package root

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/search"
	"github.com/sahilm/fuzzy"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubs(db *gorm.DB, pinecone search.SearchClientInterface, queryParams *models.ClubQueryParams) ([]models.Club, *errors.Error) {
	query := db.Model(&models.Club{})

	if queryParams.Tags != nil && len(queryParams.Tags) > 0 {
		query = query.Preload("Tags")
	}

	query = query.Where(queryParams.IntoWhere())
	query = query.Where("name != ?", "SAC")

	if queryParams.Tags != nil && len(queryParams.Tags) > 0 {
		query = query.Joins("JOIN club_tags ON club_tags.club_id = clubs.id").
			Where("club_tags.tag_id IN ?", queryParams.Tags). // add search function here
			Group("clubs.id")                                 // ensure unique club records
	}

	if queryParams.Search != "" {
		clubSearch := models.NewClubSearch(queryParams.Search)
		resultIDs, err := pinecone.Search(clubSearch)
		if err != nil {
			return nil, err
		}

		query = query.Where("id IN ?", resultIDs)
	}

	var clubs models.Clubs

	if queryParams.Search != "" {
		if err := query.Find(&clubs).Error; err != nil {
			return nil, &errors.FailedToGetClubs
		}

		// len(matches) <= len(clubs), because it discardss results that don't fuzz at all.
		matches := fuzzy.FindFrom(queryParams.Search, clubs)

		// Copy the sorted, fuzzed matches to the top of the list, and record their indices
		// in a list.
		clubsSorted := models.Clubs{}
		idxMap := map[int]bool{}
		for _, match := range matches {
			clubsSorted = append(clubsSorted, clubs[match.Index])
			idxMap[match.Index] = true
		}

		// For the remaining clubs, append them to the bottom of the list, making sure not
		// to include the fuzzed clubs again.
		for i, club := range clubs {
			_, exists := idxMap[i]
			if !exists {
				clubsSorted = append(clubsSorted, club)
			}
		}

		return clubsSorted, nil
	}

	offset := (queryParams.Page - 1) * queryParams.Limit

	result := query.Limit(queryParams.Limit).Offset(offset).Find(&clubs)
	if result.Error != nil {
		return nil, &errors.FailedToGetClubs
	}

	return clubs, nil
}

func CreateClub(db *gorm.DB, pinecone search.SearchClientInterface, userId uuid.UUID, club models.Club) (*models.Club, *errors.Error) {
	user, err := users.GetUser(db, userId)
	if err != nil {
		return nil, err
	}

	tx := db.Begin()

	club.NumMembers = 1

	if err := tx.Create(&club).Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateClub
	}

	membership := models.Membership{
		ClubID:         club.ID,
		UserID:         user.ID,
		MembershipType: models.MembershipTypeAdmin,
	}

	if err := tx.Create(&membership).Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateClub
	}

	if err := tx.Model(&user).Association("Follower").Append(&club); err != nil {
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

func UpdateClub(db *gorm.DB, pinecone search.SearchClientInterface, id uuid.UUID, club models.Club) (*models.Club, *errors.Error) {
	tx := db.Begin()

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

func DeleteClub(db *gorm.DB, pinecone search.SearchClientInterface, id uuid.UUID) *errors.Error {
	tx := db.Begin()

	var existingClub models.Club
	err := tx.First(&existingClub, id).Error
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
