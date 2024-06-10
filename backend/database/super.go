package database

import (
	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/google/uuid"
)

var SuperUserUUID uuid.UUID

func SuperUser(superUserSettings config.SuperUserSettings) (*models.User, error) {
	return &models.User{
		Role:            models.Super,
		Email:           "generatesac@gmail.com",
		Name:            "SAC Super",
		Major0:          models.ComputerScience,
		College:         models.KCCS,
		GraduationCycle: models.May,
		GraduationYear:  2025,
	}, nil
}

func SuperClub() models.Club {
	return models.Club{
		Name:        "SAC",
		Preview:     "SAC",
		Description: "SAC",
		NumMembers:  0,
		Logo:        "https://aws.amazon.com/s3",
	}
}
