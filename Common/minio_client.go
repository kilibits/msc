package common

import (
	"fmt"

	"github.com/minio/minio-go"
)

var minioClient *minio.Client
var bucketName string

//Init initialize bucket details
func Init() {

	bucketName = "myBucket"

	enableSSL := true
	var err error

	if minioClient == nil {

		minioClient, err = minio.New("endpoint", "access-token", "access-secret", enableSSL)

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	//minioClient.ListBuckets()
}

//GetSharedClient get client to access buckets
func GetSharedClient() *minio.Client {
	Init()
	return minioClient
}

//GetBucketName get bucket name
func GetBucketName() string {
	Init()
	return bucketName
}
