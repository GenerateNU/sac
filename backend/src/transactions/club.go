package transactions

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAdminIDs(db *gorm.DB, clubID int) ([]int, *errors.Error) {
	var adminIDs []models.Membership

	if err := db.Where("club_id = ? AND membership_type = ?", clubID, models.MembershipTypeAdmin).Find(&adminIDs).Error; err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: "failed to get admin for club"}
	}

	var adminIDsInt []int
	for _, admin := range adminIDs {
		adminIDsInt = append(adminIDsInt, int(admin.UserID))
	}

	return adminIDsInt, nil
}

func GetAllUsers(db *gorm.DB) ([]models.User, *errors.Error) {
	var users []models.User

	if err := db.Unscoped().Omit("password_hash").Find(&users).Error; err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: "failed to get all users"}
	}

	return users, nil
}