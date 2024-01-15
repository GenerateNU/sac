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
