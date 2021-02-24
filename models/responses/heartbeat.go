package models

import (
	"database/sql"
	"fmt"
)

type Heartbeat struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Url    string `json:"url"`
}

func (o *Heartbeat) GetHeartbeatByUserID(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT * FROM heartbeats WHERE user_id=%d LIMIT 1", o.UserID)
	return db.QueryRow(statement).Scan(&o.ID, &o.UserID, &o.Url)
}

func (o *Heartbeat) DeleteHeartbeatByUserID(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM heartbeats WHERE user_id=%d", o.UserID)
	fmt.Println(statement)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	return nil
}

func (o *Heartbeat) InsertHeartbeat(db *sql.DB) (sql.Result, error) {
	// date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	statement := fmt.Sprintf(
		"INSERT INTO heartbeats (user_id, url) "+
			"VALUES('%d', '%s')",
		o.UserID, o.Url)
	res, err := db.Exec(statement)
	if err != nil {
		return nil, err
	}
	return res, nil
}
