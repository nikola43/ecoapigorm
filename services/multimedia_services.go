package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/utils"
	"github.com/nikola43/ecoapigorm/wasabis3manager"
	"github.com/nikola43/ecoapigorm/websockets"
)

func UploadPromoImage(
	context *fiber.Ctx,
	bucketName string,
	uploadedFile *multipart.FileHeader,
	clinicName string,
	promo *models.Promo,
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
	clinicId uint,
	employeeID uint) error {

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

	if strings.Contains(cleanFilename, ".mpg") {
		cleanFilename = cleanFilename + ".mp4"
	}

	url := "https://s3.eu-central-1.wasabisys.com/ecoxportugal/" + clinicName + "/" + clientIDString + "/" + fileType + "/" + cleanFilename
	fmt.Println(url)
	//uploadedFilePath := ""
	clientFolder := "./tempFiles/" + clinicName + "/" + clientIDString

	if _, err := os.Stat(clientFolder); os.IsNotExist(err) {
		os.MkdirAll(clientFolder, os.ModePerm)
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
			}

			err = ioutil.WriteFile("tempFiles/"+clinicName+"/"+clientIDString+"/"+fileType+"/"+cleanFilename, input, 0644)
			if err != nil {
				fmt.Println("Error creating", "tempFiles/"+clinicName+"/"+clientIDString+"/"+fileType+"/"+cleanFilename)
				fmt.Println(err)
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
			imageUpdate.Filename = image.Filename
			result := database.GormDB.Where("url = ?", imageUrl).Find(&imageUpdate)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			fmt.Println("imageUpdate")
			fmt.Println(imageUpdate)

			database.GormDB.Model(&imageUpdate).Where("id = ?", imageUpdate.ID).Update("available", true)
			imageUpdate.Available = true
			imageUpdate.Url = url
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

			socketEvent := websockets.SocketEvent{
				Type:   "image",
				Action: "update",
				Data:   imageUpdate,
			}

			websockets.Emit(socketEvent, employeeID)
			websockets.Emit(socketEvent, clientID)
			if socketError := recover(); socketError != nil {
				log.Println("panic occurred:", socketError)
			}
		}()

		return err
		break
	case "video":

		var video models.Video
		// var holographic models.Holographic

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
			video = models.Video{
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

			socketEvent := websockets.SocketEvent{
				Type:   "video",
				Action: "insert",
				Data:   video,
			}

			websockets.Emit(socketEvent, employeeID)
			websockets.Emit(socketEvent, clientID)

			if socketError := recover(); socketError != nil {
				log.Println("panic occurred:", socketError)
			}

		}

		if fileType == "holographic" {
			/*
				video = models.Holographic{
					Filename:     cleanFilename,
					ClientID:     clientID,
					Url:          url,
					ThumbnailUrl: thumbUrl,
					Size:         uint(size + thumbSize),
					ClinicID:     clinicId,
				}
				database.GormDB.Create(&video)
			*/
		}

		go func() {
			fmt.Println("fileComrpess")
			err = utils.CompressMP4V2("./tempFiles/"+clinicName+"/"+clientIDString+"/"+cleanFilename,
				"./tempFiles/"+clinicName+"/"+clientIDString+"/"+fileType+"/"+cleanFilename,
				video,
				employeeID,
				clientID,
			)
			if err != nil {
				log.Fatal(err)
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
			}

			fmt.Println(videoUrl)
			fmt.Println(videoSize)

			videoUpdate := new(models.Video)
			videoUpdate.ID = insertedID
			videoUpdate.Filename = video.Filename
			result := database.GormDB.Where("url = ?", videoUrl).Find(&videoUpdate)
			if result.Error != nil {
				log.Fatal(result.Error)
			}

			database.GormDB.Model(&videoUpdate).Where("id = ?", videoUpdate.ID).Update("available", true)
			videoUpdate.Available = true
			videoUpdate.Url = url
			fmt.Println("fin")

			e := os.Remove(fmt.Sprintf("./tempFiles/%s/%s/%s", clinicName, clientIDString, cleanFilename))
			if e != nil {
				fmt.Println(e)
			}

			e = os.Remove("tempFiles/" + clinicName + "/" + clientIDString + "/" + fileType + "/" + cleanFilename + "_thumbnail.jpg")
			if e != nil {
				fmt.Println(e)
			}

			socketEvent := websockets.SocketEvent{
				Type:   "video",
				Action: "update",
				Data:   videoUpdate,
			}

			websockets.Emit(socketEvent, employeeID)
			websockets.Emit(socketEvent, clientID)

			if socketError := recover(); socketError != nil {
				log.Println("panic occurred:", socketError)
			}

		}()

		return err
		break
	case "heartbeat":

		//Read all the contents of the  original file
		bytesRead, err := ioutil.ReadFile("tempFiles/" + clinicName + "/" + clientIDString + "/" + cleanFilename)
		if err != nil {
			log.Fatal(err)
		}

		//Copy all the contents to the desitination file
		err = ioutil.WriteFile("tempFiles/"+clinicName+"/"+clientIDString+"/"+"heartbeat/"+cleanFilename, bytesRead, 0755)
		if err != nil {
			log.Fatal(err)
		}

		err = utils.ConvertAudioToMp4Aac("tempFiles/"+clinicName+"/"+clientIDString+"/"+"heartbeat/"+cleanFilename, "tempFiles/"+clinicName+"/"+clientIDString+"/"+"heartbeat/"+cleanFilename+".mp3")
		if err != nil {
			log.Fatal(err)
		}
		//ffmpeg -i input.wav -ab 192k -acodec libfaac output.mp4

		// holo
		heartbeatUrl, hearbeatSize, storeInAmazonError := wasabis3manager.WasabiS3Client.UploadObject(
			bucketName,
			clinicName,
			"tempFiles/"+clinicName+"/"+clientIDString+"/"+"heartbeat/"+cleanFilename+".mp3",
			strconv.FormatInt(int64(clientID), 10),
			fileType)

		if storeInAmazonError != nil {
			fmt.Println("storeInAmazonError.Error()")
			fmt.Println(storeInAmazonError.Error())
		}
		heartbeat := models.Heartbeat{Filename: cleanFilename, ClientID: clientID, Url: heartbeatUrl, Size: uint(hearbeatSize), ClinicID: clinicId}
		database.GormDB.Create(&heartbeat)

		e := os.Remove(fmt.Sprintf("./tempFiles/%s/%s/%s/%s", clinicName, clientIDString, "heartbeat", cleanFilename))
		e = os.Remove(fmt.Sprintf("./tempFiles/%s/%s/%s", clinicName, clientIDString, cleanFilename))
		e = os.Remove(fmt.Sprintf("./tempFiles/%s/%s/%s", clinicName, clientIDString, cleanFilename+".mp3"))
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
