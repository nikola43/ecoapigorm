package services

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/wasabis3manager"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/utils"
	"log"
	"mime/multipart"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func UploadMultimedia(
	context *fiber.Ctx,
	bucketName string,
	clinicName string,
	clientID uint,
	uploadedFile *multipart.FileHeader,
	uploadMode uint,
	clinicId uint) error {

	// Save file to root directory:
	reg, err := regexp.Compile("[^a-zA-Z0-9-.]+")
	if err != nil {
		log.Fatal(err)
	}
	cleanFilename := reg.ReplaceAllString(uploadedFile.Filename, "")
	clinicName = reg.ReplaceAllString(clinicName, "")
	clinicName = strings.ToLower(strings.ReplaceAll(clinicName, " ", ""))
	var clientIDString = strconv.FormatUint(uint64(clientID), 10)
	os.MkdirAll(fmt.Sprintf("./tempFiles/%s/%s", clinicName, clientIDString), os.ModePerm)

	err = context.SaveFile(uploadedFile, fmt.Sprintf("./tempFiles/%s/%s/%s", clinicName, clientIDString, cleanFilename))
	if err != nil {
		return err
	}

	fileType := utils.GetFileType(cleanFilename, uploadMode)
	fmt.Println("fileType")
	fmt.Println(fileType)

	if fileType == "image" {
		// image

		url, size, storeInAmazonError := wasabis3manager.WasabiS3Client.UploadObject(
			bucketName,
			clinicName,
			"tempFiles/"+clinicName+"/"+clientIDString+"/"+cleanFilename,
			clientID,
			fileType,
		)
		if storeInAmazonError != nil {
			fmt.Println(storeInAmazonError.Error())
		}
		image := models.Image{
			Filename: cleanFilename,
			ClientID: clientID,
			Url:      url,
			Size:     uint(size),
			ClinicID: clinicId,
		}
		database.GormDB.Create(&image)

		e := os.Remove(fmt.Sprintf("./tempFiles/%s/%s/%s", clinicName, clientIDString, cleanFilename))
		if e != nil {
			fmt.Println(e)
			//log.Fatal(e)
		}

		return err
	} else if fileType == "video" || fileType == "holographic" {
		// upload video
		videoUrl, videoSize, storeInAmazonError := wasabis3manager.WasabiS3Client.UploadObject(
			bucketName,
			clinicName,
			"tempFiles/"+clinicName+"/"+clientIDString+"/"+cleanFilename,
			clientID,
			fileType,
		)
		if storeInAmazonError != nil {
			return err
		}

		ThumbnailCleanFilename := "tempFiles/" + clinicName + "/" + clientIDString + "/" + cleanFilename + "_thumbnail.jpg"
		ThumbnailCleanFilename = strings.ToLower(strings.ReplaceAll(ThumbnailCleanFilename, " ", ""))

		// create thumb
		thumbnailPath, extractThumbnailFromVideoError := utils.ExtractThumbnailFromVideo("tempFiles/" + clinicName + "/" + clientIDString + "/" + cleanFilename)
		if extractThumbnailFromVideoError != nil {
			return extractThumbnailFromVideoError
		}

		// thumb video
		thumbUrl, thumbSize, storeThumbInAmazonError := wasabis3manager.WasabiS3Client.UploadObject(
			bucketName,
			clinicName,
			thumbnailPath,
			clientID,
			fileType,
		)
		if storeThumbInAmazonError != nil {
			return storeThumbInAmazonError
		}

		if fileType == "video" {
			video := models.Video{
				Filename:     cleanFilename,
				ClientID:     clientID,
				Url:          videoUrl,
				ThumbnailUrl: thumbUrl,
				Size:         uint(videoSize + thumbSize),
				ClinicID:     clinicId,
			}
			database.GormDB.Create(&video)
		}

		if fileType == "holographic" {
			video := models.Holographic{
				Filename:     cleanFilename,
				ClientID:     clientID,
				Url:          videoUrl,
				ThumbnailUrl: thumbUrl,
				Size:         uint(videoSize + thumbSize),
				ClinicID:     clinicId,
			}
			database.GormDB.Create(&video)
		}

		e := os.Remove(fmt.Sprintf("./tempFiles/%s/%s/%s", clinicName, clientIDString, cleanFilename))
		if e != nil {
			fmt.Println(e)
		}

		e = os.Remove(thumbnailPath)
		if e != nil {
			fmt.Println(e)
		}

		return err
	} else if fileType == "holo" {
		// holo
		/*
			url, size, storeInAmazonError := utils.UploadObject("tempFiles/"+cleanFilename, clientID, fileType)
			if storeInAmazonError != nil {
				fmt.Println(storeInAmazonError.Error())
			}
			image := models.Video{ClientID: clientID, Url: url, Size: uint(size)}
			database.GormDB.Create(&image)
		*/
		return err
	} else if fileType == "heartbeat" {
		// holo
		url, size, storeInAmazonError := wasabis3manager.WasabiS3Client.UploadObject(bucketName, clinicName, "tempFiles/"+clinicName+"/"+clientIDString+"/"+cleanFilename, clientID, fileType)
		if storeInAmazonError != nil {
			fmt.Println(storeInAmazonError.Error())
		}
		video := models.Heartbeat{Filename: cleanFilename, ClientID: clientID, Url: url, Size: uint(size), ClinicID: clinicId}
		database.GormDB.Create(&video)
		e := os.Remove(fmt.Sprintf("./tempFiles/%s/%s/%s", clinicName, clientIDString, cleanFilename))

		if e != nil {
			fmt.Println(e)
		}

		return err
	} else {
		return errors.New("invalid file")
	}
}

func DeleteImage(bucketName string, imageID uint) error {
	image := &models.Image{}

	result := database.GormDB.First(&image, imageID)
	if result.Error != nil {
		return result.Error
	}

	url := strings.Split(image.Url, "/")
	key := url[4] + "/" + url[5] + "/" + url[6]

	// delete from S3
	err := wasabis3manager.WasabiS3Client.DeleteObject(bucketName, &key)
	if err != nil {
		return err
	}

	// delete from DB
	result = database.GormDB.Unscoped().Delete(image)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteVideo(bucketName string, imageID uint) error {
	video := &models.Video{}

	result := database.GormDB.First(&video, imageID)
	if result.Error != nil {
		return result.Error
	}

	videoUrl := strings.Split(video.Url, "/")
	thumbnailUrl := strings.Split(video.ThumbnailUrl, "/")
	videoKey := videoUrl[4] + "/" + videoUrl[5] + "/" + videoUrl[6]
	thumbnailKey := thumbnailUrl[4] + "/" + thumbnailUrl[5] + "/" + thumbnailUrl[6]

	// delete video
	err := wasabis3manager.WasabiS3Client.DeleteObject(bucketName, &videoKey)
	if err != nil {
		return err
	}

	// delete thumbnail
	err = wasabis3manager.WasabiS3Client.DeleteObject(bucketName, &thumbnailKey)
	if err != nil {
		return err
	}

	// delete from DB
	result = database.GormDB.Unscoped().Delete(video)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteHeartbeat(bucketName string, imageID uint) error {
	heartbeat := &models.Heartbeat{}

	result := database.GormDB.First(&heartbeat, imageID)
	if result.Error != nil {
		return result.Error
	}

	url := strings.Split(heartbeat.Url, "/")
	key := url[4] + "/" + url[5] + "/" + url[6]

	// delete from S3
	err := wasabis3manager.WasabiS3Client.DeleteObject(bucketName, &key)
	if err != nil {
		return err
	}

	// delete from DB
	result = database.GormDB.Unscoped().Delete(heartbeat)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
