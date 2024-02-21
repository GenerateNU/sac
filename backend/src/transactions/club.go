package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

// Upsert Point of Contact
func UpsertPointOfContact(db *gorm.DB, pointOfContact *models.PointOfContact) (*models.PointOfContact, *errors.Error) {
	pocExist, errPOCExist := GetPointOfContact(db, pointOfContact.ID, pointOfContact.ClubID)
	if errPOCExist != nil {
		db.Model(&pointOfContact).Where("id = ?", pocExist.ID).Association("File").Replace(pocExist, pointOfContact)
	} else {
		db.Model(&pointOfContact).Association("File").Append(pointOfContact)
	}
	// err := db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "email"}, {Name: "club_id"}},
	// 	DoUpdates: clause.AssignmentColumns([]string{"name", "email", "position"}),
	// }).Create(&pointOfContact).Error
	// if err != nil {
	// 	return nil, &errors.FailedToUpsertPointOfContact
	// }
	return pointOfContact, nil
}

// Get All Point of Contacts
func GetAllPointOfContacts(db *gorm.DB, clubId uuid.UUID) ([]models.PointOfContact, *errors.Error) {
	var pointOfContacts []models.PointOfContact
	var club models.Club

	// handles error when club id is not found
	if err := db.First(&club, clubId).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.FailedToGetAllPointOfContact
		} else {
			return nil, &errors.FailedToGetClub
		}
	} else {
		// club id is found, handle error when failed to get point of contact
		if err = db.Find(&pointOfContacts).Error; err != nil {
			return nil, &errors.FailedToGetAllPointOfContact
		}
		return pointOfContacts, nil
	}
}

// Get A Point of Contact
func GetPointOfContact(db *gorm.DB, pocId uuid.UUID, clubId uuid.UUID) (*models.PointOfContact, *errors.Error) {
	var club models.Club
	if err := db.First(&club, clubId).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.ClubNotFound
		} else {
			return nil, &errors.FailedToGetClub
		}
	} else {
		var pointOfContact models.PointOfContact
		if err := db.First(&pointOfContact, pocId).Error; err != nil {
			if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
				return nil, &errors.PointOfContactNotFound
			} else {
				return nil, &errors.FailedToGetAPointOfContact
			}
		}
		return &pointOfContact, nil
	}
}

// Delete a point of contact
func DeletePointOfContact(db *gorm.DB, pocId uuid.UUID, clubId uuid.UUID) *errors.Error {
	var deletedPointOfContact models.PointOfContact
	var club models.Club

	// handles when club id is not found
	if err := db.First(&club, clubId).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return &errors.FailedToDeletePointOfContact
		} else {
			return &errors.FailedToGetClub
		}
	} else {
		// search for point of contact's email to delete
		result := db.Where("id = ?", pocId).Delete(&deletedPointOfContact)
		if result.RowsAffected == 0 {
			if result.Error == nil {
				return &errors.PointOfContactNotFound
			} else {
				return &errors.FailedToDeletePointOfContact
			}
		}
		return nil
	}
}

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

func GetClubs(db *gorm.DB, queryParams *models.ClubQueryParams) ([]models.Club, *errors.Error) {
	query := db.Model(&models.Club{})

	if queryParams.Tags != nil && len(queryParams.Tags) > 0 {
		query = query.Preload("Tags")
	}

	for key, value := range queryParams.IntoWhere() {
		query = query.Where(key, value)
	}

	if queryParams.Tags != nil && len(queryParams.Tags) > 0 {
		query = query.Joins("JOIN club_tags ON club_tags.club_id = clubs.id").
			Where("club_tags.tag_id IN ?", queryParams.Tags).
			Group("clubs.id") // ensure unique club records
	}

	var clubs []models.Club

	offset := (queryParams.Page - 1) * queryParams.Limit

	result := query.Limit(queryParams.Limit).Offset(offset).Find(&clubs)
	if result.Error != nil {
		return nil, &errors.FailedToGetClubs
	}

	return clubs, nil
}

func CreateClub(db *gorm.DB, userId uuid.UUID, club models.Club) (*models.Club, *errors.Error) {
	user, err := GetUser(db, userId)
	if err != nil {
		return nil, err
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

func GetClub(db *gorm.DB, id uuid.UUID, preloads ...OptionalQuery) (*models.Club, *errors.Error) {
	var club models.Club

	query := db

	for _, preload := range preloads {
		query = preload(query)
	}

	if err := query.First(&club, id).Error; err != nil {
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
