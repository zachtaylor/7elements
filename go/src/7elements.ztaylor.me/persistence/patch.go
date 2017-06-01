package persistence

import (
	"7elements.ztaylor.me/log"
	"io/ioutil"
	"strconv"
)

func Patch(patchpath string) {
	patch, err := GetPatch()

	if err != nil {
		if err.Error() == "no such table: patch" {
			_, err := connection.Exec("CREATE TABLE patch(patch INTEGER); INSERT INTO patch(patch) VALUES(0);")
			if err != nil {
				log.Add("Error", err).Error("patch: cannot create patch table")
				return
			}
			log.Info("patch: created new database")
		} else {
			log.Add("Error", err).Error("patch: cannot read database")
			return
		}
	}

	files, _ := ioutil.ReadDir(patchpath)
	for _, f := range files {
		fname := f.Name()
		path := patchpath + fname

		if fname[len(fname)-4:] != ".sql" {
			log.Add("File", f.Name()).Warn("patch: file is not .sql type")
			continue
		}

		if patchid, err := strconv.ParseUint(fname[:4], 10, 64); err == nil {
			if patchid == patch+1 {
				file, err := ioutil.ReadFile(path)
				if err != nil {
					log.Add("Error", err).Add("File", f.Name()).Error("patch: file read")
					continue
				}

				_, err = connection.Exec(string(file))
				if err != nil {
					log.Add("Error", err).Add("File", f.Name()).Error("patch: file exec")
					continue
				}

				patch++
				connection.Exec("UPDATE patch SET patch=?", patch)
				log.Add("Processed", f.Name()).Info("patch: applied patch")
			} else if patchid > patch {
				log.Add("File", f.Name()).Warn("patch: number not next in sequence")
			}
		} else {
			log.Add("File", f.Name()).Add("Error", err).Warn("patch: number not identified")
		}
	}
}

func GetPatch() (uint64, error) {
	var patch uint64
	row := connection.QueryRow("SELECT * FROM patch")
	err := row.Scan(&patch)

	if err != nil {
		return 0, err
	}
	return patch, nil
}
