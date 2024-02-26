package transactions

import (
	"time"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SavePasswordResetToken(db *gorm.DB, userID uuid.UUID, token string) *errors.Error {
	passwordReset := models.PasswordReset{
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour * 24).UTC(),
	}

	if err := db.Create(&passwordReset).Error; err != nil {
		return &errors.FailedToCreatePasswordReset
	}

	return nil
}

func DeletePasswordResetToken(db *gorm.DB, token string) *errors.Error {
	if err := db.Where("token = ?", token).Delete(&models.PasswordReset{}).Error; err != nil {
		return &errors.FailedToDeletePasswordReset
	}

	return nil
}

func GetPasswordResetToken(db *gorm.DB, token string) (*models.PasswordReset, *errors.Error) {
	passwordReset := models.PasswordReset{}
	if err := db.Where("token = ?", token).First(&passwordReset).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &errors.PasswordResetTokenNotFound
		}
		return nil, &errors.FailedToGetPasswordResetToken
	}

	return &passwordReset, nil
}

func GetActivePasswordResetTokenByUserID(db *gorm.DB, userID uuid.UUID) (*models.PasswordReset, *errors.Error) {
	passwordReset := models.PasswordReset{}
	if err := db.Where("user_id = ? AND expires_at > ?", userID, time.Now().UTC()).First(&passwordReset).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &errors.PasswordResetTokenNotFound
		}
		return nil, &errors.FailedToGetPasswordResetToken
	}

	return &passwordReset, nil
}
