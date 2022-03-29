package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"math/big"
	"strings"
)

var rpc1 *FtmBridge
var rpc2 *FtmBridge
var totalErrors int64
var block big.Int
var skipRows int64

func main() {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
		Credentials: credentials.NewStaticCredentials("id", "secret", ""),
	}))

	uploader := s3manager.NewUploader(sess)

	r := strings.NewReader("Hello, Reader 2!")

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("jkalina-bucket-1"),
		Key:    aws.String("testing-key"),
		Body:   r,
	})
	if err != nil {
		fmt.Printf("failed to upload file, %v", err)
		return
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
}
