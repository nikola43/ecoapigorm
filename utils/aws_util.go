package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadObject(client *s3.S3, bucket string, key *string) error {
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
