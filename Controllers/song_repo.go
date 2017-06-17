package api

import (
	"fmt"
	"net/url"
	"os"
	"time"

	common "../../Music/Common"
)

//RetrieveSong retrieve song with given id
func RetrieveSong(id string) *url.URL {

	client := common.GetSharedClient()
	bucketName := common.GetBucketName()

	found, err := client.BucketExists(bucketName)

	if err != nil {
		return nil
	}

	if found {

		presignedURL, err := client.PresignedGetObject(bucketName, id, 1000*time.Second, nil)

		if err != nil {

		}
		return presignedURL
	}

	return nil
}

func UploadSong(file *os.File) bool {
	client := common.GetSharedClient()
	bucketName := common.GetBucketName()

	found, err := client.BucketExists(bucketName)

	if err != nil {
		return false
	}

	if found {

		_, err := client.PutObject(bucketName, "object_name", file, "application/octet-stream")
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	return true
}
