package soth

import (
	"encoding/gob"
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
)

func init() {
	gob.Register(User{})
}

// User contains the information common amongst most OAuth and OAuth2 providers.
// All the "raw" data from the provider can be found in the `RawData` field.
type User struct {
	RawData           map[string]interface{}
	Provider          string
	Email             string
	Name              string
	FirstName         string
	LastName          string
	NickName          string
	Description       string
	UserID            string
	AvatarURL         string
	Location          string
	AccessToken       string
	AccessTokenSecret string
	RefreshToken      string
	ExpiresAt         time.Time
	IDToken           string
}

func (u *User) Into() *models.User {
	return &models.User{
		Role:      models.Student,
		Name:      u.Name,
		Email:     u.Email,
		AvatarURL: u.AvatarURL,
	}
}
