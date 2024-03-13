package aws

import (
	"bytes"
	"fmt"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AWSClientInterface interface {
	GetFileURL(fileURL string) *string
	UploadFile(file []byte, fileName string, folder string) (*string, *errors.Error)
	DeleteFile(fileURL string) *errors.Error
}

type AWSClient struct {
	Settings config.AWSSettings
	session  *session.Session
}

func NewAWSClient(settings config.AWSSettings) *AWSClient {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(settings.REGION.Expose()),
		Credentials: credentials.NewStaticCredentials(settings.ID.Expose(), settings.SECRET.Expose(), ""),

	})
	if err != nil {
		return nil
	}

	return &AWSClient{Settings: settings, session: sess}
}

func (c *AWSClient) GetFileURL(fileURL string) *string {
	bucket := c.Settings.BUCKET_NAME.Expose()
	fileURL = "https://" + bucket + ".s3.amazonaws.com/" + fileURL
	return &fileURL
}

func (c *AWSClient) UploadFile(file []byte, fileName string, folder string) (*string, *errors.Error) {
	svc := s3.New(c.session)

	bucket := c.Settings.BUCKET_NAME.Expose()
	key := folder + "/" + fileName

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(file),
	})
	if err != nil {
		fmt.Printf("Failed to upload data to %s/%s, %v\n", bucket, key, err)
		return nil, &errors.FailedToUploadFile
	}

	fileURL := "https://" + bucket + ".s3.amazonaws.com/" + key
	return &fileURL, nil
}

func (c *AWSClient) DeleteFile(fileURL string) *errors.Error {
	svc := s3.New(c.session)

	bucket := c.Settings.BUCKET_NAME.Expose()
	key := fileURL[len("https://"+bucket+".s3.amazonaws.com/"):]

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return &errors.FailedToDeleteFile
	}

	return nil
}