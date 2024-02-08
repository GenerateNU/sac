package database

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
)

var SuperUserUUID uuid.UUID

func SuperUser(superUserSettings config.SuperUserSettings) (*models.User, *errors.Error) {
	passwordHash, err := auth.ComputePasswordHash(superUserSettings.Password.Expose())
	if err != nil {
		return nil, &errors.FailedToComputePasswordHash
	}

	return &models.User{
		Role:         models.Super,
		NUID:         "000000000",
		Email:        "generatesac@gmail.com",
		PasswordHash: *passwordHash,
		FirstName:    "SAC",
		LastName:     "Super",
		College:      models.KCCS,
		Year:         models.First,
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
