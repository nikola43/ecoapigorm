package models

import (
	"database/sql"
	"fmt"
	u "github.com/nikola43/ecoapigorm/utils"
	"github.com/sethvargo/go-password/password"
	"strings"
	"time"
)

type Streaming struct {
	ID        uint   `json:"id"`
	ClinicID  uint   `json:"clinic_id"`
	Url       string `json:"url"`
	Code      string `json:"code"`
	Available uint   `json:"available"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"`
}

func (o *Streaming) InsertStreaming(db *sql.DB) error {
	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	codeExists := true

	for ok := true; ok; ok = codeExists {
		code, passwordErr := password.Generate(4, 0, 0, true, false)
		if passwordErr != nil {
			return passwordErr
		}

		codeQueryStatement := fmt.Sprintf("SELECT * FROM streamings WHERE code = '%s' AND available=1", code)
		codeExistsErr := db.QueryRow(codeQueryStatement).Scan()
		if codeExistsErr != nil {
			if codeExistsErr == sql.ErrNoRows {
				codeExists = false
				o.Code = strings.ToUpper(code)
			}
		}
	}

	statement := fmt.Sprintf(
		"INSERT INTO streamings (id, clinic_id, url, code,  available, updated_at, created_at)"+
			"VALUES(NULL, '%d', '%s','%s', '%d', '%s', '%s')",
		o.ClinicID, o.Url, o.Code, 1, date, date)

	fmt.Println(o)
	fmt.Println(statement)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}

	queryStatement := fmt.Sprintf("SELECT s.id, s.clinic_id, s.url, s.code, s.available, s.updated_at, s.created_at "+
		"FROM streamings as s "+
		"WHERE s.code = '%s' AND s.available = 1", o.Code)

	d := u.Info{Data: o.Code}
	d.SendStreamingMail(o.Username)

	return db.QueryRow(queryStatement).Scan(&o.ID, &o.ClinicID, &o.Url, &o.Code, &o.Available, &o.UpdatedAt, &o.CreatedAt)
}

func (o *Streaming) GetStreaming(db *sql.DB) error {
	queryStatement := fmt.Sprintf("SELECT * "+
		"FROM streamings "+
		"WHERE code = '%s'", o.Code)
	return db.QueryRow(queryStatement).Scan(&o.ID, &o.ClinicID, &o.Url, &o.Code, &o.Available, &o.UpdatedAt, &o.CreatedAt)
}

func (o *Streaming) GetStreamingByClinicId(db *sql.DB) ([]Streaming, error) {
	var list []Streaming

	statement := fmt.Sprintf("SELECT s.id, s.clinic_id, s.url,  s.code, s.available, s.updated_at, s.created_at "+
		"FROM streamings as s "+
		"WHERE s.available = 1 "+
		"AND s.clinic_id = '%d'"+
		"ORDER BY s.created_at DESC", o.ClinicID)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var o Streaming
		if err := rows.Scan(&o.ID, &o.ClinicID, &o.Url, &o.Code, &o.Available, &o.UpdatedAt, &o.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, o)
	}

	if len(list) == 0 {
		list := make([]Streaming, 0)
		return list, nil
	}

	return list, nil
}

func (o *Streaming) GetStreamings(db *sql.DB) ([]Streaming, error) {
	var list []Streaming

	statement := fmt.Sprintf("SELECT s.id, s.clinic_id, s.url, s.code, s.available, s.updated_at, s.created_at " +
		"FROM streamings as s " +
		"WHERE s.available = 1 " +
		"ORDER BY s.created_at DESC")

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var o Streaming
		if err := rows.Scan(&o.ID, &o.ClinicID, &o.Url, &o.Code, &o.Available, &o.UpdatedAt, &o.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, o)
	}

	if len(list) == 0 {
		list := make([]Streaming, 0)
		return list, nil
	}

	return list, nil
}
