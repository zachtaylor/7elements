package db

import (
	"time"

	"ztaylor.me/http/track"
)

func init() {
	httptrack.Service = HttpTrackService(1)
}

type HttpTrackService int

func (s HttpTrackService) GetAccountAddrs(name string) ([]*httptrack.LoginDetails, error) {
	rows, err := Conn.Query("SELECT name, addr, t FROM httptrack WHERE name=?",
		name,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]*httptrack.LoginDetails, 0)

	for rows.Next() {
		dat := &httptrack.LoginDetails{}
		var tbuff int64
		if err = rows.Scan(&dat.Name, &dat.Addr, &tbuff); err != nil {
			return nil, err
		}
		dat.Time = time.Unix(tbuff, 0)
		data = append(data, dat)
	}

	return data, nil
}

func (s HttpTrackService) UpdateQuery(data *httptrack.LoginDetails) error {
	if res, err := Conn.Exec(
		"UPDATE httptrack SET heat = heat + 1, t=? WHERE name=? AND addr=?",
		data.Time.Unix(),
		data.Name,
		data.Addr,
	); err != nil {
		return err
	} else if rownum, _ := res.RowsAffected(); rownum < 1 {
		return ErrUpdateFailed
	} else {
		return nil
	}
}

func (s HttpTrackService) InsertQuery(data *httptrack.LoginDetails) error {
	if _, err := Conn.Exec(
		"INSERT INTO httptrack (name, addr, heat, t) VALUES (?, ?, 0, ?)",
		data.Name,
		data.Addr,
		data.Time.Unix(),
	); err != nil {
		return err
	} else {
		return nil
	}
}

func (s HttpTrackService) SaveAccountAddr(name string, addr string, t time.Time) error {
	data := &httptrack.LoginDetails{name, addr, t}
	if err := s.UpdateQuery(data); err == nil {
		return nil
	} else if err == ErrUpdateFailed {
		return s.InsertQuery(data)
	} else {
		return err
	}
}
