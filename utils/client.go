package utils

import (
	"fmt"

	"github.com/minio/minio-go"
)

//Client minio client and bucket structure
type Client struct {
	MinioClient *minio.Client
	Bucket      string
}

//MscClient minio client
var MscClient *Client

func init() {

	bucketName := "myBucket"
	enableSSL := true
	var err error
	var minioClient *minio.Client

	if minioClient == nil {
		minioClient, err = minio.New("endpoint", "access-token", "access-secret", enableSSL)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	MscClient = &Client{
		Bucket:      bucketName,
		MinioClient: minioClient,
	}
}
