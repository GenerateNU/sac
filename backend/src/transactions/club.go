package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

func GetClubs(db *gorm.DB, limit int, offset int) ([]models.Club, *errors.Error) {
	var clubs []models.Club
	result := db.Limit(limit).Offset(offset).Find(&clubs)
	if result.Error != nil {
		return nil, &errors.FailedToGetClubs
	}

	return clubs, nil
}

func CreateClub(db *gorm.DB, userId uuid.UUID, club models.Club) (*models.Club, *errors.Error) {
	user, err := GetUser(db, userId)
	if err != nil {
		return nil, &errors.UserNotFound
	}

	tx := db.Begin()

	if err := tx.Create(&club).Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateClub
	}

	if err := tx.Model(&club).Association("Admin").Append(user); err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateClub
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateClub
	}

	return &club, nil
}

func GetClub(db *gorm.DB, id uuid.UUID) (*models.Club, *errors.Error) {
	var club models.Club
	if err := db.First(&club, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.ClubNotFound
		} else {
			return nil, &errors.FailedToGetClub
		}
	}

	return &club, nil
}

func UpdateClub(db *gorm.DB, id uuid.UUID, club models.Club) (*models.Club, *errors.Error) {
	result := db.Model(&models.User{}).Where("id = ?", id).Updates(club)
	if result.Error != nil {
		if stdliberrors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &errors.UserNotFound
		} else {
			return nil, &errors.FailedToUpdateClub
		}
	}
	var existingClub models.Club

	err := db.First(&existingClub, id).Error
	if err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.ClubNotFound
		} else {
			return nil, &errors.FailedToCreateClub
		}
	}

	if err := db.Model(&existingClub).Updates(&club).Error; err != nil {
		return nil, &errors.FailedToUpdateUser
	}

	return &existingClub, nil
}

func DeleteClub(db *gorm.DB, id uuid.UUID) *errors.Error {
	if result := db.Delete(&models.Club{}, id); result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.ClubNotFound
		} else {
			return &errors.FailedToDeleteClub
		}
	}

	return nil
}


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

	if err := db.Model(&club).Association("Tag").Find(&tags) ; err != nil {
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