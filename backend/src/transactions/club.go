package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAdminIDs(db *gorm.DB, clubID uuid.UUID) ([]uuid.UUID, *errors.Error) {
	var adminIDs []models.Membership

	if err := db.Where("club_id = ? AND membership_type = ?", clubID, models.MembershipTypeAdmin).Find(&adminIDs).Error; err != nil {
		return nil, &errors.FailedtoGetAdminIDs
	}

	adminUUIDs := make([]uuid.UUID, 0)
	for _, adminID := range adminIDs {
		adminUUIDs = append(adminUUIDs, adminID.ClubID)
	}

	return adminUUIDs, nil
}

func GetClubs(db *gorm.DB, limit int, offset int) ([]models.Club, *errors.Error) {
	var clubs []models.Club
	result := db.Limit(limit).Offset(offset).Find(&clubs)
	if result.Error != nil {
		return nil, &errors.FailedToGetClubs
	}

	return clubs, nil
}

func CreateClub(db *gorm.DB, userId uuid.UUID, club models.Club) (*models.Club, *errors.Error) {
	user, err := GetUser(db, userId)
	if err != nil {
		return nil, &errors.UserNotFound
	}

	tx := db.Begin()

	if err := tx.Create(&club).Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateClub
	}

	if err := tx.Model(&club).Association("Admin").Append(user); err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateClub
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateClub
	}

	return &club, nil
}

func GetClub(db *gorm.DB, id uuid.UUID) (*models.Club, *errors.Error) {
	var club models.Club
	if err := db.First(&club, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.ClubNotFound
		} else {
			return nil, &errors.FailedToGetClub
		}
	}

	return &club, nil
}

func UpdateClub(db *gorm.DB, id uuid.UUID, club models.Club) (*models.Club, *errors.Error) {
	result := db.Model(&models.User{}).Where("id = ?", id).Updates(club)
	if result.Error != nil {
		if stdliberrors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &errors.UserNotFound
		} else {
			return nil, &errors.FailedToUpdateClub
		}
	}
	var existingClub models.Club

	err := db.First(&existingClub, id).Error
	if err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.ClubNotFound
		} else {
			return nil, &errors.FailedToCreateClub
		}
	}

	if err := db.Model(&existingClub).Updates(&club).Error; err != nil {
		return nil, &errors.FailedToUpdateUser
	}

	return &existingClub, nil
}

func DeleteClub(db *gorm.DB, id uuid.UUID) *errors.Error {
	if result := db.Delete(&models.Club{}, id); result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.ClubNotFound
		} else {
			return &errors.FailedToDeleteClub
		}
	}
	return nil
}

func GetClubMembers(db *gorm.DB, clubID uuid.UUID) ([]models.User, *errors.Error) {
	var users []models.User

	club, err := GetClub(db, clubID)
	if err != nil {
		return nil, &errors.UserNotFound
	}

	if err := db.Model(&club).Association("Member").Find(&users); err != nil {
		return nil, &errors.FailedToGetMembers
	}
	return users, nil
}

func CreateMembership(db *gorm.DB, clubID uuid.UUID, userID uuid.UUID) *errors.Error {
	club, err := GetClub(db, clubID)
	if err != nil {
		return &errors.ClubNotFound
	}

	user, err := GetUserWithMemberships(db, userID)
	if err != nil {
		return &errors.UserNotFound
	}

	if err := db.Model(&club).Association("Member").Replace(append(club.Member, *user)); err != nil {
		return &errors.FailedToUpdateUser
	}

	return nil
}

func CreateMembershipsByEmail(db *gorm.DB, clubID uuid.UUID, emails []string) *errors.Error {
	club, err := GetClub(db, clubID)
	if err != nil {
		return &errors.ClubNotFound
	}

	var users []models.User
	result := db.Where("email IN ?", emails).Find(&users).Error
	if result != nil {
		return &errors.UserNotFound
	}

	// append all found users to the club's member list
	currentMembers := club.Member
	for _, user := range users {
		currentMembers = append(currentMembers, user)
	}

	// update the association to use the newly calculated member list
	if err := db.Model(&club).Association("Member").Replace(currentMembers); err != nil {
		return &errors.FailedToUpdateClub
	}

	return nil
}

func DeleteMembership(db *gorm.DB, clubID uuid.UUID, userID uuid.UUID) *errors.Error {
	user, err := GetUser(db, userID)
	if err != nil {
		return &errors.UserNotFound
	}

	club, err := GetClub(db, clubID)
	if err != nil {
		return &errors.ClubNotFound
	}

	if err := db.Model(&club).Association("Member").Delete(user); err != nil {
		return &errors.FailedToUpdateClub
	}
	return nil
}

func DeleteMemberships(db *gorm.DB, clubID uuid.UUID, userIDs []uuid.UUID) *errors.Error {
	club, err := GetClub(db, clubID)
	if err != nil {
		return &errors.ClubNotFound
	}

	var users []models.User
	result := db.Where("ID IN ?", userIDs).Find(&users)

	if result != nil {
		return &errors.UserNotFound
	}

	for _, user := range users {
		if err := db.Model(&club).Association("Member").Delete(user); err != nil {
			return &errors.FailedToUpdateClub
		}
	}
	return nil
}
func GetUserFollowersForClub(db *gorm.DB, club_id uuid.UUID) ([]models.User, *errors.Error) {
	var users []models.User
	club, err := GetClub(db, club_id)
	if err != nil {
		return nil, &errors.ClubNotFound
	}

	if err := db.Model(&club).Association("Follower").Find(&users); err != nil {
		return nil, &errors.FailedToGetClubFollowers
	}
	return users, nil
}
