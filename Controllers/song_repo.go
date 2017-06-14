package api

import (
	"time"
	"net/url"
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
