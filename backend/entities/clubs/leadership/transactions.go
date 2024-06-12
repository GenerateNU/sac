package leadership

import (
	"errors"

	"github.com/GenerateNU/sac/backend/database/cache"
	"github.com/GenerateNU/sac/backend/entities/files/base"
	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateClubLeader(db *gorm.DB, leader models.Leader, fileIntegrationUpload func() (*models.FileInfo, error), fileIntegrationDelete func(string) error) (*models.Leader, error) {
	fileInfo, err := fileIntegrationUpload()
	if err != nil {
		return nil, err
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := db.Create(&leader).Error; err != nil {
		return nil, err
	}

	file, err := base.CreateFile(tx, leader.ID, "leadership", *fileInfo)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		if err := fileIntegrationDelete(fileInfo.FileURL); err != nil {
			return nil, err
		}

		return nil, err
	}

	leader.PhotoFile = *file

	return &leader, nil
}

func GetClubLeadership(db *gorm.DB, clubID uuid.UUID) ([]models.Leader, error) {
	db = cache.SetUseCache(db, true)

	var leadership []models.Leader
	result := db.Preload("PhotoFile").Where("club_id = ?", clubID).Find(&leadership)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, result.Error
	}

	return leadership, nil
}

func GetClubLeader(db *gorm.DB, clubID uuid.UUID, leaderID uuid.UUID) (*models.Leader, error) {
	db = cache.SetUseCache(db, true)

	var leader models.Leader
	if err := db.Preload("PhotoFile").First(&leader, "id = ? AND club_id = ?", leaderID, clubID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &leader, nil
}

func GetClubLeaderByClubIDAndEmail(db *gorm.DB, clubID uuid.UUID, email string) (*models.Leader, error) {
	db = cache.SetUseCache(db, true)

	var leader models.Leader
	if err := db.First(&leader, "email = ? AND club_id = ?", email, clubID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &leader, nil
}

func UpdateClubLeader(db *gorm.DB, clubID uuid.UUID, leaderID uuid.UUID, updatedLeader models.Leader) (*models.Leader, error) {
	leader, err := GetClubLeader(db, clubID, leaderID)
	if err != nil {
		return nil, err
	}
	if err := db.Model(&leader).Clauses(clause.Returning{}).Updates(&updatedLeader).Error; err != nil {
		return nil, err
	}

	return leader, nil
}

func DeleteClubLeader(db *gorm.DB, clubID uuid.UUID, leaderID uuid.UUID, fileIntegrationDelete func(string) error) error {
	leader, err := GetClubLeader(db, clubID, leaderID)
	if err != nil {
		return err
	}

	if err := fileIntegrationDelete(leader.PhotoFile.FileURL); err != nil {
		return err
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	err = base.DeleteFile(tx, leader.PhotoFile.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if result := db.Delete(&models.Leader{}, "id = ? AND club_id = ?", leaderID, clubID); result.RowsAffected == 0 {
		tx.Rollback()
		if result.Error == nil {
			return utilities.ErrNotFound
		}
		return result.Error
	}

	return tx.Commit().Error
}
