package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	u "github.com/nikola43/ecoapigorm/utils"
)

type UserPromo struct {
	EmitterUserID uint   `json:"emitter_user_id"`
	Title         string `json:"title"`
	Text          string `json:"text"`
	Available     uint   `json:"available"`
	UpdatedAt     string `json:"updated_at"`
	CreatedAt     string `json:"created_at"`
	StartAt       string `json:"start_at"`
	EndAt         string `json:"end_at"`
	Week          uint   `json:"week"`
}

func (o *UserPromo) InsertUserPromo(db *sql.DB) error {

	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	statement := fmt.Sprintf("INSERT INTO users_promos (emitter_user_id, title, text, available, updated_at, created_at, start_at, end_at, week) "+
		"VALUES('%d', '%s', '%s', '%d','%s', '%s', '%s', '%s', '%d')",
		o.EmitterUserID, o.Title, o.Text, 1, date, date, o.StartAt, o.EndAt, o.Week)
	_, err := db.Exec(statement)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// get clinic id from emitter user
	var clinic_id = 0
	statement = fmt.Sprintf("SELECT clinic_id FROM users WHERE id = %d AND available = 1 AND rol != 'Cliente' ", o.EmitterUserID)
	err = db.QueryRow(statement).Scan(&clinic_id)
	if err != nil {
		fmt.Println("Error buscando el propietario de la clinica")
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(clinic_id)

	// get array of push tokens from clinic id
	var tokens []string

	switch o.Week {
	case 41:
		statement = fmt.Sprintf("SELECT push_token FROM `users` WHERE clinic_id = %d", clinic_id)
		break
	case 12:
		statement = fmt.Sprintf("SELECT u.push_token "+
			"FROM users as u "+
			"INNER JOIN calculators as c "+
			"ON u.id = c.user_id "+
			"WHERE u.clinic_id = %d "+
			"AND c.week BETWEEN %d AND 19", clinic_id, o.Week)
		break
	case 20:
		statement = fmt.Sprintf("SELECT u.push_token "+
			"FROM users as u "+
			"INNER JOIN calculators as c "+
			"ON u.id = c.user_id "+
			"WHERE u.clinic_id = %d "+
			"AND c.week BETWEEN %d AND 27", clinic_id, o.Week)
		break
	case 28:
		statement = fmt.Sprintf("SELECT u.push_token "+
			"FROM users as u "+
			"INNER JOIN calculators as c "+
			"ON u.id = c.user_id "+
			"WHERE u.clinic_id = %d "+
			"AND c.week BETWEEN %d AND 40", clinic_id, o.Week)
		break
	}

	rows, err := db.Query(statement)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var currentToken string
		if err := rows.Scan(&currentToken); err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println(currentToken)
		if len(currentToken) > 0 {
			tokens = append(tokens, currentToken)
		}
	}

	if len(tokens) == 0 {
		tokens = make([]string, 0)
	}

	fmt.Println(tokens)

	// oneSignalNotificationContent := PushNotificationContent{Text: o.Title + "\r\n" + o.Text}

	//SendPushToAllUsers(oneSignalNotificationContent)
	// SendPushToUsersRange(tokens, oneSignalNotificationContent)
	return nil
}

func (o *UserPromo) GetUserPromosByWeek(db *sql.DB) ([]UserPromo, error) {
	statement := " "

	if o.Week >= 12 && o.Week <= 19 || o.Week == 41 {
		statement = fmt.Sprintf("SELECT * FROM users_promos WHERE emitter_user_id = %d AND available = 1 AND week BETWEEN 12 AND 19 OR week = 41 ORDER BY created_at DESC", o.EmitterUserID)
	} else if o.Week >= 20 && o.Week <= 27 || o.Week == 41 {
		statement = fmt.Sprintf("SELECT * FROM users_promos WHERE emitter_user_id = %d AND available = 1  AND week BETWEEN 20 AND 27 OR week = 41 ORDER BY created_at DESC", o.EmitterUserID)
	} else if o.Week >= 28 && o.Week <= 36 || o.Week == 41 {
		statement = fmt.Sprintf("SELECT * FROM users_promos WHERE emitter_user_id = %d AND available = 1  AND week BETWEEN 28 AND 36 OR week = 41 ORDER BY created_at DESC", o.EmitterUserID)
	} else {
		list := make([]UserPromo, 0)
		return list, nil
	}
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var list []UserPromo

	for rows.Next() {
		var o UserPromo
		if err := rows.Scan(&o.EmitterUserID, &o.Title, &o.Text, &o.Available, &o.UpdatedAt, &o.CreatedAt, &o.StartAt, &o.EndAt, &o.Week); err != nil {
			return nil, err
		}
		endAt, _ := time.Parse("2006-01-02T15:04:05Z", o.StartAt)
		if time.Now().Before(endAt) {
			o.Available = 1
		} else {
			o.Available = 0
		}
		list = append(list, o)
	}

	if len(list) == 0 {
		list := make([]UserPromo, 0)
		return list, nil
	}

	return list, nil
}

func (o *UserPromo) GetUserPromosByEmitterID(db *sql.DB) ([]UserPromo, error) {
	statement := fmt.Sprintf("SELECT * FROM users_promos WHERE emitter_user_id = %d ORDER BY created_at DESC", o.EmitterUserID)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var list []UserPromo

	for rows.Next() {
		var o UserPromo
		if err := rows.Scan(&o.EmitterUserID, &o.Title, &o.Text, &o.Available, &o.UpdatedAt, &o.CreatedAt, &o.StartAt, &o.EndAt, &o.Week); err != nil {
			return nil, err
		}
		/*
		endAt, _ := time.Parse("2006-01-02T15:04:05Z", o.StartAt)

		if time.Now().Before(endAt) {
			o.Available = 1
		} else {
			o.Available = 0
		}
		*/
		list = append(list, o)
	}

	if len(list) == 0 {
		list := make([]UserPromo, 0)
		return list, nil
	}

	return list, nil
}

func SendPushToUsersRange(users []string, pushNotificationContent PushNotificationContent) {
	oneSignalApiKey := "Basic ODQzYWViMmEtZjE5Zi00OWZkLTkxY2YtMjE0Y2RlYTEwMGRi"
	oneSignalSendNotificationUrl := "https://onesignal.com/api/v1/notifications"
	oneSignalAppID := "e9c66ae8-78aa-4cc3-a4df-bf7ddab449d4"

	// create push
	oneSignalPushNotification := PushNotificationRange{
		AppID:            oneSignalAppID,
		IncludePlayerIds: users,
		Contents:         pushNotificationContent}

	// parse object to json
	notificationJsonString, loginErr := json.Marshal(oneSignalPushNotification)
	if loginErr != nil {
		fmt.Println("error")
	}

	fmt.Println(string(notificationJsonString))

	// send request
	pushNotificationResult := PushNotificationResult{}
	pushNotificationResultErr := json.Unmarshal([]byte(
		u.PostRequest(oneSignalSendNotificationUrl,
			string(notificationJsonString),
			oneSignalApiKey)),
		&pushNotificationResult)

	// check response
	if pushNotificationResultErr != nil {
		fmt.Println(pushNotificationResultErr.Error())
	}

	fmt.Println(pushNotificationResult)
}

// func (o *UserPromo) GetUserPromosByReceiverID(db *sql.DB) ([]UserPromo, error) {
// 	statement := fmt.Sprintf(
// 		"SELECT us.emitter_user_id, us.receiver_user_id, c.name as emitter_clinic_name, CONCAT(u.name , ' ', u.lastname) as receiver_user_name,  us.title, us.text, us.updated_at, us.created_at, us.start_at, us.end_at"+
// 			" FROM users_promos as us"+
// 			" LEFT OUTER JOIN users as u"+
// 			" ON us.receiver_user_id = u.id"+
// 			" LEFT OUTER JOIN clinics as c"+
// 			" ON us.emitter_user_id = c.user_id "+
// 			" WHERE us.receiver_user_id = %d", o.ReceiverUserID)

// 	rows, err := db.Query(statement)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer func() {
// 		_ = rows.Close()
// 	}()

// 	var list []UserPromo

// 	for rows.Next() {
// 		var o UserPromo
// 		if err := rows.Scan(&o.EmitterUserID, &o.EmitterClinicName, &o.ReceiverUserName, &o.Title, &o.Text, &o.UpdatedAt, &o.CreatedAt, &o.StartAt, &o.EndAt); err != nil {
// 			return nil, err
// 		}
// 		endAt, _ := time.Parse("2006-01-02T15:04:05Z", o.StartAt)
// 		if time.Now().Before(endAt) {
// 			o.Available = 1
// 		} else {
// 			o.Available = 0
// 		}
// 		list = append(list, o)
// 	}

// 	if len(list) == 0 {
// 		list := make([]UserPromo, 0)
// 		return list, nil
// 	}

// 	return list, nil
// }

func (o *UserPromo) UpdateUserPromo(db *sql.DB) error {
	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	statement := fmt.Sprintf("UPDATE users_promos SET title='%s', text='%s', updated_at='%s', start_at='%s', end_at='%s', week=%d "+
		"WHERE emitter_user_id=%d AND created_at='%s'",
		o.Title, o.Text, date, o.StartAt, o.EndAt, o.Week, o.EmitterUserID, o.CreatedAt)
	_, err := db.Exec(statement)

	// get clinic id from emitter user
	var clinic_id = 0
	statement = fmt.Sprintf("SELECT clinic_id FROM users WHERE id = %d AND available = 1 AND rol != 'Cliente' ", o.EmitterUserID)
	err = db.QueryRow(statement).Scan(&clinic_id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// get array of push tokens from clinic id
	var tokens []string

	switch o.Week {
	case 41:
		statement = fmt.Sprintf("SELECT push_token FROM `users` WHERE clinic_id = %d", clinic_id)
		break
	case 12:
		statement = fmt.Sprintf("SELECT u.push_token "+
			"FROM users as u "+
			"INNER JOIN calculators as c "+
			"ON u.id = c.user_id "+
			"WHERE u.clinic_id = %d "+
			"AND c.week BETWEEN %d AND 19", clinic_id, o.Week)
		break
	case 20:
		statement = fmt.Sprintf("SELECT u.push_token "+
			"FROM users as u "+
			"INNER JOIN calculators as c "+
			"ON u.id = c.user_id "+
			"WHERE u.clinic_id = %d "+
			"AND c.week BETWEEN %d AND 27", clinic_id, o.Week)
		break
	case 28:
		statement = fmt.Sprintf("SELECT u.push_token "+
			"FROM users as u "+
			"INNER JOIN calculators as c "+
			"ON u.id = c.user_id "+
			"WHERE u.clinic_id = %d "+
			"AND c.week BETWEEN %d AND 40", clinic_id, o.Week)
		break
	}

	rows, err := db.Query(statement)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var currentToken string
		if err := rows.Scan(&currentToken); err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println(currentToken)
		if len(currentToken) > 0 {
			tokens = append(tokens, currentToken)
		}
	}

	if len(tokens) == 0 {
		tokens = make([]string, 0)
	}

	fmt.Println(tokens)

	// oneSignalNotificationContent := PushNotificationContent{Text: o.Title + "\r\n" + o.Text}

	//SendPushToAllUsers(oneSignalNotificationContent)
	// SendPushToUsersRange(tokens, oneSignalNotificationContent)
	return err
}

func (o *UserPromo) DeleteUserPromo(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM users_promos WHERE emitter_user_id=%d AND title='%s'", o.EmitterUserID, o.Title)
	fmt.Println(statement)
	_, err := db.Exec(statement)
	return err
}

func (o *UserPromo) GetUserPromos(db *sql.DB) ([]UserPromo, error) {
	statement := fmt.Sprintf(
		"SELECT us.emitter_user_id, us.title, us.text, us.updated_at, us.created_at, us.start_at, us.end_at, us.week" +
			" FROM users_promos as us" +
			" LEFT OUTER JOIN clinics as c" +
			" ON us.emitter_user_id = c.user_id ORDER BY created_at DESC")

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var list []UserPromo

	for rows.Next() {
		var o UserPromo
		if err := rows.Scan(&o.EmitterUserID, &o.Title, &o.Text, &o.UpdatedAt, &o.CreatedAt, &o.StartAt, &o.EndAt, &o.Week); err != nil {
			return nil, err
		}

		/*
			endAt, _ := time.Parse("2006-01-02T15:04:05Z", o.StartAt)
			if endAt.Before(time.Now()) {
				o.Available = 1
			} else {
				o.Available = 0
			}
		*/
		list = append(list, o)
	}
	return list, nil
}

/*
func SendPushToUsersRange(users []string, pushNotificationContent PushNotificationContent) {
	oneSignalApiKey := "Basic ODQzYWViMmEtZjE5Zi00OWZkLTkxY2YtMjE0Y2RlYTEwMGRi"
	oneSignalSendNotificationUrl := "https://onesignal.com/api/v1/notifications"
	oneSignalAppID := "e9c66ae8-78aa-4cc3-a4df-bf7ddab449d4"

	// create push
	oneSignalPushNotification := PushNotification{
		AppID:            oneSignalAppID,
		IncludePlayerIds: users,
		Contents:         pushNotificationContent}

	// parse object to json
	notificationJsonString, loginErr := json.Marshal(oneSignalPushNotification)
	if loginErr != nil {
		fmt.Println("error")
	}

	// send request
	pushNotificationResult := PushNotificationResult{}
	pushNotificationResultErr := json.Unmarshal([]byte(
		u.PostRequest(oneSignalSendNotificationUrl,
			string(notificationJsonString),
			oneSignalApiKey)),
		&pushNotificationResult)

	// check response
	if pushNotificationResultErr != nil {
		fmt.Println(pushNotificationResultErr.Error())
	}

	fmt.Println(pushNotificationResult)
}
*/

func SendPushToAllUsers(pushNotificationContent PushNotificationContent) {
	oneSignalApiKey := "Basic ODQzYWViMmEtZjE5Zi00OWZkLTkxY2YtMjE0Y2RlYTEwMGRi"
	oneSignalSendNotificationUrl := "https://onesignal.com/api/v1/notifications"
	oneSignalAppID := "e9c66ae8-78aa-4cc3-a4df-bf7ddab449d4"

	// create push
	oneSignalPushNotification := PushNotificationAll{
		AppID:            oneSignalAppID,
		IncludedSegments: []string{"All"},
		Contents:         pushNotificationContent}

	// parse object to json
	notificationJsonString, loginErr := json.Marshal(oneSignalPushNotification)
	if loginErr != nil {
		fmt.Println("error")
	}

	fmt.Println(string(notificationJsonString))

	// send request
	pushNotificationResult := PushNotificationResult{}
	pushNotificationResultErr := json.Unmarshal([]byte(
		u.PostRequest(oneSignalSendNotificationUrl,
			string(notificationJsonString),
			oneSignalApiKey)),
		&pushNotificationResult)

	// check response
	if pushNotificationResultErr != nil {
		fmt.Println(pushNotificationResultErr.Error())
	}

	fmt.Println(pushNotificationResult)
}
