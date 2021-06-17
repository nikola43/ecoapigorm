package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/models/promos"
	"github.com/nikola43/ecoapigorm/socketinstance"
	"github.com/nikola43/ecoapigorm/utils"
	"github.com/nikola43/ecoapigorm/wasabis3manager"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func UploadPromoImage(
	context *fiber.Ctx,
	bucketName string,
	uploadedFile *multipart.FileHeader,
	clinicName string,
	promo *promos.Promo,
) error {
	// sanitize file name
	reg, _ := regexp.Compile("[^a-zA-Z0-9-.]+")
	cleanFilename := reg.ReplaceAllString(uploadedFile.Filename, "")
	clinicName = reg.ReplaceAllString(clinicName, "")
	clinicName = strings.ToLower(strings.ReplaceAll(clinicName, " ", ""))

	//uploadedFilePath := ""
	promosFolder := "./tempFiles/" + clinicName + "/" + "promos" + "/" + "image" + "/"

	if _, err := os.Stat(promosFolder); os.IsNotExist(err) {
		os.MkdirAll(promosFolder, os.ModePerm)
	}

	err := context.SaveFile(uploadedFile, promosFolder+"/"+cleanFilename)
	if err != nil {
		return err
	}

	url, _, storeInAmazonError := wasabis3manager.WasabiS3Client.UploadObject(
		bucketName,
		clinicName,
		"tempFiles/"+clinicName+"/"+"promos"+"/"+"image"+"/"+cleanFilename,
		"promos",
		"image",
	)
	if storeInAmazonError != nil {
		fmt.Println(storeInAmazonError.Error())
	}

	database.GormDB.Model(&promo).Where("id = ?", promo.ID).Update("image_url", url)

	return nil
}


func UploadMultimedia(
	context *fiber.Ctx,
	bucketName string,
	clinicName string,
	clientID uint,
	uploadedFile *multipart.FileHeader,
	uploadMode uint,
	clinicId uint) error {

	var insertedID uint

	// sanitize file name
	reg, _ := regexp.Compile("[^a-zA-Z0-9-.]+")
	cleanFilename := reg.ReplaceAllString(uploadedFile.Filename, "")
	clinicName = reg.ReplaceAllString(clinicName, "")
	clinicName = strings.ToLower(strings.ReplaceAll(clinicName, " ", ""))

	clientIDString := strconv.FormatUint(uint64(clientID), 10)
	fileType := utils.GetFileType(cleanFilename, uploadMode)

	if strings.Contains(cleanFilename, ".avi") {
		cleanFilename = cleanFilename + ".mp4"
	}

	url := "https://s3.eu-central-1.wasabisys.com/steleros/" + clinicName + "/" + clientIDString + "/" + fileType + "/" + cleanFilename

	//uploadedFilePath := ""
	clientFolder := "./tempFiles/" + clinicName + "/" + clientIDString

	if _, err := os.Stat(clientFolder); os.IsNotExist(err) {
		os.MkdirAll(clientFolder, os.ModePerm)
	}

	e := os.Remove(fmt.Sprintf(clientFolder + "/" + cleanFilename))
	if e != nil {
		fmt.Println(e)
		//panic(e)
	}

	err := context.SaveFile(uploadedFile, clientFolder+"/"+cleanFilename)
	if err != nil {
		return err
	}

	file, err := os.Open(clientFolder + "/" + cleanFilename)
	if err != nil {
		fmt.Printf("err opening file: %s", err)
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()

	fmt.Println("fileType")
	fmt.Println(fileType)

	if _, err := os.Stat(clientFolder + "/" + fileType); os.IsNotExist(err) {
		os.MkdirAll(clientFolder+"/"+fileType, os.ModePerm)
	}

	switch fileType {
	case "image":

		// image
		image := models.Image{
			Filename:  cleanFilename,
			ClientID:  clientID,
			Url:       url,
			Size:      uint(size),
			ClinicID:  clinicId,
			Available: false,
		}
		database.GormDB.Create(&image)

		insertedID = image.ID
		fmt.Println(insertedID)

		go func() {
			/*
				err = utils.CompressImage("tempFiles/"+clinicName+"/"+clientIDString+"/"+cleanFilename, "tempFiles/"+clinicName+"/"+clientIDString+"/"+fileType+"/"+cleanFilename)
				if err != nil {
					fmt.Println(err.Error())
				}
			*/

			input, err := ioutil.ReadFile("tempFiles/" + clinicName + "/" + clientIDString + "/" + cleanFilename)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = ioutil.WriteFile("tempFiles/"+clinicName+"/"+clientIDString+"/"+fileType+"/"+cleanFilename, input, 0644)
			if err != nil {
				fmt.Println("Error creating", "tempFiles/"+clinicName+"/"+clientIDString+"/"+fileType+"/"+cleanFilename)
				fmt.Println(err)
				return
			}

			imageUrl, imageSize, storeInAmazonError := wasabis3manager.WasabiS3Client.UploadObject(
				bucketName,
				clinicName,
				//"tempFiles/"+clinicName+"/"+clientIDString+"/"+fileType+"/"+cleanFilename,
				"tempFiles/"+clinicName+"/"+clientIDString+"/"+fileType+"/"+cleanFilename,
				strconv.FormatInt(int64(clientID), 10),
				fileType,
			)
			if storeInAmazonError != nil {
				fmt.Println(storeInAmazonError.Error())
			}

			imageUpdate := new(models.Image)
			imageUpdate.ID = insertedID
			result := database.GormDB.Where("id = ?", imageUrl).Find(&imageUpdate)
			if result.Error != nil {
				log.Fatal(result.Error)
				return
			}

			database.GormDB.Model(&imageUpdate).Where("id = ?", imageUpdate.ID).Update("available", true)
			imageUpdate.Available = true
			fmt.Println(imageUrl)
			fmt.Println(imageSize)

			e := os.Remove("tempFiles/" + clinicName + "/" + clientIDString + "/" + cleanFilename)
			if e != nil {
				fmt.Println(e)
				//panic(e)
			}

			e = os.Remove(fmt.Sprintf("tempFiles/" + clinicName + "/" + clientIDString + "/" + fileType + "/" + cleanFilename))
			if e != nil {
				fmt.Println(e)
				//panic(e)
			}

			socketEvent := models.SocketEvent{
				Type:   "image",
				Action: "update",
				Data:   imageUpdate,
			}

			b, _ := json.Marshal(socketEvent)
			socketinstance.SocketInstance.Emit(b)
		}()

		return err
		break
	case "video":

		ThumbnailCleanFilename := "tempFiles/" + clinicName + "/" + clientIDString + "/" + cleanFilename + "_thumbnail.jpg"
		ThumbnailCleanFilename = strings.ToLower(strings.ReplaceAll(ThumbnailCleanFilename, " ", ""))

		// create thumb
		extractThumbnailFromVideoError := utils.ExtractThumbnailFromVideo("tempFiles/"+clinicName+"/"+clientIDString+"/"+cleanFilename, "tempFiles/"+clinicName+"/"+clientIDString+"/"+fileType+"/"+cleanFilename+"_thumbnail.jpg")
		if extractThumbnailFromVideoError != nil {
			return extractThumbnailFromVideoError
		}

		// thumb video
		thumbUrl, thumbSize, storeThumbInAmazonError := wasabis3manager.WasabiS3Client.UploadObject(
			bucketName,
			clinicName,
			"tempFiles/"+clinicName+"/"+clientIDString+"/"+fileType+"/"+cleanFilename+"_thumbnail.jpg",
			strconv.FormatInt(int64(clientID), 10),
			fileType,
		)
		if storeThumbInAmazonError != nil {
			return storeThumbInAmazonError
		}

		if fileType == "video" {
			video := models.Video{
				Filename:     cleanFilename,
				ClientID:     clientID,
				Url:          url,
				ThumbnailUrl: thumbUrl,
				Size:         uint(size + thumbSize),
				ClinicID:     clinicId,
				Available:    false,
			}
			database.GormDB.Create(&video)

			insertedID = video.ID
			fmt.Println(insertedID)
		}

		if fileType == "holographic" {
			video := models.Holographic{
				Filename:     cleanFilename,
				ClientID:     clientID,
				Url:          url,
				ThumbnailUrl: thumbUrl,
				Size:         uint(size + thumbSize),
				ClinicID:     clinicId,
			}
			database.GormDB.Create(&video)
		}

		go func() {
			fmt.Println("fileComrpess")
			err := utils.CompressMP4("tempFiles/"+clinicName+"/"+clientIDString+"/"+cleanFilename, "tempFiles/"+clinicName+"/"+clientIDString+"/"+fileType+"/"+cleanFilename)
			if err != nil {
				log.Fatal(err)
				return
			}

			// upload video
			videoUrl, videoSize, storeInAmazonError := wasabis3manager.WasabiS3Client.UploadObject(
				bucketName,
				clinicName,
				"tempFiles/"+clinicName+"/"+clientIDString+"/"+fileType+"/"+cleanFilename,
				strconv.FormatInt(int64(clientID), 10),
				fileType,
			)
			if storeInAmazonError != nil {
				log.Fatal(storeInAmazonError)
				return
			}

			fmt.Println(videoUrl)
			fmt.Println(videoSize)

			videoUpdate := new(models.Video)
			videoUpdate.ID = insertedID
			result := database.GormDB.Where("id = ?", videoUrl).Find(&videoUpdate)
			if result.Error != nil {
				log.Fatal(result.Error)
				return
			}

			database.GormDB.Model(&videoUpdate).Where("id = ?", videoUpdate.ID).Update("available", true)
			videoUpdate.Available = true

			fmt.Println("fin")

			/*
				e := os.Remove(fmt.Sprintf("./tempFiles/%s/%s/%s", clinicName, clientIDString, cleanFilename))
				if e != nil {
					fmt.Println(e)
				}

				e = os.Remove("tempFiles/"+clinicName+"/"+clientIDString+"/"+fileType+"/"+cleanFilename+"_thumbnail.jpg")
				if e != nil {
					fmt.Println(e)
				}
			*/

			socketEvent := models.SocketEvent{
				Type:   "video",
				Action: "update",
				Data:   videoUpdate,
			}

			b, _ := json.Marshal(socketEvent)
			socketinstance.SocketInstance.Emit(b)
		}()

		return err
		break
	case "heartbeat":

		// holo
		url, size, storeInAmazonError := wasabis3manager.WasabiS3Client.UploadObject(bucketName, clinicName, "tempFiles/"+clinicName+"/"+clientIDString+"/"+cleanFilename, strconv.FormatInt(int64(clientID), 10), fileType)
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
		break
	}

	return errors.New("invalid file")
}

func DeleteImage(bucketName string, imageID uint) error {
	image := &models.Image{}

	result := database.GormDB.First(&image, imageID)
	if result.Error != nil {
		return result.Error
	}

	url := strings.Split(image.Url, "/")
	key := url[4] + "/" + url[5] + "/" + url[6] + "/" + url[7]

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
	videoKey := videoUrl[4] + "/" + videoUrl[5] + "/" + videoUrl[6] + "/" + videoUrl[7]
	thumbnailKey := thumbnailUrl[4] + "/" + thumbnailUrl[5] + "/" + thumbnailUrl[6] + "/" + thumbnailUrl[7]

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
	key := url[4] + "/" + url[5] + "/" + url[6] + "/" + url[7]

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
