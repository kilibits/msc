package api

import (
	"time"

	common "../../Music/Common"
)

//RetrieveSong retrieve song with given id
func RetrieveSong(id string) string {

	client := common.GetSharedClient()
	bucketName := common.GetBucketName()

	found, err := client.BucketExists(bucketName)

	if err != nil {
		return
	}

	if found {



	presignedURL, err := client.PresignedGetObject(bucketName, id, 1000*time.Second, nil)

	if err != nil {

	}
	return presignedURL.String()
	}
}
