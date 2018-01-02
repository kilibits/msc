package utils

import (
	"fmt"
	"net/url"
	"os"
	"time"
)

//RetrieveSong retrieve song with given id
func (client *Client) RetrieveSong(id string) *url.URL {
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

//UploadSong upload song to bucket
func (client *Client) UploadSong(file *os.File) bool {
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
