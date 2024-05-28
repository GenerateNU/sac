package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"
	search_types "github.com/GenerateNU/sac/backend/search/types"
	"github.com/GenerateNU/sac/backend/transactions"
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

	req, err = http.NewRequest("POST", fmt.Sprintf("%s/_bulk", constants.SEARCH_URL), bytes.NewBufferString(requestBodyBuilder.String()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		stdout.WriteString(fmt.Sprintf("Error making _bulk request for club seeding: %s\n", err.Error()))
	}

	if resp.StatusCode != 200 {
		stdout.WriteString(fmt.Sprintf("Error making _bulk request for club seeding: response returned %d code\n", resp.StatusCode))
	}

	stdout.WriteString("Seeding clubs finished")

	return nil
}

// FIXME: see fixme comment for SeedClubs()
func SeedEvents(db *gorm.DB) error {
	stdout := os.Stdout

	stdout.WriteString("Deleting existing event index...\n")
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/events", constants.SEARCH_URL), nil)
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

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

	req, err = http.NewRequest("POST", fmt.Sprintf("%s/_bulk", constants.SEARCH_URL), bytes.NewBufferString(requestBodyBuilder.String()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		stdout.WriteString(fmt.Sprintf("Error making _bulk request for event seeding: %s\n", err.Error()))
	}

	if resp.StatusCode != 200 {
		stdout.WriteString(fmt.Sprintf("Error making _bulk request for event seeding: response returned %d code\n", resp.StatusCode))
	}

	stdout.WriteString("Seeding events finished")

	return nil
}
