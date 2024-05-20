package base

import (
	"context"
	"errors"
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(ctx context.Context, db *gorm.DB, user *models.User) (*models.User, error) {
	if err := db.WithContext(ctx).Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, utilities.ErrDuplicate
		}
		return nil, err
	}

	return user, nil
}

func GetToken(ctx context.Context, db *gorm.DB, token string, tokenType models.VerificationType) (*models.Verification, error) {
	var tokenModel models.Verification
	if err := db.WithContext(ctx).Where("token = ? AND type = ?", token, tokenType).First(&tokenModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &tokenModel, nil
}

func GetActiveTokenByUserID(ctx context.Context, db *gorm.DB, userID uuid.UUID, tokenType models.VerificationType) (*models.Verification, error) {
	var token models.Verification
	if err := db.WithContext(ctx).Where("user_id = ? AND expires_at > ? AND type = ?", userID, time.Now().UTC(), tokenType).First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &token, nil
}

func DeleteToken(ctx context.Context, db *gorm.DB, token string, tokenType models.VerificationType) error {
	if err := db.WithContext(ctx).Where("token = ? AND type = ?", token, tokenType).Delete(&models.Verification{}).Error; err != nil {
		return err
	}

	return nil
}

func SaveToken(ctx context.Context, db *gorm.DB, userID uuid.UUID, token string, tokenType models.VerificationType, expiry time.Time) error {
	tokenModel := models.Verification{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiry,
		Type:      tokenType,
	}

	if err := db.WithContext(ctx).Create(&tokenModel).Error; err != nil {
		return err
	}

	return nil
}

func VerifyEmail(ctx context.Context, db *gorm.DB, id uuid.UUID, token string) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := users.UpdateEmailVerification(ctx, tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := DeleteToken(ctx, tx, token, models.EmailVerificationType); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func PasswordReset(ctx context.Context, db *gorm.DB, id uuid.UUID, hash string, token string) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := users.UpdatePassword(tx, id, hash); err != nil {
		tx.Rollback()
		return err
	}

	if err := DeleteToken(ctx, tx, token, models.PasswordResetType); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
