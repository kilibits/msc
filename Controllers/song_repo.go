package api

import (
	"fmt"
	"net/url"
	"os"
	"time"

	common "../../msc/Common"
)

//RetrieveSong retrieve song with given id
func RetrieveSong(id string) *url.URL {

	client := common.GetClient()
	bucketName := client.Bucket
	minioClient := client.MinioClient

	found, err := minioClient.BucketExists(bucketName)

	if err != nil {
		return nil
	}

	if found {

		presignedURL, err := minioClient.PresignedGetObject(bucketName, id, 1000*time.Second, nil)

		if err != nil {

		}
		return presignedURL
	}

	return nil
}

func UploadSong(file *os.File) bool {
	client := common.GetClient()
	bucketName := client.Bucket
	minioClient := client.MinioClient

	found, err := minioClient.BucketExists(bucketName)

	if err != nil {
		return false
	}

	if found {

		_, err := minioClient.PutObject(bucketName, "object_name", file, "application/octet-stream")
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	return true
}
