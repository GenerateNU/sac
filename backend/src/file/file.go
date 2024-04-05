package file

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

type FileType string

const (
	IMAGE       FileTypes = "image"
	VIDEO       FileTypes = "video"
	AUDIO       FileTypes = "audio"
	DOCUMENT    FileTypes = "document"
	APPLICATION FileTypes = "application"
	TEXT        FileTypes = "text"
	ALL         FileTypes = "all"
)

type FileClientInterface interface {
	GetFileURL(fileURL string) *string
	UploadFile(folder string, fileHeader *multipart.FileHeader, allowedTypes []FileTypes) (*models.FileInfo, *errors.Error)
	DeleteFile(fileURL string) *errors.Error
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
	fileURL = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", aw.Settings.BUCKET_NAME.Expose(), fileURL)
	return &fileURL
}

func preProcessFile(fileHeader *multipart.FileHeader) (*string, []byte, *errors.Error) {
	if fileHeader.Size > 5000000 {
		return nil, nil, &errors.InvalidFileSize
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, nil, &errors.FailedToOpenFile
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, &errors.FailedToReadFile
	}

	defer file.Close()

	fileName := generateUniqueFileName(fileHeader.Filename)

	return &fileName, fileBytes, nil
}

func generateUniqueFileName(fileName string) string {
	ext := fileName[strings.LastIndex(fileName, "."):]
	fileNameWithoutExt := fileName[:len(fileName)-len(ext)]
	uniqueFileName := fmt.Sprintf("%s-%s%s", fileNameWithoutExt, uuid.New().String(), ext)
	return uniqueFileName
}

func isFileTypeAllowed(fileType string, allowedTypes []FileTypes) bool {
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

func (aw *AWSClient) UploadFile(folder string, fileHeader *multipart.FileHeader, allowedTypes []FileTypes) (*models.FileInfo, *errors.Error) {
	fileType := fileHeader.Header.Get("Content-Type")
	if !isFileTypeAllowed(fileType, allowedTypes) {
		return nil, &errors.InvalidFileType
	}

	fileName, file, err := preProcessFile(fileHeader)
	if err != nil {
		return nil, err
	}

	svc := s3.New(aw.session)

	bucket := aw.Settings.BUCKET_NAME.Expose()
	key := fmt.Sprintf("%s/%s", folder, *fileName)

	_, s3Err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(file),
	})
	if s3Err != nil {
		fmt.Printf("Failed to upload data to %s/%s, %v\n", bucket, key, err)
		return nil, &errors.FailedToUploadFile
	}

	fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key)
	return &models.FileInfo{
		FileName:  *fileName,
		FileType:  fileHeader.Header.Get("Content-Type"),
		FileSize:  int(fileHeader.Size),
		FileURL:   fileURL,
		ObjectKey: key,
	}, nil
}

func (aw *AWSClient) DeleteFile(fileURL string) *errors.Error {
	svc := s3.New(aw.session)

	bucket := aw.Settings.BUCKET_NAME.Expose()
	key := fileURL[len(fmt.Sprintf("https://%s.s3.amazonaws.com/", bucket)):]

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return &errors.FailedToDeleteFile
	}

	return nil
}
