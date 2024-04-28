package database

import (
	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/google/uuid"
)

var SuperUserUUID uuid.UUID

func SuperUser(superUserSettings config.SuperUserSettings) (*models.User, *errors.Error) {
	passwordHash, err := auth.ComputeHash(superUserSettings.Password.Expose())
	if err != nil {
		return nil, &errors.FailedToComputePasswordHash
	}

	return &models.User{
		Role:            models.Super,
		Email:           "generatesac@gmail.com",
		PasswordHash:    *passwordHash,
		FirstName:       "SAC",
		LastName:        "Super",
		Major0:          models.ComputerScience,
		College:         models.KCCS,
		GraduationCycle: models.May,
		GraduationYear:  2025,
	}, nil
}

func SuperClub() models.Club {
	return models.Club{
		Name:             "SAC",
		Preview:          "SAC",
		Description:      "SAC",
		NumMembers:       0,
		IsRecruiting:     true,
		RecruitmentCycle: models.Always,
		RecruitmentType:  models.Application,
		ApplicationLink:  "https://generatenu.com/apply",
		Logo:             "https://aws.amazon.com/s3",
	}
}
