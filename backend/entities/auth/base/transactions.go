package base

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/oauth"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FindOrCreateUser(ctx context.Context, db *gorm.DB, user oauth.User) (*models.User, error) {
	var sacUser models.User
	if err := db.WithContext(ctx).Where("email = ?", user.Email).First(&sacUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := createUser(ctx, db, &models.User{Email: user.Email, Name: user.Name, Role: models.Student}); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &sacUser, nil
}

func GetToken(db *gorm.DB, token string) (*models.Verification, error) {
	tokenModel := models.Verification{}
	if err := db.Where("token = ?", token).First(&tokenModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &tokenModel, nil
}

func GetActiveTokenByUserID(db *gorm.DB, userID uuid.UUID) (*models.Verification, error) {
	token := models.Verification{}
	if err := db.Where("user_id = ? AND expires_at > ?", userID, time.Now().UTC()).First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &token, nil
}

func DeleteToken(db *gorm.DB, token string) error {
	if err := db.Where("token = ?", token).Delete(&models.Verification{}).Error; err != nil {
		return err
	}

	return nil
}

func EnqueueVerification(ctx context.Context, db *gorm.DB, userID uuid.UUID, email string) error {
	tx := db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	otp, err := auth.GenerateOTP(constants.OTP_LENGTH)
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}

	verification := models.Verification{
		UserID:    userID,
		Email:     email,
		Token:     *otp,
		ExpiresAt: time.Now().Add(constants.OTP_EXPIRY).UTC(),
	}

	if err := tx.Create(&verification).Error; err != nil {
		return err
	}

	task := models.VerificationTask{
		UserID:         userID,
		VerificationID: verification.ID,
	}

	if err := tx.Create(&task).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

func createUser(ctx context.Context, db *gorm.DB, user *models.User) error {
	tx := db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	welcomeTask := models.WelcomeTask{Name: user.Name, Email: user.Email}
	if err := tx.Create(&welcomeTask).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
