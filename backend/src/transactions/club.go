package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/gofiber/fiber/v2"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func GetContacts(db *gorm.DB, limit int, offset int) ([]models.Contact, *errors.Error) {
	var contacts []models.Contact
	result := db.Limit(limit).Offset(offset).Find(&contacts)
	if result.Error != nil {
		return nil, &errors.FailedToGetContacts
	}

	return contacts, nil
}

func GetClubContacts(db *gorm.DB, id uuid.UUID) ([]models.Contact, *errors.Error) {
	var club models.Club
	if err := db.Preload("Contact").First(&club, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.ClubNotFound
		} else {
			return nil, &errors.FailedToGetContacts
		}
	}

	if club.Contact == nil {
		return nil, &errors.FailedToGetContacts
	}

	return club.Contact, nil
}

func GetContact(db *gorm.DB, id uuid.UUID) (*models.Contact, *errors.Error) {
	var contact models.Contact
	if err := db.First(&contact, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.ContactNotFound
		} else {
			return nil, &errors.FailedToGetContact
		}
	}

	return &contact, nil
}

func PutContact(db *gorm.DB, clubID uuid.UUID, contact models.Contact) (*models.Contact, *errors.Error) {
	// if the club already has a contact of the same type, update the existing contact
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "club_id"}, {Name: "type"}},
		DoUpdates: clause.AssignmentColumns([]string{"content"}),
	}).Create(&contact).Error

	if err != nil {

		// if the foreign key (clubID) constraint is violated, return a club not found error
		if stdliberrors.Is(err, gorm.ErrForeignKeyViolated) {
			return nil, &errors.ClubNotFound
		}
		return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToUpdateContact.Message}
	}
	return &contact, nil
}

func DeleteContact(db *gorm.DB, id uuid.UUID) *errors.Error {
	if result := db.Delete(&models.Contact{}, id); result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.ContactNotFound
		} else {
			return &errors.FailedToDeleteContact
		}
	}
	return nil
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
