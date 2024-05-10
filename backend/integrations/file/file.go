package file

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

type FileType string

const (
	IMAGE       FileType = "image"
	VIDEO       FileType = "video"
	AUDIO       FileType = "audio"
	DOCUMENT    FileType = "document"
	APPLICATION FileType = "application"
	TEXT        FileType = "text"
	ALL         FileType = "all"
)

type FileClientInterface interface {
	GetFileURL(fileURL string) *string
	UploadFile(folder string, fileHeader *multipart.FileHeader, allowedTypes []FileType) (*models.FileInfo, error)
	DeleteFile(fileURL string) error
}

type AWSClient struct {
	Settings config.AWSSettings
	session  *session.Session
}

func NewAWSProvider(settings config.AWSSettings) FileClientInterface {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(settings.REGION.Expose()),
		Credentials: credentials.NewStaticCredentials(settings.ID.Expose(), settings.SECRET.Expose(), ""),
	})
	if err != nil {
		return nil
	}

	return &AWSClient{Settings: settings, session: sess}
}

func (aw *AWSClient) GetFileURL(fileURL string) *string {
	fileURL = fmt.Sprintf("https://s3.amazonaws.com/%s/%s", aw.Settings.BUCKET_NAME.Expose(), fileURL)
	return &fileURL
}

func preProcessFile(fileHeader *multipart.FileHeader) (*string, []byte, error) {
	if fileHeader.Size > constants.MAX_FILE_SIZE {
		return nil, nil, errors.New("file size is too large")
	}

	file, err := fileHeader.Open()
	defer file.Close()
	if err != nil {
		return nil, nil, errors.New("failed to open file")
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, errors.New("failed to read file")
	}

	fileName := generateUniqueFileName(fileHeader.Filename)

	return &fileName, fileBytes, nil
}

func generateUniqueFileName(fileName string) string {
	ext := fileName[strings.LastIndex(fileName, "."):]
	fileNameWithoutExt := fileName[:len(fileName)-len(ext)]
	uniqueFileName := fmt.Sprintf("%s-%s%s", fileNameWithoutExt, uuid.New().String(), ext)
	return uniqueFileName
}

func isFileTypeAllowed(fileType string, allowedTypes []FileType) bool {
	if len(allowedTypes) == 0 {
		return false
	}

	mainType := strings.Split(fileType, "/")[0]
	if len(allowedTypes) == 1 && allowedTypes[0] == ALL {
		return true
	}

	for _, allowedType := range allowedTypes {
		allowedMainType := strings.Split(string(allowedType), "/")[0]
		if allowedMainType == mainType {
			return true
		}
	}
	return false
}

func (aw *AWSClient) UploadFile(folder string, fileHeader *multipart.FileHeader, allowedTypes []FileType) (*models.FileInfo, error) {
	fileType := fileHeader.Header.Get("Content-Type")
	if !isFileTypeAllowed(fileType, allowedTypes) {
		return nil, errors.New("file type is not allowed")
	}

	fileName, file, err := preProcessFile(fileHeader)
	if err != nil {
		return nil, err
	}

	svc := s3.New(aw.session)

	bucket := aw.Settings.BUCKET_NAME.Expose()
	key := fmt.Sprintf("%s/%s", folder, *fileName)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(file),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %s", err)
	}

	return &models.FileInfo{
		FileName:  *fileName,
		FileType:  fileHeader.Header.Get("Content-Type"),
		FileSize:  int(fileHeader.Size),
		FileURL:   fmt.Sprintf("https://s3.amazonaws.com/%s/%s", bucket, key),
		ObjectKey: key,
	}, nil
}

func (aw *AWSClient) DeleteFile(fileURL string) error {
	svc := s3.New(aw.session)

	bucket := aw.Settings.BUCKET_NAME.Expose()
	key := fileURL[len(fmt.Sprintf("https://s3.amazonaws.com/%s/", bucket)):]

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file: %s", err)
	}

	return nil
}
