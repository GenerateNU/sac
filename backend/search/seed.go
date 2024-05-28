package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"
	search_types "github.com/GenerateNU/sac/backend/search/types"
	"github.com/GenerateNU/sac/backend/transactions"
	"github.com/GenerateNU/sac/backend/utilities"
	"gorm.io/gorm"
)

// FIXME: For now it is performant enough to just delete the index and redo, but in the future
// we may have to start worrying about updating in place for this, maybe a migration-esque strategy as well. (???)
func SeedClubs(db *gorm.DB) error {
	stdout := os.Stdout

	stdout.WriteString("Deleting existing club index...\n")
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/clubs", constants.SEARCH_URL), nil)
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	var clubs []models.Club

	if err := transactions.PreloadTag()(db).Find(&clubs).Error; err != nil {
		return err
	}

	var requestBodyBuilder strings.Builder

	for _, club := range clubs {
		indexData := search_types.BulkRequestCreate{}
		indexData.Create.Index = "clubs"
		indexData.Create.Id = club.ID.String()
		indexJson, err := json.Marshal(indexData)
		if err != nil {
			return err
		}

		clubData := club.ToSearchDocument()
		clubJson, err := json.Marshal(clubData)
		if err != nil {
			return err
		}

		requestBodyBuilder.WriteString(fmt.Sprintf("%s\n%s\n", indexJson, clubJson))
	}

	resp, err := utilities.Request(http.MethodPost, fmt.Sprintf("%s/_bulk", constants.SEARCH_URL), []byte(requestBodyBuilder.String()), utilities.JSON())
	if err != nil {
		stdout.WriteString(fmt.Sprintf("Error making _bulk request for club seeding: %s\n", err.Error()))
	}
	defer resp.Body.Close()

	if !utilities.IsOk(resp.StatusCode) {
		stdout.WriteString(fmt.Sprintf("Error making _bulk request for club seeding: response returned %d code\n", resp.StatusCode))
	}

	stdout.WriteString("Seeding clubs finished")

	return nil
}

// FIXME: see fixme comment for SeedClubs()
func SeedEvents(db *gorm.DB) error {
	stdout := os.Stdout

	stdout.WriteString("Deleting existing event index...\n")
	resp, err := utilities.Request(http.MethodDelete, fmt.Sprintf("%s/events", constants.SEARCH_URL), nil, utilities.JSON())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var events []models.Event

	if err := db.Preload("Tag").Preload("Clubs").Not("is_draft = ?", true).Not("is_archived = ?", true).Find(&events).Error; err != nil {
		return err
	}

	var requestBodyBuilder strings.Builder

	for _, event := range events {
		indexData := search_types.BulkRequestCreate{}
		indexData.Create.Index = "events"
		indexData.Create.Id = event.ID.String()
		indexJson, err := json.Marshal(indexData)
		if err != nil {
			return err
		}

		eventData := event.ToSearchDocument()
		eventJson, err := json.Marshal(eventData)
		if err != nil {
			return err
		}

		requestBodyBuilder.WriteString(fmt.Sprintf("%s\n%s\n", indexJson, eventJson))
	}

	resp, err = utilities.Request(http.MethodPost, fmt.Sprintf("%s/_bulk", constants.SEARCH_URL), []byte(requestBodyBuilder.String()), utilities.JSON())
	if err != nil {
		stdout.WriteString(fmt.Sprintf("Error making _bulk request for event seeding: %s\n", err.Error()))
	}
	defer resp.Body.Close()

	if !utilities.IsOk(resp.StatusCode) {
		stdout.WriteString(fmt.Sprintf("Error making _bulk request for event seeding: response returned %d code\n", resp.StatusCode))
	}

	stdout.WriteString("Seeding events finished")

	return nil
}
