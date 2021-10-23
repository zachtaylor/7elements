package runtime

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/zachtaylor/7elements/db"
	"taylz.io/db/patch"
	"taylz.io/env"
	"taylz.io/http/hash"
	"taylz.io/http/session"
	"taylz.io/keygen"
	"taylz.io/log"
	"taylz.io/types"
)

func Parse(env env.Service, ex_patch int) (*T, error) {
	var isdev bool
	if env["ENV"] == "dev" {
		isdev = true
	}

	passhash := hash.NewMD5SaltHash(env["DB_PWSALT"])

	var conn *db.DB
	var err error
	if conn, err = db.OpenEnv(env); err != nil {
		return nil, err
	} else if patch, err := patch.Get(conn); err != nil {
		return nil, err
	} else if patch != ex_patch {
		return nil, errors.New("Patch mismatch: " + strconv.FormatInt(int64(patch), 10))
	}

	var logger *log.T
	filePath := types.NewSource(0).File()
	for i := 0; i < 3; i++ {
		filePath = filePath[:strings.LastIndex(filePath, "/")]
	}
	loglvl, err := log.GetLevel(env["LOG_LEVEL"])
	if err != nil {
		return nil, err
	} else if isdev {
		logger = log.Lining(log.IOLiner(&log.ColorFormat{
			Colors:     log.DefaultColorMap(),
			ColorMsg:   true,
			ColorField: true,
			SrcFmt: log.RestringSourceFormatter(
				log.DetailSourceFormatter(),
				log.RestringerMiddleware(log.RestringerCutPrefixes{filePath}, log.RestringerLenExact(40)),
			),
			TimeFmt: log.DefaultTimeFormatter(),
		}, os.Stdout))
	} else {
		logger = log.Lining(log.LevelLiner(loglvl, log.IOLiner(&log.ColorFormat{
			SrcFmt:  log.ClassicSourceFormatter(filePath),
			TimeFmt: log.DefaultTimeFormatter(),
		}, log.DailyRotatingFile(env["LOG_PATH"]))))
	}

	sessionSettings := session.DefaultSettings(keygen.NewFunc(8))
	if isdev {
		sessionSettings.Lifetime = 10 * types.Minute
	} else {
		sessionSettings.Secure = true
		sessionSettings.Lifetime = 1 * types.Hour
	}

	wsKeygen := keygen.NewFunc(12)
	chatKeygen := keygen.NewFunc(4)

	return New(
		isdev,
		passhash,
		conn,
		logger,
		sessionSettings,
		wsKeygen,
		chatKeygen,
	)
}
