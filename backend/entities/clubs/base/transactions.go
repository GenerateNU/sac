package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/integrations/search"
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/sahilm/fuzzy"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubs(db *gorm.DB, pinecone search.SearchClientInterface, queryParams *models.ClubQueryParams, pagination middleware.Pagination) ([]models.Club, error) {
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
			return nil, err
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

	if err := query.Scopes(pagination.IntoScope(query)).Find(&clubs).Error; err != nil {
		return nil, err
	}

	return clubs, nil
}

func CreateClub(db *gorm.DB, pinecone search.SearchClientInterface, userId uuid.UUID, club models.Club) (*models.Club, error) {
	user, err := users.GetUser(db, userId)
	if err != nil {
		return nil, err
	}

	tx := db.Begin()

	club.NumMembers = 1

	if err := tx.Create(&club).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	membership := models.Membership{
		ClubID:         club.ID,
		UserID:         user.ID,
		MembershipType: models.MembershipTypeAdmin,
	}

	if err := tx.Create(&membership).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&user).Association("Follower").Append(&club); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := pinecone.Upsert([]search.Searchable{&club}); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &club, nil
}

func UpdateClub(db *gorm.DB, pinecone search.SearchClientInterface, id uuid.UUID, club models.Club) (*models.Club, error) {
	tx := db.Begin()

	var existingClub models.Club

	err := tx.First(&existingClub, id).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	if err := tx.Model(&existingClub).Updates(&club).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if pinecone.Upsert([]search.Searchable{&existingClub}) != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &existingClub, nil
}

func DeleteClub(db *gorm.DB, pinecone search.SearchClientInterface, id uuid.UUID) error {
	tx := db.Begin()

	var existingClub models.Club
	err := tx.First(&existingClub, id).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utilities.ErrNotFound
		}
		return err
	}

	if err := pinecone.Delete([]search.Searchable{&existingClub}); err != nil {
		tx.Rollback()
		return err
	}

	if result := tx.Delete(&models.Club{}, id); result.RowsAffected == 0 {
		tx.Rollback()
		if result.Error == nil {
			return utilities.ErrNotFound
		}
		return result.Error
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
