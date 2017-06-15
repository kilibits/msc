package common

import (
	"fmt"

	"github.com/minio/minio-go"
)

var minioClient *minio.Client
var bucketName string

func init() {

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
}

//GetSharedClient get client to access buckets
func GetSharedClient() *minio.Client {
	return minioClient
}

//GetBucketName get bucket name
func GetBucketName() string {
	return bucketName
}
