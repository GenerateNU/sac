package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAdminIDs(db *gorm.DB, clubID uuid.UUID) ([]uuid.UUID, *errors.Error) {
	var adminIDs []models.Membership

	if err := db.Where("club_id = ? AND membership_type = ?", clubID, models.MembershipTypeAdmin).Find(&adminIDs).Error; err != nil {
		return nil, &errors.FailedtoGetAdminIDs
	}

	adminUUIDs := make([]uuid.UUID, 0)
	for _, adminID := range adminIDs {
		adminUUIDs = append(adminUUIDs, adminID.ClubID)
	}

	return adminUUIDs, nil
}

func GetClubs(db *gorm.DB, params *models.ClubQueryParams) ([]models.Club, *errors.Error) {
	offset := (*&params.Page - 1) * *&params.Limit
	var clubs []models.Club

	db_call := db
	if params.Tags != nil {
		//We only want to query tags if necessary
		//i.e. when we filter
		db_call = db.Preload("Tag")
	}
	query := db_call.Where(params.IntoWhere())
	if params.Tags != nil {
		query = query.Where("id IN (SELECT club_id FROM club_tags WHERE tag_id IN ?)", params.Tags)
	}
	result := query.Limit(params.Limit).Offset(offset).Find(&clubs)

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
