package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/awsmanager"
	"github.com/nikola43/ecoapigorm/services"
	"github.com/nikola43/ecoapigorm/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

func UploadMultimedia(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)
	uploadMode, _ := strconv.ParseUint(context.Params("upload_mode"), 10, 64)
	uploadedFile, err := context.FormFile("file")
	employeeTokenClaims, err := utils.GetEmployeeTokenClaims(context)
	if err != nil {
		return context.SendStatus(fiber.StatusUnauthorized)
	}

	bucketName := strings.ToLower(strings.ReplaceAll(employeeTokenClaims.CompanyName, " ", ""))
	if err != nil {
		return err
	}

	err = services.UploadMultimedia(context, bucketName, uint(clientID), uploadedFile, uint(uploadMode))
	if err != nil {
		return err
	}

	e := os.Remove("./tempFiles/" + uploadedFile.Filename)
	if e != nil {
		log.Fatal(e)
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"file": uploadedFile.Filename,
	})
}

func DownloadAllMultimediaContentByClientID(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)

	// download images
	images, err := services.GetAllImagesByClientID(uint(clientID))
	os.Mkdir("tempFiles/"+context.Params("client_id")+"/images", os.ModePerm)
	for _, image := range images {

		fmt.Println(image.Url)
		filename := strings.Split(image.Url, "/")[len(strings.Split(image.Url, "/"))-1]
		fmt.Println(filename)
		err = awsmanager.AwsManager.DownloadObject(filename, "tempFiles/"+context.Params("client_id")+"/images/"+filename)
		if err != nil {
			fmt.Println(err)
		}
	}
	// download videos
	/*
		videos, err := services.GetAllVideosByClientID(uint(clientID))
		os.Mkdir("tempFiles/"+context.Params("client_id")+"/images", os.ModePerm)
		for _, video := range videos {
			filename := strings.Split(video.Url, "/")[len(strings.Split(video.Url, "/"))-1]
			err = DownloadFile("tempFiles/"+context.Params("client_id")+"/"+filename, video.Url)
			if err != nil {
				fmt.Println(err)
			}
		}
	*/

	return context.SendStatus(fiber.StatusOK)
}

func DeleteImage(context *fiber.Ctx) error {
	id, _ := strconv.ParseUint(context.Params("id"), 10, 64)

	err := services.DeleteImage(uint(id))
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println(id)

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})

}

func DeleteVideo(context *fiber.Ctx) error {
	id, _ := strconv.ParseUint(context.Params("video_id"), 10, 64)

	fmt.Println(id)

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})

}
func DeleteHolographic(context *fiber.Ctx) error {
	id, _ := strconv.ParseUint(context.Params("holopraphic_id"), 10, 64)

	fmt.Println(id)

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})

}

func DeleteHeartbeat(context *fiber.Ctx) error {
	id, _ := strconv.ParseUint(context.Params("heartbeat_id"), 10, 64)

	fmt.Println(id)

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})

}
