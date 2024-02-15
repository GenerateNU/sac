package transactions

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserTags(db *gorm.DB, id uuid.UUID) ([]models.Tag, *errors.Error) {
	var tags []models.Tag

	user, err := GetUser(db, id)
	if err != nil {
		return nil, err
	}

	if err := db.Model(&user).Association("Tag").Find(&tags); err != nil {
		return nil, &errors.FailedToGetTag
	}
	return tags, nil
}

func CreateUserTags(db *gorm.DB, id uuid.UUID, tags []models.Tag) ([]models.Tag, *errors.Error) {
	user, err := GetUser(db, id, PreloadTag())
	if err != nil {
		return nil, err
	}

	if err := db.Model(&user).Association("Tag").Append(tags); err != nil {
		return nil, &errors.FailedToUpdateUser
	}

	return tags, nil
}
