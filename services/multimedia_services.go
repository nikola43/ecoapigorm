package services

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/awsmanager"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/utils"
	"log"
	"mime/multipart"
	"os"
)

func UploadMultimedia(context *fiber.Ctx, bucketName string, clientID uint, uploadedFile *multipart.FileHeader, uploadMode uint) error {
	//fmt.Println(context)
	//fmt.Println(clientID)
	//fmt.Println(uploadedFile)

	// Save file to root directory:
	err := context.SaveFile(uploadedFile, fmt.Sprintf("./tempFiles/%s", uploadedFile.Filename))
	if err != nil {
		return err
	}

	fileType := utils.GetFileType(uploadedFile.Filename, uploadMode)
	fmt.Println("fileType")
	fmt.Println(fileType)

	if fileType == "image" {
		// image

		url, size, storeInAmazonError := awsmanager.AwsManager.UploadObject(bucketName, "tempFiles/"+uploadedFile.Filename, clientID, fileType)
		if storeInAmazonError != nil {
			fmt.Println(storeInAmazonError.Error())
		}
		image := models.Image{ClientID: clientID, Url: url, Size: uint(size)}
		database.GormDB.Create(&image)

		return err
	} else if fileType == "video" || fileType == "holographic" {
		// upload video
		videoUrl, videoSize, storeInAmazonError := awsmanager.AwsManager.UploadObject(bucketName,"tempFiles/"+uploadedFile.Filename, clientID, fileType)
		if storeInAmazonError != nil {
			return err
		}

		// create thumb
		thumbnailPath, extractThumbnailFromVideoError := utils.ExtractThumbnailFromVideo("tempFiles/"+uploadedFile.Filename)
		if extractThumbnailFromVideoError != nil {
			return extractThumbnailFromVideoError
		}

		// thumb video
		thumbUrl, thumbSize, storeThumbInAmazonError := awsmanager.AwsManager.UploadObject(bucketName,thumbnailPath, clientID, fileType)
		if storeThumbInAmazonError != nil {
			return storeThumbInAmazonError
		}

		if fileType == "video" {
			video := models.Video{ClientID: clientID, Url: videoUrl, ThumbnailUrl: thumbUrl, Size: uint(videoSize + thumbSize)}
			database.GormDB.Create(&video)
		}

		if fileType == "holographic" {
			video := models.Holographic{ClientID: clientID, Url: videoUrl, ThumbnailUrl: thumbUrl, Size: uint(videoSize + thumbSize)}
			database.GormDB.Create(&video)
		}

		e := os.Remove(thumbnailPath)
		if e != nil {
			log.Fatal(e)
		}

		return err
	} else if fileType == "holo" {
		// holo
		/*
			url, size, storeInAmazonError := utils.UploadObject("tempFiles/"+uploadedFile.Filename, clientID, fileType)
			if storeInAmazonError != nil {
				fmt.Println(storeInAmazonError.Error())
			}
			image := models.Video{ClientID: clientID, Url: url, Size: uint(size)}
			database.GormDB.Create(&image)
		*/
		return err
	} else if fileType == "heartbeat" {
		// holo
		url, size, storeInAmazonError := awsmanager.AwsManager.UploadObject(bucketName,"tempFiles/"+uploadedFile.Filename, clientID, fileType)
		if storeInAmazonError != nil {
			fmt.Println(storeInAmazonError.Error())
		}
		video := models.Heartbeat{ClientID: clientID, Url: url, Size: uint(size)}
		database.GormDB.Create(&video)
		return err
	} else {
		return errors.New("invalid file")
	}
}

func DeleteImage(imageID uint) error {
	image := &models.Image{}

	result := database.GormDB.First(&image, imageID)
	if result.Error != nil {
		return result.Error
	}

	// delete from S3

	// delete from DB
	result = database.GormDB.Unscoped().Delete(image)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
