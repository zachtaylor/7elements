package db

import (
	"io/ioutil"
	"strconv"
	"time"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

func EnvPatchDir() string {
	return env.Default("DB_PATCHES", "patches")
}

func Patch() int {
	patchdir := EnvPatchDir()
	log := log.Add("patch-path", patchdir)

	patch, err := scanPatch()

	if err != nil {
		if err.Error() == "no such table: patch" {
			_, err := Connection.Exec("CREATE TABLE patch(patch INTEGER); INSERT INTO patch(patch) VALUES(0);")
			if err != nil {
				log.Add("Error", err).Error("patch: cannot create patch table")
				return 0
			}
			log.Info("patch: created new database")
		} else {
			log.Add("Error", err).Error("patch: cannot read database")
			return 0
		}
	}

	files, err := ioutil.ReadDir(patchdir)
	if err != nil {
		log.Add("Error", err).Error("patch: cannot read patch-path")
	}

	for _, f := range files {
		patch = tryApplyPatch(patch, patchdir, f.Name())
	}

	return patch

}

func tryApplyPatch(patch int, filepath, filename string) int {
	log := log.Add("File", filename)

	if filename[len(filename)-4:] != ".sql" {
		log.Warn("patch: file is not .sql type")
		return patch
	}

	patchid, err := strconv.ParseInt(filename[:4], 10, 64)
	if err != nil {
		log.Add("Error", err).Error("patch: file is not .sql type")
	}

	if int(patchid) <= patch {
		return patch
	} else if int(patchid) > patch+1 {
		log.Warn("patch: number not next in sequence")
		return patch
	}

	file, err := ioutil.ReadFile(filepath + filename)

	if err != nil {
		log.Add("Error", err).Error("patch: file read")
		return patch
	}

	tStart := time.Now()
	if _, err = Connection.Exec(string(file)); err != nil {
		log.Add("Error", err).Error("patch: failed")
		return patch
	}

	patch++
	Connection.Exec("UPDATE patch SET patch=?", patch)

	log.Add("Time", time.Now().Sub(tStart)).Info("patch")
	return patch
}

func scanPatch() (int, error) {
	var patch int
	row := Connection.QueryRow("SELECT * FROM patch")
	err := row.Scan(&patch)
	if err != nil {
		return 0, err
	}
	return patch, nil
}
