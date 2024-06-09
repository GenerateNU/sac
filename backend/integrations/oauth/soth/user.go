package soth

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/markbates/goth"
)

func Into(u goth.User) *models.User {
	return &models.User{
		Role:      models.Student,
		Name:      u.Name,
		Email:     u.Email,
		AvatarURL: u.AvatarURL,
	}
}
