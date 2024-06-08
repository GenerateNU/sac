package mocks

import (
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/file"
)

type AWSClient struct{}

func NewAWSClient() file.FileClientInterface {
	return &AWSClient{}
}

func (c *AWSClient) UploadFile(folder string, fileHeader *multipart.FileHeader, allowedTypes []file.FileType) (*models.FileInfo, error) {
	return nil, nil
}

func (c *AWSClient) DeleteFile(fileURL string) error {
	return nil
}

func (c *AWSClient) GetFileURL(fileURL string) *string {
	return nil
}
