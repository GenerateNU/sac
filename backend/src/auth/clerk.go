package auth

import (
	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/clerkinc/clerk-sdk-go/clerk"
)

type ClerkServiceInterface interface {
	// Register(email, password string) *errors.Error
	// Login(email, password string) (string, *errors.Error)
	GetAllUsers() (string, error)
}

type ClerkService struct {
	Settings config.ClerkSettings
	client   clerk.Client
}

func NewClerkService(settings config.ClerkSettings) *ClerkService {
	client, err := clerk.NewClient(settings.APIKey.Expose())
	if err != nil {
		return nil
	}

	return &ClerkService{
		Settings: settings,
		client:   client,
	}
}

// func (c *ClerkService) Register(email, password string) *errors.Error {
// 	_, err := clerk.(c.context, email, password)
// 	if err != nil {
// 		return errors.NewError(err)
// 	}

// 	return nil
// }

func (c *ClerkService) GetAllUsers(limit int, offset int) ([]clerk.User, error) {
	users, err := c.client.Users().ListAll(clerk.ListAllUsersParams{
		Limit:  &limit,
		Offset: &offset,
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}