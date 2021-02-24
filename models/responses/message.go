package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Message struct {
	ID        uint   `json:"id"`
	Text      string `json:"text"`
	Available uint   `json:"available"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

func (o *Message) insertMessage(db *sql.DB) error {
	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	statement := fmt.Sprintf(
		"INSERT INTO messages (text, available, updated_at, created_at) " +
		"VALUES('%s', '%d', '%s', '%s')",
		o.Text, o.Available, date, date)
	
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	return nil
}

func (o *Message) getMessage(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT * FROM messages WHERE id=%d AND available = 1", o.ID)
	return db.QueryRow(statement).Scan(&o.ID, &o.Text, &o.Available, &o.UpdatedAt, &o.CreatedAt)
}

func (o *Message) updateMessage(db *sql.DB) error {
	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	statement := fmt.Sprintf(
		"UPDATE messages " +
		"SET text='%s', available='%d', updated_at='%s' " +
		"WHERE id=%d",
		o.Text, o.Available, date, o.ID)
	_, err := db.Exec(statement)
	return err
}

func (o *Message) deleteMessage(db *sql.DB) error {
	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	statement := fmt.Sprintf(
		"UPDATE messages " +
			"SET available='%d', updated_at='%s' " +
			"WHERE id=%d",
		0, date, o.ID)
	_, err := db.Exec(statement)
	fmt.Println(statement)
	return err
}

func getMessages(db *sql.DB) ([]Message, error) {
	var list []Message

	rows, err := db.Query("SELECT * FROM messages WHERE available = 1")

	if err != nil {
		return nil, err
	}


	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var o Message
		if err := rows.Scan(&o.ID, &o.Text, &o.Available, &o.UpdatedAt, &o.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, nil
}
