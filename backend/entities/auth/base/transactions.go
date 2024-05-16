package base

import (
	"errors"
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) (*models.User, error) {
	if err := db.Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, utilities.ErrDuplicate
		}
		return nil, err
	}

	return user, nil
}

func GetToken(db *gorm.DB, token string, tokenType models.VerificationType) (*models.Verification, error) {
	tokenModel := models.Verification{}
	if err := db.Where("token = ? AND type = ?", token, tokenType).First(&tokenModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &tokenModel, nil
}

func GetActiveTokenByUserID(db *gorm.DB, userID uuid.UUID, tokenType models.VerificationType) (*models.Verification, error) {
	token := models.Verification{}
	if err := db.Where("user_id = ? AND expires_at > ? AND type = ?", userID, time.Now().UTC(), tokenType).First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &token, nil
}

func DeleteToken(db *gorm.DB, token string, tokenType models.VerificationType) error {
	if err := db.Where("token = ? AND type = ?", token, tokenType).Delete(&models.Verification{}).Error; err != nil {
		return err
	}

	return nil
}

func SaveToken(db *gorm.DB, userID uuid.UUID, token string, tokenType models.VerificationType, expiry time.Time) error {
	tokenModel := models.Verification{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiry,
		Type:      tokenType,
	}

	if err := db.Create(&tokenModel).Error; err != nil {
		return err
	}

	return nil
}
