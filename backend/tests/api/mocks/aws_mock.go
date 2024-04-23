package mocks

import (
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/file"
)

type AWSMockClient struct{}

func NewAWSMockClient() file.FileClientInterface {
	return &AWSMockClient{}
}

func (c *AWSMockClient) UploadFile(folder string, fileHeader *multipart.FileHeader, allowedTypes []file.FileType) (*models.FileInfo, *errors.Error) {
	return nil, nil
}

func (c *AWSMockClient) DeleteFile(fileURL string) *errors.Error {
	return nil
}

func (c *AWSMockClient) GetFileURL(fileURL string) *string {
	return nil
}
