package utils

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	maxPartSize = int64(1024 * 1024 * 1024 * 3) // 3GB MAX FILE
	maxRetries  = 3
)

func UploadObject(filepath string, clientID uint, tipo string) (string, int64, error) {
	awsBucketName := GetEnvVariable("AWS_BUCKET_NAME")
	s3Client := InitializeAWSConnection(
		GetEnvVariable("AWS_ACCESS_KEY"),
		GetEnvVariable("AWS_SECRET_KEY"),
		GetEnvVariable("AWS_ENDPOINT"),
		GetEnvVariable("AWS_BUCKET_REGION"))

	file, err := os.Open("./" + filepath)
	if err != nil {
		fmt.Printf("err opening file: %s", err)
		return "", 0, err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	fileType := http.DetectContentType(buffer)
	file.Read(buffer)

	fmt.Println("file.Name()")
	fmt.Println(file.Name())

	path := "/" + strconv.FormatInt(int64(clientID), 10) + "/" + tipo + "/" + strings.Split(file.Name(), "/")[2]
	input := &s3.CreateMultipartUploadInput{
		Bucket:      aws.String(awsBucketName),
		Key:         aws.String(path),
		ACL:         aws.String("public-read"),
		ContentType: aws.String(fileType),
	}
	resp, err := s3Client.CreateMultipartUpload(input)
	if err != nil {
		fmt.Println(err.Error())
		return "No se pudo conectar", 0, err
	}

	var curr, partLength int64
	var remaining = size
	var completedParts []*s3.CompletedPart
	partNumber := 1
	for curr = 0; remaining != 0; curr += partLength {
		if remaining < maxPartSize {
			partLength = remaining
		} else {
			partLength = maxPartSize
		}
		completedPart, err := UploadPart(s3Client, resp, buffer[curr:curr+partLength], partNumber)
		if err != nil {
			fmt.Println(err.Error())
			err := AbortMultipartUpload(s3Client, resp)
			if err != nil {
				fmt.Println(err.Error())
			}
			return "No se pudo subir", 0, err
		}
		remaining -= partLength
		partNumber++
		completedParts = append(completedParts, completedPart)
	}
	completeResponse, err := CompleteMultipartUpload(s3Client, resp, completedParts)
	if err != nil {
		fmt.Println(err.Error() + "abc")
		return "No se pudo subir", 0, err
	}
	var url *string

	url = completeResponse.Location

	return *url, size, nil
}

func DeleteObject(client *s3.S3, bucket string, key *string) error {
	request := &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    key,
	}

	res, err := client.DeleteObject(request)
	if err != nil {
		return err
	}
	fmt.Println(err)
	fmt.Println(res)
	return nil
}

func UploadPart(svc *s3.S3, resp *s3.CreateMultipartUploadOutput, fileBytes []byte, partNumber int) (*s3.CompletedPart, error) {
	tryNum := 1
	partInput := &s3.UploadPartInput{
		Body:          bytes.NewReader(fileBytes),
		Bucket:        resp.Bucket,
		Key:           resp.Key,
		PartNumber:    aws.Int64(int64(partNumber)),
		UploadId:      resp.UploadId,
		ContentLength: aws.Int64(int64(len(fileBytes))),
	}

	for tryNum <= maxRetries {
		uploadResult, err := svc.UploadPart(partInput)
		if err != nil {
			if tryNum == maxRetries {
				if aerr, ok := err.(awserr.Error); ok {
					return nil, aerr
				}
				return nil, err
			}
			tryNum++
		} else {
			return &s3.CompletedPart{
				ETag:       uploadResult.ETag,
				PartNumber: aws.Int64(int64(partNumber)),
			}, nil
		}
	}
	return nil, nil
}

func AbortMultipartUpload(svc *s3.S3, resp *s3.CreateMultipartUploadOutput) error {
	abortInput := &s3.AbortMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
	}
	_, err := svc.AbortMultipartUpload(abortInput)
	return err
}

func CompleteMultipartUpload(svc *s3.S3, resp *s3.CreateMultipartUploadOutput, completedParts []*s3.CompletedPart) (*s3.CompleteMultipartUploadOutput, error) {
	completeInput := &s3.CompleteMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: completedParts,
		},
	}
	return svc.CompleteMultipartUpload(completeInput)
}

func InitializeAWSConnection(accessKey, secretKey, endpoint, bucketRegion string) *s3.S3 {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(bucketRegion),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession := session.New(s3Config)
	S3Session := s3.New(newSession)
	return S3Session
}
