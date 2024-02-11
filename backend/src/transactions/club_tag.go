package transactions

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

// Create tags for a club
func CreateClubTags(db *gorm.DB, id uuid.UUID, tags []models.Tag) ([]models.Tag, *errors.Error) {
	user, err := GetClub(db, id)
	if err != nil {
		return nil, &errors.UserNotFound
	}

	if err := db.Model(&user).Association("Tag").Replace(tags); err != nil {
		return nil, &errors.FailedToUpdateUser
	}

	return tags, nil
}

// Get tags for a club
func GetClubTags(db *gorm.DB, id uuid.UUID) ([]models.Tag, *errors.Error) {
	var tags []models.Tag

	club, err := GetClub(db, id)
	if err != nil {
		return nil, &errors.ClubNotFound
	}

	if err := db.Model(&club).Association("Tag").Find(&tags); err != nil {
		return nil, &errors.FailedToGetTag
	}
	return tags, nil
}

// Delete tag for a club
func DeleteClubTag(db *gorm.DB, id uuid.UUID, tagId uuid.UUID) *errors.Error {
	club, err := GetClub(db, id)
	if err != nil {
		return &errors.ClubNotFound
	}

	tag, err := GetTag(db, tagId)
	if err != nil {
		return &errors.TagNotFound
	}

	if err := db.Model(&club).Association("Tag").Delete(&tag); err != nil {
		return &errors.FailedToUpdateClub
	}
	return nil
}
