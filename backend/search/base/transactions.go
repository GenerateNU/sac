package base

import (
	"fmt"
	"strings"
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/search/types"
	"gorm.io/gorm"
)

// Use this over strings.join so elements can be single quoted if nessecary.
func joinQuoted(elems []string, seperator string) string {
	elemsQuoted := make([]string, len(elems))

	for i, elem := range elems {
		elemsQuoted[i] = fmt.Sprintf("'%s'", elem)
	}

	return strings.Join(elemsQuoted, seperator)
}

func SearchClubs(db *gorm.DB, query *types.ClubSearchRequest) (*types.SearchResult[models.Club], error) {
	var clubs []models.Club

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
		whereClauses = append(whereClauses,
			fmt.Sprintf("club_tags.tag_id IN (%s)", joinQuoted(query.Tags, ",")))
	}

	finalQuery := fmt.Sprintf(
		"SELECT * FROM clubs %s WHERE %s",
		innerJoin,
		strings.Join(whereClauses, " AND "))

	if err := dbQuery.Raw(finalQuery).Scan(&clubs).Error; err != nil {
		return nil, err
	}

	// return results
	return &types.SearchResult[models.Club]{
		Results: clubs,
	}, nil
}

func SearchEvents(db *gorm.DB, query *types.EventSearchRequest) (*types.SearchResult[models.Event], error) {
	var events []models.Event

	dbQuery := db.Model(&events)

	var whereClauses []string
	var innerJoins []string

	if query.Search != "" {
		whereClauses =
			append(whereClauses,
				fmt.Sprintf("eventsearch_index_col @@ plainto_tsquery('%s')", query.Search))
	}

	if !query.StartTime.IsZero() {
		whereClauses =
			append(whereClauses,
				fmt.Sprintf("start_time >= %s", query.StartTime.Format(time.DateTime)))
	}

	if !query.EndTime.IsZero() {
		whereClauses =
			append(whereClauses,
				fmt.Sprintf("end_time <= %s", query.EndTime.Format(time.DateTime)))
	}

	if len(query.Tags) != 0 {
		innerJoins = append(innerJoins, "INNER JOIN event_tags ON events.id = event_tags.event_id")
		whereClauses = append(whereClauses,
			fmt.Sprintf("event_tags.tag_id IN (%s)", joinQuoted(query.Tags, ",")))
	}

	if len(query.Clubs) != 0 {
		innerJoins = append(innerJoins, "INNER JOIN club_events ON events.id = club_events.event_id")
		whereClauses = append(whereClauses,
			fmt.Sprintf("club_events.club_id IN (%s)", joinQuoted(query.Clubs, ",")))
	}

	finalQuery := fmt.Sprintf(
		"SELECT * FROM events %s WHERE %s",
		strings.Join(innerJoins, " "),
		strings.Join(whereClauses, " AND "))

	if err := dbQuery.Raw(finalQuery).Scan(&events).Error; err != nil {
		return nil, err
	}

	// return results
	return &types.SearchResult[models.Event]{
		Results: events,
	}, nil
}
