package helpers

import (
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/email"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/file"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/search"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

// AWSMockClient
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

// PineconeMockClient
type PineconeMockClient struct{}

// Connects to an existing Pinecone index, using the host and keys provided in settings.
func NewPineconeMockClient() search.SearchClientInterface {
	return &PineconeMockClient{}
}

func (c *PineconeMockClient) Seed(db *gorm.DB) *errors.Error {
	return nil
}

func (c *PineconeMockClient) Upsert(items []search.Searchable) *errors.Error {
	return nil
}

func (c *PineconeMockClient) Delete(items []search.Searchable) *errors.Error {
	return nil
}

func (c *PineconeMockClient) Search(item search.Searchable, topK int) ([]string, *errors.Error) {
	return []string{}, nil
}

// OpenAIMockClient
type OpenAIMockClient struct{}

func NewOpenAIMockClient() search.AIClientInterface {
	return &OpenAIMockClient{}
}

func (c *OpenAIMockClient) CreateEmbedding(items []search.Searchable) ([]search.Embedding, *errors.Error) {
	return []search.Embedding{}, nil
}

// ResendMockClient
type ResendMockClient struct{}

func NewResendMockClient() email.EmailClientInterface {
	return &ResendMockClient{}
}

func (c *ResendMockClient) SendPasswordResetEmail(name, email, token string) *errors.Error {
	return nil
}

func (c *ResendMockClient) SendEmailVerification(email, code string) *errors.Error {
	return nil
}

func (c *ResendMockClient) SendWelcomeEmail(name, email string) *errors.Error {
	return nil
}

func (c *ResendMockClient) SendPasswordChangedEmail(name, email string) *errors.Error {
	return nil
}

// JWTMockClient
type JWTMockClient struct{}

func NewJWTMockClient() auth.JWTClientInterface {
	return &JWTMockClient{}
}

func (c *JWTMockClient) GenerateTokenPair(accessClaims, refreshClaims auth.Claims) (*auth.Token, *errors.Error) {
	return &auth.Token{}, nil
}

func (c *JWTMockClient) GenerateToken(claims auth.Claims, tokenType auth.JWTType) ([]byte, *errors.Error) {
	return []byte{}, nil
}

func (c *JWTMockClient) RefreshToken(token, refreshToken string, tokenType auth.JWTType, newClaims jwt.MapClaims) ([]byte, *errors.Error) {
	return []byte{}, nil
}

func (c *JWTMockClient) ExtractClaims(tokenString string, tokenType auth.JWTType) (jwt.MapClaims, *errors.Error) {
	return jwt.MapClaims{}, nil
}

func (c *JWTMockClient) ParseToken(tokenString string, tokenType auth.JWTType) (*jwt.Token, *errors.Error) {
	return &jwt.Token{}, nil
}

func (c *JWTMockClient) IsTokenValid(tokenString string, tokenType auth.JWTType) (bool, *errors.Error) {
	return true, nil
}
