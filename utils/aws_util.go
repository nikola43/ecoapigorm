package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DeleteObject(client *s3.S3, bucket string, key *string) (deleteError error) {
	request := &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    key,
	}

	res, err := client.DeleteObject(request)
	fmt.Println(err)
	fmt.Println(res)
	return
}
