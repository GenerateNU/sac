package recruitment

import (
	"errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateClubRecruitment(db *gorm.DB, clubID *uuid.UUID, recruitment *models.Recruitment) (*models.Club, error) {
	return UpdateClubRecruitment(db, clubID, recruitment)
}

func UpdateClubRecruitment(db *gorm.DB, clubID *uuid.UUID, recruitment *models.Recruitment) (*models.Club, error) {
	var club models.Club
	if err := db.Model(&club).Where("id = ?", clubID).Update("recruitment", recruitment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &club, nil
}

func DeleteClubRecruitment(db *gorm.DB, clubID *uuid.UUID) error {
	_, err := UpdateClubRecruitment(db, clubID, nil)
	return err
}
