package models

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Clinic struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	ClinicOwner string `json:"clinic_owner"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Users       uint   `json:"users"`
	Available   uint   `json:"available"`
	Employers []Employer `gorm:"many2many:employer_clinics;"`
}

func (o *Clinic) InsertClinic(db *sql.DB) (*Clinic, error) {
	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	id := sql.NullInt64{Int64: 0, Valid: true}
	_, idError := id.Value()
	if idError != nil {
		return nil, idError
	}
	statement := fmt.Sprintf(
		"INSERT INTO clinics (user_id, name, address, available, updated_at, created_at) "+
			"VALUES('%d', '%s', '%s', '%d', '%s', '%s')",
		o.UserID, o.Name, o.Address, 1, date, date)

	res, err := db.Exec(statement)
	if err != nil {
		return nil, err
	}

	insertedId, err := res.LastInsertId()
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	statement = fmt.Sprintf("UPDATE users SET clinic_id=%d, updated_at='%s' "+"WHERE id=%d", insertedId, date, o.UserID)
	_, err = db.Exec(statement)
	o.ID = uint(insertedId)
	_ = o.GetClinic(db)
	return o, nil
}

func (o *Clinic) GetClinic(db *sql.DB) error {
	statement := fmt.Sprintf(
		"SELECT c.id, c.user_id, u.username as clinic_owner, c.name, c.address, c.available, c.updated_at, c.created_at "+
			"FROM clinics as c "+
			"INNER JOIN users as u "+
			"ON c.user_id = u.id "+
			"WHERE c.id = %d AND c.available = 1 ORDER BY c.created_at DESC", o.ID)
	return db.QueryRow(statement).Scan(&o.ID, &o.UserID, &o.ClinicOwner, &o.Name, &o.Address, &o.Available, &o.UpdatedAt, &o.CreatedAt)
}

func (o *Clinic) GetClinicByUserID(db *sql.DB) error {
	statement := fmt.Sprintf(
		"SELECT c.id, c.user_id, CONCAT(u.name , ' ', u.lastname) as clinic_owner, c.name, c.address, c.available, c.updated_at, c.created_at "+
			"FROM clinics as c "+
			"INNER JOIN users as u "+
			"ON c.user_id = u.id "+
			"WHERE user.id = %d AND c.available = 1 ORDER BY c.created_at DESC", o.UserID)
	return db.QueryRow(statement).Scan(&o.ID, &o.UserID, &o.ClinicOwner, &o.Name, &o.Address, &o.Available, &o.UpdatedAt, &o.CreatedAt)
}

func (o *Clinic) UpdateClinic(db *sql.DB) (*Clinic, error) {
	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	statement := fmt.Sprintf(
		"UPDATE clinics "+
			"SET name='%s', user_id='%d', address='%s', updated_at='%s' "+
			"WHERE id=%d",
		o.Name, o.UserID, o.Address, date, o.ID)
	_, err := db.Exec(statement)
	fmt.Println(statement)
	fmt.Println(err)

	statement = fmt.Sprintf("UPDATE users SET clinic_id=%d, updated_at='%s' "+"WHERE id=%d", o.ID, date, o.UserID)
	_, err = db.Exec(statement)
	fmt.Println(err)

	_ = o.GetClinic(db)

	return o, err
}

func (o *Clinic) DeleteClinic(db *sql.DB) error {
	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	statement := fmt.Sprintf(
		"UPDATE clinics "+
			"SET available='%d', updated_at='%s' "+
			"WHERE id=%d",
		0, date, o.ID)
	_, err := db.Exec(statement)
	return err
}

/*
DROP FUNCTION IF EXISTS F_COUNT_USERS_BY_CLINIC;


DELIMITER $$

CREATE FUNCTION F_COUNT_USERS_BY_CLINIC(clinic_id INTEGER) RETURNS INT
BEGIN
  DECLARE usersByClinic INT unsigned DEFAULT 0;

  SELECT COUNT(u.id)
  		INTO usersByClinic
        FROM clinics as c
        INNER JOIN users as u
        ON c.id = u.clinic_id
        AND c.available = 1
        AND u.available = 1
        AND c.id = clinic_id;
  RETURN usersByClinic;
END $$

DELIMITER ;

SELECT F_COUNT_USERS_BY_CLINIC(1) FROM `clinics`
*/

func (o *Clinic) GetClinics(db *sql.DB) ([]Clinic, error) {
	statement := fmt.Sprintf(
		"SELECT cl.id, cl.user_id, u.username as clinic_owner, cl.name, F_COUNT_USERS_BY_CLINIC(cl.id) as users, cl.address, cl.available, cl.updated_at, cl.created_at" +
			" FROM clinics as cl" +
			" INNER JOIN users as u" +
			" ON cl.user_id = u.id" +
			" AND cl.available = 1" +
			" GROUP BY cl.id" +
			" ORDER BY cl.created_at DESC")

	var list []Clinic

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var o Clinic
		if err := rows.Scan(&o.ID, &o.UserID, &o.ClinicOwner, &o.Name, &o.Users, &o.Address, &o.Available, &o.UpdatedAt, &o.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, o)
	}

	if len(list) == 0 {
		list := make([]Clinic, 0)
		return list, nil
	}

	return list, nil
}

func GetUsersByClinicID(db *sql.DB, id uint) ([]User, error) {
	//var list []User
	//statement := fmt.Sprintf(
	//	"SELECT u.`id`, u.`clinic_id`, c.`name`, u.`username`, u.`password`, u.`name`, u.`lastname`,"+
	//		" u.`phone`, u.`rol`, u.`available`, u.`token`, u.`updated_at`, u.`created_at`"+
	//		" FROM `users` AS u "+
	//		"INNER JOIN clinics AS c ON u.`clinic_id` = c.`id` "+
	//		"WHERE u.`id` = %d AND u.`available` = 1 ORDER BY c.created_at DESC", id)
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
	//	if err := rows.Scan(&o.ID, &o.ClinicID, &o.Username, &o.Password, &o.Name, &o.Lastname, &o.Phone, &o.Rol, &o.Available, &o.FirebaseToken, &o.UpdatedAt, &o.CreatedAt); err != nil {
	//		return nil, err
	//	}
	//	list = append(list, o)
	//}
	//return list, nil
	return nil,nil
}
