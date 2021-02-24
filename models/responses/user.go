package models

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"

	u "github.com/nikola43/ecoapigorm/utils"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username      string         `json:"username"`
	Password      string         `json:"password"`
	Name          string `json:"name"`
	Lastname      string `json:"lastname"`
	Phone         string `json:"phone"`
}

type UserRecovery struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Token     string `json:"token"`
	Available string `json:"available"`
	Password  string `json:"password"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

func (o *User) InsertUser(db *sql.DB) (error, map[string]interface{}) {

	//if resp, ok := o.Validate(db); !ok {
	//	return nil, resp
	//}
	//
	//date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	//statement := fmt.Sprintf(
	//	"INSERT INTO users (clinic_id, username, password, name, lastname, phone, rol, available, updated_at, created_at) "+
	//		"VALUES('%d', '%s', '%s', '%s', '%s', '%s', '%s', '%d', '%s', '%s')",
	//	o.ClinicID.Int64, strings.ToLower(o.Username), hashAndSalt([]byte("babyandme")), o.Name.String, o.Lastname.String, o.Phone.String, o.Rol, 1, date, date)
	//res, err := db.Exec(statement)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return err, nil
	//}
	//
	//o.Password = ""
	//lastId, _ := res.LastInsertId()
	//o.ID = uint(lastId)
	//response := u.Message(http.StatusOK, "success")
	//response["user"] = o
	//
	////d := u.Info{Data: o.Username}
	////d.SendWelcomeMail(o.Username)

	//return nil, response
	return nil, nil
}

func (o *User) GetUserFromUsername(db *sql.DB) error {
	//statement := fmt.Sprintf(
	//	"SELECT u.`id`, u.`clinic_id`, c.`name`, u.`username`, u.`name`, u.`lastname`,"+
	//		" u.`phone`, u.`rol`, u.`available`, u.`token`, u.`firebase_token`, u.`push_token`, u.`device_type`, u.`updated_at`, u.`created_at`, cal.`week`"+
	//		" FROM `users` AS u "+
	//		"INNER JOIN calculators AS cal ON u.`id` = cal.`user_id` "+
	//		"INNER JOIN clinics AS c ON u.`clinic_id` = c.`id` "+
	//		"WHERE u.`username`= '%s'  AND u.`available` = 1", o.Username)
	//return db.QueryRow(statement).Scan(&o.ID, &o.ClinicID, &o.Username, &o.Name, &o.Lastname, &o.Phone, &o.Rol, &o.Available, &o.Token, &o.FirebaseToken, &o.PushToken, &o.DeviceType, &o.UpdatedAt, &o.CreatedAt)
	return nil
}

func (o *User) GetUser(db *sql.DB) error {
	//statement := fmt.Sprintf(
	//	"SELECT u.`id`, u.`clinic_id`, c.`name`, u.`username`, u.`name`, u.`lastname`,"+
	//		" u.`phone`, u.`rol`, u.`available`, u.`token`, u.`firebase_token`, u.`push_token`, u.`device_type`, u.`updated_at`, u.`created_at`, cal.`week`"+
	//		" FROM `users` AS u "+
	//		"INNER JOIN calculators AS cal ON u.`id` = cal.`user_id` "+
	//		"INNER JOIN clinics AS c ON u.`clinic_id` = c.`id` "+
	//		"WHERE u.`id` = %d AND u.`available` = 1", o.ID)
	//return db.QueryRow(statement).Scan(&o.ID, &o.ClinicID, &o.Username, &o.Name, &o.Lastname, &o.Phone, &o.Rol, &o.Available, &o.Token, &o.FirebaseToken, &o.PushToken, &o.DeviceType, &o.UpdatedAt, &o.CreatedAt)
	return nil
}

func (o *User) UpdateDeviceType(db *sql.DB) error {
	//date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	//statement := fmt.Sprintf(
	//	"UPDATE users "+
	//		"SET device_type='%s', updated_at='%s' "+
	//		"WHERE username='%s'",
	//	o.DeviceType.String, date, o.Username)
	//_, err := db.Exec(statement)
	//return err
	return nil
}

func (o *User) UpdateUserApiTokenByUsername(db *sql.DB) error {
	//date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	//statement := fmt.Sprintf(
	//	"UPDATE users "+
	//		"SET token='%s', updated_at='%s' "+
	//		"WHERE username='%s'",
	//	o.Token.String, date, o.Username)
	//_, err := db.Exec(statement)
	//return err
	return nil
}

func (o *User) UpdateUserFirebaseTokenByUsername(db *sql.DB) error {
	//date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	//statement := fmt.Sprintf(
	//	"UPDATE users "+
	//		"SET firebase_token='%s', updated_at='%s' "+
	//		"WHERE username='%s'",
	//	o.FirebaseToken.String, date, o.Username)
	//_, err := db.Exec(statement)
	//return err
	return nil
}

func (o *User) UpdateUserPushTokenByUsername(db *sql.DB) error {
	//date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	//statement := fmt.Sprintf(
	//	"UPDATE users "+
	//		"SET push_token='%s', updated_at='%s' "+
	//		"WHERE username='%s'",
	//	o.PushToken.String, date, o.Username)
	//_, err := db.Exec(statement)
	//return err
	return nil
}

func (o *User) RemoveApiAndPushTokens(db *sql.DB) error {
	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	statement := fmt.Sprintf(
		"UPDATE users "+
			"SET token='%s', push_token='%s', updated_at='%s' "+
			"WHERE username='%s'",
		"", "", date, o.Username)
	_, err := db.Exec(statement)
	return err
}

func (o *User) ActivateUser(db *sql.DB) error {
	//statement := fmt.Sprintf(
	//	"UPDATE users SET available = %d WHERE id = %d", o.Available, o.ID)
	//_, err := db.Exec(statement)
	//fmt.Println(statement)
	//fmt.Println(err)
	//return err
	return nil
}

func (o *User) UpdateUser(db *sql.DB) error {
	//date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	//statement := fmt.Sprintf(
	//	"UPDATE users "+
	//		"SET available=1, clinic_id=%d, username='%s', name='%s', lastname='%s', phone='%s', rol='%s', updated_at='%s' "+
	//		"WHERE id=%d",
	//	o.ClinicID.Int64, o.Username, o.Name.String, o.Lastname.String, o.Phone.String, o.Rol, date, o.ID)
	//_, err := db.Exec(statement)
	//return err
	return nil
}

func (o *User) DeleteUser(db *sql.DB) error {
	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	statement := fmt.Sprintf(
		"UPDATE users "+
			"SET available='%d', updated_at='%s' "+
			"WHERE id=%d",
		0, date, o.ID)
	_, err := db.Exec(statement)
	return err
}

func (o *User) GetNonAdminUsers(db *sql.DB) ([]User, error) {
	//var list []User
	//	//
	//	//statement := fmt.Sprintf(
	//	//	"SELECT u.`id`, u.`clinic_id`, c.`name`, u.`username`, u.`name`, u.`lastname`," +
	//	//		" u.`phone`, u.`rol`, u.`available`, u.`token`, u.`firebase_token`, u.`push_token`, u.`device_type`, u.`updated_at`, u.`created_at`, cal.`week`" +
	//	//		" FROM `users` AS u " +
	//	//		"INNER JOIN calculators AS cal ON u.`id` = cal.`user_id` " +
	//	//		"INNER JOIN clinics AS c ON u.`clinic_id` = c.`id` " +
	//	//		"WHERE u.`available` = 1 AND u.rol = 'Cliente' ORDER BY u.`created_at` DESC")
	//	//
	//	//rows, err := db.Query(statement)
	//	//
	//	//if err != nil {
	//	//	return nil, err
	//	//}
	//	//
	//	//defer func() {
	//	//	_ = rows.Close()
	//	//}()
	//	//
	//	//for rows.Next() {
	//	//	var o User
	//	//	if err := rows.Scan(&o.ID, &o.ClinicID, &o.Username, &o.Name, &o.Lastname, &o.Phone, &o.Rol, &o.Available, &o.Token, &o.FirebaseToken, &o.PushToken, &o.DeviceType, &o.UpdatedAt, &o.CreatedAt); err != nil {
	//	//		return nil, err
	//	//	}
	//	//	list = append(list, o)
	//	//}
	//	//
	//	//if len(list) == 0 {
	//	//	list := make([]User, 0)
	//	//	return list, nil
	//	//}
	//	//
	//	//return list, nil
	return nil, nil
}

func (o *User) GetNonOwnersUsers(db *sql.DB) ([]User, error) {
	//var list []User
	//
	//statement :=
	//	"SELECT u.`id`, u.`clinic_id`, c.`name`, u.`username`, u.`name`, u.`lastname`," +
	//		" u.`phone`, u.`rol`, u.`available`, u.`token`, u.`firebase_token`, u.`push_token`, u.`device_type`, u.`updated_at`, u.`created_at`, cal.`week`" +
	//		" FROM `users` AS u " +
	//		"INNER JOIN calculators AS cal ON u.`id` = cal.`user_id` " +
	//		"INNER JOIN clinics AS c ON u.`clinic_id` = c.`id` " +
	//		"WHERE u.`available` = 1 " +
	//		"AND u.`rol` != 'Cliente' " +
	//		"AND u.`rol` != 'Master' " +
	//		"AND u.`id` NOT IN (SELECT user_id FROM clinics) " +
	//		"ORDER BY u.`created_at` DESC"
	//
	//rows, err := db.Query(statement)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//defer func() {
	//	_ = rows.Close()
	//}()
	//
	//for rows.Next() {
	//	var o User
	//	if err := rows.Scan(&o.ID, &o.ClinicID, &o.Username, &o.Name, &o.Lastname, &o.Phone, &o.Rol, &o.Available, &o.Token, &o.FirebaseToken, &o.PushToken, &o.DeviceType, &o.UpdatedAt, &o.CreatedAt); err != nil {
	//		return nil, err
	//	}
	//	list = append(list, o)
	//}
	//
	//if len(list) == 0 {
	//	list := make([]User, 0)
	//	return list, nil
	//}
	//
	//return list, nil
	return nil,nil
}

func (o *User) GetUsersByClinicID(db *sql.DB) ([]User, error) {
	//var list []User
	//
	//statement := fmt.Sprintf(
	//	"SELECT u.`id`, u.`clinic_id`, c.`name`, u.`username`, u.`name`, u.`lastname`,"+
	//		" u.`phone`, u.`rol`, u.`available`, u.`token`, u.`firebase_token`, u.`push_token`, u.`device_type`, u.`updated_at`, u.`created_at`, cal.`week`"+
	//		" FROM `users` AS u "+
	//		"INNER JOIN calculators AS cal ON u.`id` = cal.`user_id` "+
	//		"INNER JOIN clinics AS c ON u.`clinic_id` = c.`id` "+
	//		"WHERE u.`clinic_id` = %d "+
	//		"AND u.`rol` != 'Propietario' "+
	//		"AND u.`rol` != 'Master' "+
	//		"ORDER BY u.`created_at` DESC", o.ID)
	//
	//rows, err := db.Query(statement)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//defer func() {
	//	_ = rows.Close()
	//}()
	//
	//for rows.Next() {
	//	var o User
	//	if err := rows.Scan(&o.ID, &o.ClinicID, &o.Username, &o.Name, &o.Lastname, &o.Phone, &o.Rol, &o.Available, &o.Token, &o.FirebaseToken, &o.PushToken, &o.DeviceType, &o.UpdatedAt, &o.CreatedAt); err != nil {
	//		return nil, err
	//	}
	//	list = append(list, o)
	//}
	//
	//if len(list) == 0 {
	//	list := make([]User, 0)
	//	return list, nil
	//}
	//
	//return list, nil
	return nil,nil
}

func (o *User) GetUsers(db *sql.DB) ([]User, error) {
	//var list []User
	//
	//statement := fmt.Sprintf(
	//	"SELECT u.`id`, u.`clinic_id`, c.`name`, u.`username`, u.`name`, u.`lastname`," +
	//		" u.`phone`, u.`rol`, u.`available`, u.`token`, u.`firebase_token`, u.`push_token`, u.`device_type`, u.`updated_at`, u.`created_at`, cal.`week`" +
	//		" FROM `users` AS u " +
	//		"INNER JOIN calculators AS cal ON u.`id` = cal.`user_id` " +
	//		"INNER JOIN clinics AS c ON u.`clinic_id` = c.`id` "+
	//		"ORDER BY u.`created_at` DESC")
	//rows, err := db.Query(statement)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//defer func() {
	//	_ = rows.Close()
	//}()
	//
	//for rows.Next() {
	//	var o User
	//	if err := rows.Scan(&o.ID, &o.ClinicID, &o.Username, &o.Name, &o.Lastname, &o.Phone, &o.Rol, &o.Available, &o.Token, &o.FirebaseToken, &o.PushToken, &o.DeviceType, &o.UpdatedAt, &o.CreatedAt); err != nil {
	//		return nil, err
	//	}
	//	list = append(list, o)
	//}
	//
	//if len(list) == 0 {
	//	list := make([]User, 0)
	//	return list, nil
	//}
	//
	//return list, nil
	return nil,nil
}

func (o *User) Validate(db *sql.DB) (map[string]interface{}, bool) {
	// Generate random password
	res, passwordGenerateErr := password.Generate(8, 2, 0, false, false)
	if passwordGenerateErr != nil {
		log.Fatal(passwordGenerateErr)
	}
	o.Password = res

	// get username from db
	var username string
	statement := fmt.Sprintf("SELECT username FROM users WHERE username = '%s'", o.Username)
	queryErr := db.QueryRow(statement).Scan(&username)

	// check for any error
	if queryErr != nil && queryErr != sql.ErrNoRows {
		return u.Message(http.StatusForbidden, queryErr.Error()), false
	}

	// check if username exist
	if username != "" {
		return u.Message(http.StatusUnprocessableEntity, "Username already in use by another user."), false
	}
	return u.Message(http.StatusOK, "Requirement passed"), true
}

func (o *User) GetRefreshToken(db *sql.DB, username string, password string) (string, error) {
 return "", nil
}


func (o *User) Login(db *sql.DB, username string, password string) (*User, error) {
	//fmt.Println(o)
	//
	//// get cast type
	//deviceType := 0
	//
	//switch o.DeviceType.String {
	//case "android":
	//	deviceType = 1
	//	break
	//case "ios":
	//	deviceType = 0
	//	break
	//case "web":
	//	deviceType = 2
	//	break
	//default:
	//	deviceType = 2
	//	break
	//}
	//
	//// get user
	//user := &User{}
	//
	//statement := fmt.Sprintf(
	//	"SELECT u.`id`, u.`clinic_id`, c.`name`, u.`username`,  u.`password`, u.`name`, u.`lastname`,"+
	//		" u.`phone`, u.`rol`, u.`available`, u.`token`, u.`firebase_token`, u.`push_token`, u.`device_type`, u.`updated_at`, u.`created_at`, cal.`week`, c.`user_id`, cal.`updated_at` as last_week_calc "+
	//		" FROM `users` AS u "+
	//		"INNER JOIN calculators AS cal ON u.`id` = cal.`user_id` "+
	//		"INNER JOIN clinics AS c ON u.`clinic_id` = c.`id` "+
	//		"WHERE u.`available` = 1 "+
	//		"AND u.`username` = '%s'", o.Username)
	//
	//err := db.QueryRow(statement).Scan(&user.ID, &user.ClinicID, &user.Username, &user.Password, &user.Name, &user.Lastname, &user.Phone, &user.Rol, &user.Available, &user.Token, &user.FirebaseToken, &user.PushToken, &user.DeviceType, &user.UpdatedAt, &user.CreatedAt)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// check if password match
	//match := comparePasswords(user.Password, []byte(password))
	//if match == false {
	//	return nil, sql.ErrNoRows
	//}
	//
	//// remove password
	//user.Password = ""
	//
	//// update lastest login device type
	//updateError := o.UpdateDeviceType(db)
	//if updateError != nil {
	//	fmt.Print(updateError.Error())
	//}
	//
	//// Declare the expiration time of the token
	//// here, we have kept it as 60 days
	//expirationTime := time.Now().Add(100160 * time.Minute)
	//
	//// Create the JWT claims, which includes the username and expiry time
	//claims := &FirebaseToken{
	//	Username: user.Username,
	//	StandardClaims: jwt.StandardClaims{
	//		ExpiresAt: expirationTime.Unix(),
	//	},
	//}
	//
	//// generate api token
	//apiToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	//apiTokenString, _ := apiToken.SignedString([]byte(os.Getenv("token_password")))
	//user.Token = sql.NullString{String: apiTokenString, Valid: true}
	//
	//// update api token
	//updateApiError := user.UpdateUserApiTokenByUsername(db)
	//if updateApiError != nil {
	//	fmt.Println(updateApiError.Error())
	//}
	//
	//
	//oneSignalToken := "Basic NWMwYTNmZTEtOGU4Ni00OTlhLWFlZjgtNGJkNTY1MTFiZWM2"
	//oneSignalInsertUserUrl := "https://onesignal.com/api/v1/players"
	//oneSignalAppID := "0f19ed5b-2b9b-4492-a8ef-298845ce7d21"
	//
	//fmt.Println(o.FirebaseToken)
	//
	//// create onesignal user
	//oneSignalUser := &OneSignalUser{AppID: oneSignalAppID, DeviceType: deviceType, Identifier: o.FirebaseToken.String}
	//oneSignalUserString, loginErr := json.Marshal(oneSignalUser)
	//if loginErr != nil {
	//	fmt.Println("error")
	//}
	//fmt.Println("oneSignalUser")
	//fmt.Println(oneSignalUserString)
	//
	//// make insert request
	//oneSignalInsertUserResult := OneSignalInsertUserResult{}
	//oneSignalInsertUserResultErr := json.Unmarshal([]byte(u.PostRequest(oneSignalInsertUserUrl, string(oneSignalUserString), oneSignalToken)), &oneSignalInsertUserResult)
	//fmt.Println("oneSignalInsertUserResult")
	//fmt.Println(oneSignalInsertUserResult)
	//fmt.Println("oneSignalInsertUserResultErr")
	//fmt.Println(oneSignalInsertUserResultErr)
	//if oneSignalInsertUserResultErr != nil {
	//	fmt.Println(oneSignalInsertUserResultErr.Error())
	//}
	//
	//fmt.Println("oneSignalInsertUserResult.ID")
	//fmt.Println(oneSignalInsertUserResult.ID)
	//// update user push token
	////user.PushToken = sql.NullString{String: oneSignalInsertUserResult.ID, Valid: true}
	//user.PushToken = sql.NullString{String: oneSignalInsertUserResult.ID, Valid: true}
	//updatePushTokenError := user.UpdateUserPushTokenByUsername(db)
	//if updatePushTokenError != nil {
	//	fmt.Println(updatePushTokenError)
	//}
	//
	////u.WriteLog(username)
	////u.WriteLog(user.Rol)
	//
	//if deviceType == 2 && user.Rol == "Master" || user.Rol == "Propietario" {
	//	return user, nil
	//}
	//if (deviceType == 0 || deviceType == 1) && user.Rol == "Cliente" {
	//	return user, nil
	//} else {
	//	return nil, nil
	//}
	return nil,nil
}

func (o *User) GetUserToken(db *sql.DB) error {
	//statement := fmt.Sprintf("SELECT id FROM users WHERE `token`= '%s' AND `available` = 1", o.Token.String)
	//err := db.QueryRow(statement).Scan(&o.ID)
	//if err != nil {
	//	return err
	//}
	//return nil
	return nil
}

func (o *User) PasswordRecovery(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT id FROM users WHERE `username`= '%s'  AND `available` = 1", o.Username)
	err := db.QueryRow(statement).Scan(&o.ID)
	if err != nil {
		return err
	}

	apiTokenString := u.GenerateTokenUsername(o.Username)

	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	statement1 := fmt.Sprintf(
		"INSERT INTO users_recovery (user_id, token, available, updated_at, created_at) "+
			"VALUES('%d', '%s', '%d', '%s', '%s')",
		o.ID, apiTokenString, 1, date, date)
	_, err = db.Exec(statement1)



	d := u.Info{Data: apiTokenString}
	d.SendMailRecovery(o.Username)

	return err
}

func (o *UserRecovery) CheckPassword(db *sql.DB) bool {
	statement := fmt.Sprintf("SELECT id,user_id FROM users_recovery WHERE `token`= '%s'  AND `available` = 1", o.Token)
	db.QueryRow(statement).Scan(&o.ID, &o.UserID)
	if o.ID != 0 {
		return true
	} else {
		return false
	}
}

func (o *UserRecovery) PasswordRecoverySecond(db *sql.DB) error {
	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))

	find_user_query := fmt.Sprintf(
		"SELECT id,user_id FROM users_recovery WHERE `token`= '%s'  AND `available` = 1", o.Token)
	db.QueryRow(find_user_query).Scan(&o.ID, &o.UserID)
	fmt.Println(o)

	var username string

	find_username_query := fmt.Sprintf(
		"SELECT username FROM users WHERE `id`= '%d'  AND `available` = 1", o.UserID)
	db.QueryRow(find_username_query).Scan(&username)
	update_user_password := fmt.Sprintf("UPDATE users SET password='%s', updated_at='%s' WHERE id='%d'",
		hashAndSalt([]byte(o.Password)), date, o.UserID)
	_, err := db.Exec(update_user_password)

	disable_recovery := fmt.Sprintf("UPDATE users_recovery SET available='0', updated_at='%s' WHERE id='%d'", date, o.ID)
	db.Exec(disable_recovery)
	d := u.Info{Data: o.Password}
	d.SendPasswordRecoveryConfirmation(username)
	return err
}

func (o *User) ChangePassword(db *sql.DB) error {
	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	statement := fmt.Sprintf("UPDATE users SET password='%s', updated_at='%s' WHERE username='%s'",
		hashAndSalt([]byte(o.Password)), date, o.Username)
	_, err := db.Exec(statement)
	return err
}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
