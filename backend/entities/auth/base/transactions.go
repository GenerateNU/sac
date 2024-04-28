package base

import (
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetToken(db *gorm.DB, token string, tokenType models.VerificationType) (*models.Verification, *errors.Error) {
	tokenModel := models.Verification{}
	if err := db.Where("token = ? AND type = ?", token, tokenType).First(&tokenModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &errors.TokenNotFound
		}
		return nil, &errors.FailedToGetToken
	}

	return &tokenModel, nil
}

func GetActiveTokenByUserID(db *gorm.DB, userID uuid.UUID, tokenType models.VerificationType) (*models.Verification, *errors.Error) {
	token := models.Verification{}
	if err := db.Where("user_id = ? AND expires_at > ? AND type = ?", userID, time.Now().UTC(), tokenType).First(&token).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &errors.TokenNotFound
		}
		return nil, &errors.FailedToGetToken
	}

	return &token, nil
}

func DeleteToken(db *gorm.DB, token string, tokenType models.VerificationType) *errors.Error {
	if err := db.Where("token = ? AND type = ?", token, tokenType).Delete(&models.Verification{}).Error; err != nil {
		return &errors.FailedToDeleteToken
	}

	return nil
}

func SaveToken(db *gorm.DB, userID uuid.UUID, token string, tokenType models.VerificationType, expiry time.Time) *errors.Error {
	tokenModel := models.Verification{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiry,
		Type:      tokenType,
	}

	if err := db.Create(&tokenModel).Error; err != nil {
		return &errors.FailedToSaveToken
	}

	return nil
}
