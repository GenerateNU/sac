package transactions

import (
	"backend/src/models"

	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User

	if err := db.Unscoped().Omit("password_hash").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUser(db *gorm.DB, id string, payload models.User) (models.User, error) {
	var existingUser models.User

	if err := db.First(&existingUser, id).Error; err != nil {
		return models.User{}, err
	}

	if err := db.Model(&existingUser).Updates(&payload).Error; err != nil {
		return models.User{}, err
	}

	return existingUser, nil
}