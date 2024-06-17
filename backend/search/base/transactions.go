package base

import (
	"fmt"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/search/types"
	"gorm.io/gorm"
)

func SearchClubs(db *gorm.DB, query *types.ClubSearchRequest) (*types.SearchResult[models.Club], error) {
	var clubs []models.Club
	// make a giant fucking where query
	dbQuery := db.Model(&clubs)

	var whereClauses []string
	innerJoin := ""

	if query.Search != "" {
		whereClauses =
			append(whereClauses,
				fmt.Sprintf("clubsearch_index_col @@ plainto_tsquery('%s')", query.Search))
	}

	if query.MaxMembers != 0 {
		whereClauses =
			append(whereClauses,
				fmt.Sprintf("num_members <= %d", query.MaxMembers))
	}

	if query.MinMembers != 0 {
		whereClauses =
			append(whereClauses,
				fmt.Sprintf("num_members >= %d", query.MinMembers))
	}

	if len(query.Tags) != 0 {
		innerJoin = "INNER JOIN club_tags ON clubs.id = club_tags.club_id"
		concatenatedTags := ""
		for i, tag := range query.Tags {
			concatenatedTags += fmt.Sprintf("'%s'", tag)

			if i < len(query.Tags)-1 {
				concatenatedTags += ","
			}
		}

		whereClauses = append(whereClauses,
			fmt.Sprintf("club_tags.tag_id IN (%s)", concatenatedTags))
	}

	whereClausesJoined := ""
	for i, clause := range whereClauses {
		whereClausesJoined += clause

		if i < len(whereClauses)-1 {
			whereClausesJoined += " AND "
		}
	}

	finalQuery := fmt.Sprintf("SELECT * FROM clubs %s WHERE %s", innerJoin, whereClausesJoined)

	if err := dbQuery.Raw(finalQuery).Scan(&clubs).Error; err != nil {
		return nil, err
	}

	// return results
	return &types.SearchResult[models.Club]{
		Results: clubs,
	}, nil
}

func SearchEvents(db *gorm.DB, query *types.EventSearchRequest) (*types.SearchResult[models.Event], error) {
	return nil, nil
}
