package file

import (
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
)

type AWSMockClient struct{}

func NewAWSMockClient() FileClientInterface {
	return &AWSMockClient{}
}

func (c *AWSMockClient) UploadFile(folder string, fileHeader *multipart.FileHeader, allowedTypes []FileType) (*models.FileInfo, *errors.Error) {
	return nil, nil
}

func (c *AWSMockClient) DeleteFile(fileURL string) *errors.Error {
	return nil
}

func (c *AWSMockClient) GetFileURL(fileURL string) *string {
	return nil
}
