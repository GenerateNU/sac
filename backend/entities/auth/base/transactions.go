package base

import (
	"context"
	"errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth"
	"gorm.io/gorm"
)

func FindOrCreateUser(ctx context.Context, db *gorm.DB, user soth.User) (*models.User, error) {
	var sacUser models.User
	if err := db.WithContext(ctx).Where("email = ?", user.Email).First(&sacUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user, err := createUser(ctx, db, *user.Into())
			if err != nil {
				return nil, err
			}
			return user, nil
		} else {
			return nil, err
		}
	}

	return &sacUser, nil
}

func createUser(ctx context.Context, db *gorm.DB, user models.User) (*models.User, error) {
	tx := db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	welcomeTask := models.WelcomeTask{Name: user.Name, Email: user.Email}
	if err := tx.Create(&welcomeTask).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, tx.Commit().Error
}
