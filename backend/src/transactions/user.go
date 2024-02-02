package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) (*models.User, *errors.Error) {
	if err := db.Create(user).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, &errors.UserAlreadyExists
		} else {
			return nil, &errors.FailedToCreateUser
		}
	}

	return user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*models.User, *errors.Error) {
	var user models.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, &errors.UserNotFound
	}

	return &user, nil
}

func GetUsers(db *gorm.DB, limit int, offset int) ([]models.User, *errors.Error) {
	var users []models.User

	if err := db.Omit("password_hash").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, &errors.FailedToGetUsers
	}

	return users, nil
}

func GetUser(db *gorm.DB, id uuid.UUID) (*models.User, *errors.Error) {
	var user models.User
	if err := db.Omit("password_hash").First(&user, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.UserNotFound
		} else {
			return nil, &errors.FailedToGetUser
		}
	}

	return &user, nil
}

func GetUserWithFollowers(db *gorm.DB, id uuid.UUID) (*models.User, *errors.Error) {
	var user models.User
	if err := db.Preload("Follower").Omit("password_hash").First(&user, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.UserNotFound
		} else {
			return nil, &errors.FailedToGetUser
		}
	}

	return &user, nil
}

func UpdateUser(db *gorm.DB, id uuid.UUID, user models.User) (*models.User, *errors.Error) {
	var existingUser models.User

	err := db.First(&existingUser, id).Error
	if err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.UserNotFound
		} else {
			return nil, &errors.FailedToUpdateTag
		}
	}

	if err := db.Model(&existingUser).Updates(&user).Error; err != nil {
		return nil, &errors.FailedToUpdateUser
	}

	return &existingUser, nil
}

func DeleteUser(db *gorm.DB, id uuid.UUID) *errors.Error {
	result := db.Delete(&models.User{}, id)
	if result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.UserNotFound
		} else {
			return &errors.FailedToDeleteUser
		}
	}
	return nil
}

// Create following for a user
func CreateFollowing(db *gorm.DB, userId uuid.UUID, clubId uuid.UUID) *errors.Error {

	user, err := GetUserWithFollowers(db, userId)
	if err != nil {
		return &errors.UserNotFound
	}
	club, err := GetClub(db, clubId)
	if err != nil {
		return &errors.ClubNotFound
	}

	if err := db.Model(&user).Association("Follower").Replace(append(user.Follower, *club)); err != nil {
		return &errors.FailedToUpdateUser
	}
	return nil
}

// Delete following for a user
func DeleteFollowing(db *gorm.DB, userId uuid.UUID, clubId uuid.UUID) *errors.Error {
	user, err := GetUser(db, userId)
	if err != nil {
		return &errors.UserNotFound
	}
	club, err := GetClub(db, clubId)
	if err != nil {
		return &errors.ClubNotFound
	}
	//What to return here?
	//Should we return User or Success message?
	if err := db.Model(&user).Association("Follower").Delete(club); err != nil {
		return &errors.FailedToUpdateUser
	}
	return nil
}

// Get all following for a user

func GetClubFollowing(db *gorm.DB, userId uuid.UUID) ([]models.Club, *errors.Error) {
	var clubs []models.Club

	user, err := GetUser(db, userId)
	if err != nil {
		return nil, &errors.UserNotFound
	}

	if err := db.Model(&user).Association("Follower").Find(&clubs); err != nil {
		return nil, &errors.FailedToGetUserFollowing
	}
	return clubs, nil
}
