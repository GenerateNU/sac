package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/search"

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

func CreateClub(db *gorm.DB, pinecone *search.PineconeClient, userId uuid.UUID, club models.Club) (*models.Club, *errors.Error) {
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

	if err := pinecone.Upsert([]search.Searchable{&club}); err != nil {

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

func UpdateClub(db *gorm.DB, pinecone *search.PineconeClient, id uuid.UUID, club models.Club) (*models.Club, *errors.Error) {
	tx := db.Begin()

	// TODO: is this block needed? (updating user based off club data)
	result := tx.Model(&models.User{}).Where("id = ?", id).Updates(club)
	if result.Error != nil {
		tx.Rollback()
		if stdliberrors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &errors.UserNotFound
		} else {
			return nil, &errors.FailedToUpdateClub
		}
	}
	var existingClub models.Club

	err := tx.First(&existingClub, id).Error
	if err != nil {
		tx.Rollback()
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.ClubNotFound
		} else {
			return nil, &errors.FailedToCreateClub
		}
	}

	if err := tx.Model(&existingClub).Updates(&club).Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToUpdateUser
	}

	if pinecone.Upsert([]search.Searchable{&existingClub}) != nil {
		tx.Rollback()
		return nil, &errors.FailedToUpsertToPinecone
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToUpdateClub
	}

	return &existingClub, nil
}

func DeleteClub(db *gorm.DB, pinecone *search.PineconeClient, id uuid.UUID) *errors.Error {
	tx := db.Begin()

	var existingClub models.Club
	err := tx.First(&existingClub, id)
	if err != nil {
		tx.Rollback()
		return &errors.ClubNotFound
	}

	pineconeErr := pinecone.Delete([]search.Searchable{&existingClub})
	if pineconeErr != nil {
		tx.Rollback()
		return &errors.FailedToDeleteClub
	}

	if result := tx.Delete(&models.Club{}, id); result.RowsAffected == 0 {
		tx.Rollback()
		if result.Error == nil {
			return &errors.ClubNotFound
		} else {
			return &errors.FailedToDeleteClub
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &errors.FailedToDeleteClub
	}

	return nil
}
