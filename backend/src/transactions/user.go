package transactions

import (
	"backend/src/models"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User

	if err := db.Unscoped().Omit("password_hash").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(db *gorm.DB, id string) (*models.User, error) {
	var user models.User
	if err := db.Omit("role").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	fmt.Printf("%+v\n", user)

	return &user, nil
}
