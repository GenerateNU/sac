package tag

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/tags"
	"github.com/GenerateNU/sac/backend/entities/users"
	"github.com/GenerateNU/sac/backend/transactions"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserTags(db *gorm.DB, id uuid.UUID) ([]models.Tag, error) {
	var tags []models.Tag

	user, err := users.GetUser(db, id)
	if err != nil {
		return nil, err
	}

	if err := db.Model(&user).Association("Tag").Find(&tags); err != nil {
		return nil, err
	}
	return tags, nil
}

func CreateUserTags(db *gorm.DB, id uuid.UUID, tags []models.Tag) ([]models.Tag, error) {
	user, err := users.GetUser(db, id, transactions.PreloadTag())
	if err != nil {
		return nil, err
	}

	if err := db.Model(&user).Association("Tag").Append(tags); err != nil {
		return nil, err
	}

	return tags, nil
}

func DeleteUserTag(db *gorm.DB, id uuid.UUID, tagID uuid.UUID) error {
	user, err := users.GetUser(db, id, transactions.PreloadTag())
	if err != nil {
		return err
	}

	tag, err := tags.GetTag(db, tagID)
	if err != nil {
		return err
	}

	if err := db.Model(&user).Association("Tag").Delete(&tag); err != nil {
		return err
	}

	return nil
}
