package base

import (
	"context"
	"errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/oauth"
	"gorm.io/gorm"
)

func FindOrCreateUser(ctx context.Context, db *gorm.DB, user oauth.User) (*models.User, error) {
	var sacUser models.User
	if err := db.WithContext(ctx).Where("email = ?", user.Email).First(&sacUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := createUser(ctx, db, user.Into()); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &sacUser, nil
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
