package tags

import (
	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/tags"
	"github.com/GenerateNU/sac/backend/transactions"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

func CreateClubTags(db *gorm.DB, id uuid.UUID, tags []models.Tag) ([]models.Tag, error) {
	club, err := clubs.GetClub(db, id, transactions.PreloadTag())
	if err != nil {
		return nil, err
	}

	if err := db.Model(&club).Association("Tag").Append(tags); err != nil {
		return nil, err
	}

	return tags, nil
}

func GetClubTags(db *gorm.DB, id uuid.UUID) ([]models.Tag, error) {
	var tags []models.Tag

	club, err := clubs.GetClub(db, id, transactions.PreloadTag())
	if err != nil {
		return nil, err
	}

	if err := db.Model(&club).Association("Tag").Find(&tags); err != nil {
		return nil, err
	}
	return tags, nil
}

func DeleteClubTag(db *gorm.DB, id uuid.UUID, tagID uuid.UUID) error {
	club, err := clubs.GetClub(db, id, transactions.PreloadTag())
	if err != nil {
		return err
	}

	tag, err := tags.GetTag(db, tagID)
	if err != nil {
		return err
	}

	if err := db.Model(&club).Association("Tag").Delete(&tag); err != nil {
		return err
	}
	return nil
}
