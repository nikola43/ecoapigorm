package models

import (
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	u "github.com/nikola43/ecoapigorm/utils"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type MultimediaContentSize struct {
	Size uint `json:"size"`
}
type MultimediaContent struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Url       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
	PregnancyID    uint   `json:"pregnancy_id"`
}

type VideoThumbnail struct {
	ID        uint   `json:"id"`
	VideoID   uint   `json:"video_id"`
	Thumbnail string `json:"thumbnail"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
}

func (o *MultimediaContent) InsertMultimediaContent(db *sql.DB) (sql.Result, error) {
	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	statement := fmt.Sprintf(
		"INSERT INTO multimedia_contents (user_id, name, type, url, updated_at, created_at) "+
			"VALUES('%d', '%s', '%s', '%s', '%s', '%s')",
		o.UserID, o.Name, o.Type, o.Url, date, date)

	res, err := db.Exec(statement)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (o *VideoThumbnail) InsertVideoThumbnail(db *sql.DB) error {
	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	statement := fmt.Sprintf(
		"INSERT INTO video_thumbnails (video_id, thumbnail, updated_at, created_at) "+
			"VALUES('%d', '%s', '%s', '%s')",
		o.VideoID, o.Thumbnail, date, date)

	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	return nil
}

func (o *MultimediaContent) GetRangeImagesByUserID(db *sql.DB, start, end uint) ([]MultimediaContent, error) {
	var list []MultimediaContent
	statement := fmt.Sprintf("SELECT * FROM multimedia_contents WHERE user_id = %d AND id BETWEEN %d AND %d AND type = 'image' AND deleted_at = '2000-01-01 01:01:01'", o.UserID, start, end)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var o MultimediaContent
		if err := rows.Scan(&o.ID, &o.UserID, &o.Name, &o.Type, &o.Url, &o.UpdatedAt, &o.CreatedAt, &o.DeletedAt); err != nil {
			return nil, err
		}
		list = append(list, o)
	}

	if len(list) == 0 {
		list := make([]MultimediaContent, 0)
		return list, nil
	}

	return list, nil
}

func (o *MultimediaContent) GetImagesByUserID(db *sql.DB) ([]MultimediaContent, error) {
	var list []MultimediaContent
	statement := fmt.Sprintf("SELECT * FROM multimedia_contents WHERE user_id = %d AND type = 'image' AND deleted_at = '2000-01-01 01:01:01'", o.UserID)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var o MultimediaContent
		if err := rows.Scan(&o.ID, &o.UserID, &o.Name, &o.Type, &o.Url, &o.UpdatedAt, &o.CreatedAt, &o.DeletedAt); err != nil {
			return nil, err
		}
		list = append(list, o)
	}

	if len(list) == 0 {
		list := make([]MultimediaContent, 0)
		return list, nil
	}

	return list, nil
}

func (o *MultimediaContent) GetVideosByUserID(db *sql.DB) ([]MultimediaContent, error) {
	var list []MultimediaContent
	statement := fmt.Sprintf("SELECT M.*, T.thumbnail as thumbnail FROM multimedia_contents M INNER JOIN video_thumbnails T ON M.id = T.video_id WHERE M.user_id = %d AND M.type = 'video' AND M.deleted_at = '2000-01-01 01:01:01' GROUP BY M.id", o.UserID)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var o MultimediaContent
		if err := rows.Scan(&o.ID, &o.UserID, &o.Name, &o.Type, &o.Url, &o.UpdatedAt, &o.CreatedAt, &o.DeletedAt, &o.Thumbnail); err != nil {
			return nil, err
		}
		list = append(list, o)
	}

	if len(list) == 0 {
		list := make([]MultimediaContent, 0)
		return list, nil
	}

	return list, nil
}

func (o *MultimediaContent) GetHolographiesByUserID(db *sql.DB) ([]MultimediaContent, error) {
	var list []MultimediaContent
	statement := fmt.Sprintf("SELECT M.*, T.thumbnail as thumbnail FROM multimedia_contents M INNER JOIN video_thumbnails T ON M.id = T.video_id WHERE M.user_id = %d AND M.type = 'holography' AND M.deleted_at = '2000-01-01 01:01:01' GROUP BY M.id", o.UserID)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var o MultimediaContent
		if err := rows.Scan(&o.ID, &o.UserID, &o.Name, &o.Type, &o.Url, &o.UpdatedAt, &o.CreatedAt, &o.DeletedAt, &o.Thumbnail); err != nil {
			return nil, err
		}
		list = append(list, o)
	}

	if len(list) == 0 {
		list := make([]MultimediaContent, 0)
		return list, nil
	}

	return list, nil
}

func (o *MultimediaContent) getMultimediaContent(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT * FROM multimedia_contents WHERE id=%d AND deleted_at = '1000-01-01 00:00:00'", o.ID)
	return db.QueryRow(statement).Scan(&o.ID, &o.UserID, &o.Name, &o.Type, &o.Url, &o.UpdatedAt, &o.CreatedAt, &o.DeletedAt)
}

func (o *MultimediaContent) updateMultimediaContent(db *sql.DB) error {
	date := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	statement := fmt.Sprintf(
		"UPDATE multimedia_contents "+
			"SET user_id='%d', name='%s', type='%s', url='%s', updated_at='%s' "+
			"WHERE id=%d",
		o.UserID, o.Name, o.Type, o.Url, date, o.ID)
	_, err := db.Exec(statement)
	return err
}

func (o *MultimediaContent) DeleteMultimediaContent(db *sql.DB, S3Session *s3.S3, awsBucketName string) error {

	multimedia := MultimediaContent{}
	statement := fmt.Sprintf("SELECT * FROM multimedia_contents WHERE id = %d", o.ID)
	deletePath := ""
	err := db.QueryRow(statement).Scan(&multimedia.ID, &multimedia.UserID, &multimedia.Name, &multimedia.Type, &multimedia.Url, &multimedia.UpdatedAt, &multimedia.CreatedAt, &multimedia.DeletedAt)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(multimedia.Name) > 0 {
		s := strconv.FormatUint(uint64(multimedia.UserID), 10)

		if multimedia.Type == "video" || multimedia.Type == "holography" {
			// delete thumb
			deletePath = s + "/" + multimedia.Type + "/" + multimedia.Name + "-thumbnail.jpg"
			err = u.DeleteObject(S3Session, awsBucketName, aws.String(deletePath))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(deletePath)

			statement = fmt.Sprintf("DELETE FROM video_thumbnails WHERE video_id=%d", multimedia.ID)
			res, err := db.Exec(statement)
			count, err := res.RowsAffected()
			fmt.Println(count)

			if err != nil {
				fmt.Println(err)
			}

			// delete video
			deletePath = s + "/" + multimedia.Type + "/" + multimedia.Name
			err = u.DeleteObject(S3Session, awsBucketName, aws.String(deletePath))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(deletePath)

			statement = fmt.Sprintf("DELETE FROM multimedia_contents WHERE id=%d", o.ID)
			res2, err := db.Exec(statement)
			count2, err := res2.RowsAffected()
			fmt.Println(count2)
			fmt.Println(count2)

			if err != nil {
				fmt.Println(err)
			}

		} else {
			deletePath = s + "/" + multimedia.Type + "/" + multimedia.Name
			err = u.DeleteObject(S3Session, awsBucketName, aws.String(deletePath))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(deletePath)

			statement = fmt.Sprintf("DELETE FROM multimedia_contents WHERE id=%d", o.ID)
			res, err := db.Exec(statement)
			count, err := res.RowsAffected()

			if err != nil {
				fmt.Println(err)
			}
			if count == 0 {
				return sql.ErrNoRows
			}
			fmt.Println(count)
		}
	}
	return err
}

func getMultimediaContents(db *sql.DB) ([]MultimediaContent, error) {
	var list []MultimediaContent

	rows, err := db.Query("SELECT * FROM multimedia_contents WHERE deleted_at = '2000-01-01 01:01:01'")

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var o MultimediaContent
		if err := rows.Scan(&o.ID, &o.UserID, &o.Name, &o.Type, &o.Url, &o.UpdatedAt, &o.CreatedAt, &o.DeletedAt); err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, nil
}

func (o *MultimediaContent) getMultimediaContentPhotoByUserID(db *sql.DB) ([]MultimediaContent, error) {
	var list []MultimediaContent
	statement := fmt.Sprintf("SELECT * FROM multimedia_contents WHERE user_id=%d AND type='Imagen'  AND deleted_at = '2000-01-01 01:01:01'", o.UserID)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var o MultimediaContent
		if err := rows.Scan(&o.ID, &o.UserID, &o.Name, &o.Type, &o.Url, &o.UpdatedAt, &o.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, nil
}

func (o *MultimediaContent) getMultimediaContentVideoByUserID(db *sql.DB) ([]MultimediaContent, error) {
	var list []MultimediaContent
	statement := fmt.Sprintf("SELECT * FROM multimedia_contents WHERE user_id=%d AND type='Video'  AND deleted_at = '2000-01-01 01:01:01'", o.UserID)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var o MultimediaContent
		if err := rows.Scan(&o.ID, &o.UserID, &o.Name, &o.Type, &o.Url, &o.UpdatedAt, &o.CreatedAt, &o.DeletedAt); err != nil {
			return nil, err
		}
		list = append(list, o)
	}

	if len(list) == 0 {
		return list, sql.ErrNoRows
	}

	return list, nil
}
