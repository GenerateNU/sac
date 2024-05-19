package mocks

import (
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/file"
)

type AWSMockClient struct{}

func NewAWSMockClient() file.FileClientInterface {
	return &AWSMockClient{}
}

func (c *AWSMockClient) UploadFile(folder string, fileHeader *multipart.FileHeader, allowedTypes []file.FileType) (*models.FileInfo, error) {
	return nil, nil
}

func (c *AWSMockClient) DeleteFile(fileURL string) error {
	return nil
}

func (c *AWSMockClient) GetFileURL(fileURL string) *string {
	return nil
}
