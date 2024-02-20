package transactions

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateMember(db *gorm.DB, userId uuid.UUID, clubId uuid.UUID) *errors.Error {
	user, err := GetUser(db, userId)
	if err != nil {
		return err
	}

	club, err := GetClub(db, clubId)
	if err != nil {
		return err
	}

	// this doesnt mean the user is a member of the club, it could be empty
	if err := db.Model(&user).Association("Member").Find(&club); err == nil {
		return &errors.AlreadyMemberOfClub
	}

	if err := db.Model(&user).Association("Member").Append(&club); err != nil {
		return &errors.FailedToUpdateUser
	}

	return nil
}

func DeleteMember(db *gorm.DB, userId uuid.UUID, clubId uuid.UUID) *errors.Error {
	user, err := GetUser(db, userId, PreloadMember())
	if err != nil {
		return err
	}

	club, err := GetClub(db, clubId, PreloadMember())
	if err != nil {
		return err
	}

	if err := db.Model(&user).Association("Member").Find(&club); err != nil {
		return &errors.UserNotMemberOfClub
	}

	if err := db.Model(&user).Association("Member").Delete(&club); err != nil {
		return &errors.FailedToUpdateUser
	}

	// userMemberClubIDs := make([]uuid.UUID, len(user.Member))

	// for i, club := range user.Member {
	// 	userMemberClubIDs[i] = club.ID
	// }

	// if !slices.Contains(userMemberClubIDs, club.ID) {
	// 	return &errors.UserNotMemberOfClub
	// }

	// if err := db.Model(&user).Association("Member").Delete(club); err != nil {
	// 	return &errors.FailedToUpdateUser
	// }

	return nil
}

func GetClubMembership(db *gorm.DB, userId uuid.UUID) ([]models.Club, *errors.Error) {
	var clubs []models.Club

	user, err := GetUser(db, userId)
	if err != nil {
		return nil, err
	}

	if err := db.Model(&user).Association("Member").Find(&clubs); err != nil {
		return nil, &errors.FailedToGetUserMemberships
	}

	return clubs, nil
}
