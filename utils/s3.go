package utils

import (
	"bytes"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// DownloadFromS3 download file from s3 using aws-sdk
func DownloadFromS3(bucket, item string) (data io.Reader, err error) {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	downloader := s3manager.NewDownloader(sess)

	buff := &aws.WriteAtBuffer{}

	_, err = downloader.Download(buff,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})

	data = bytes.NewReader(buff.Bytes())

	if err != nil {
		return nil, err
	}

	return data, nil
}
