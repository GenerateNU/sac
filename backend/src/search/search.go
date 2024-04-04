package search

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/garrettladley/mattress"
	"gorm.io/gorm"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/utilities"
)

type SearchClientInterface interface {
	Seed(db *gorm.DB) *errors.Error
	Upsert(items []Searchable) *errors.Error
	Delete(items []Searchable) *errors.Error
	Search(item Searchable, topK int) ([]string, *errors.Error)
}

type PineconeClient struct {
	Settings     config.PineconeSettings
	IndexName    *mattress.Secret[string]
	openAIClient AIClientInterface
	dev          bool
}

// Connects to an existing Pinecone index, using the host and keys provided in settings.
func NewPineconeClient(openAIClient AIClientInterface, settings config.PineconeSettings, dev bool) SearchClientInterface {
	return &PineconeClient{
		Settings:     settings,
		openAIClient: openAIClient,
		dev:          dev,
	}
}

// Seeds the pinecone index with the clubs currently in the database.
func (c *PineconeClient) Seed(db *gorm.DB) *errors.Error {
	if c.dev {
		return nil
	}

	var clubs []models.Club

	if err := db.Find(&clubs).Error; err != nil {
		return &errors.ClubSeedingFailed
	}

	searchables := make([]Searchable, len(clubs))
	for i := range clubs {
		searchables = append(searchables, &clubs[i])
	}

	var chunks [][]Searchable
	chunkSize := 50

	for i := 0; i < len(searchables); i += chunkSize {
		end := i + chunkSize

		if end > len(searchables) {
			end = len(searchables)
		}

		chunks = append(chunks, searchables[i:end])
	}

	for index, chunk := range chunks {
		log.Info("Uploading chunk #%d (of %d) to pinecone...\n", index+1, len(chunks))
		err := c.Upsert(chunk)
		if err != nil {
			return &errors.ClubSeedingFailed
		}
	}

	return nil
}

func pineconeRequest(settings config.PineconeSettings, req *http.Request) *http.Request {
	return utilities.ApplyModifiers(req,
		utilities.HeaderKV("Api-Key", settings.APIKey.Expose()),
		utilities.AcceptJSON(),
		utilities.JSON(),
	)
}

type Vector struct {
	ID     string    `json:"id"`
	Values []float32 `json:"values"`
}

type PineconeUpsertRequestBody struct {
	Vectors   []Vector `json:"vectors"`
	Namespace string   `json:"namespace"`
}

// Inserts the given list of searchables to the Pinecone index.
func (c *PineconeClient) Upsert(items []Searchable) *errors.Error {
	if len(items) == 0 {
		return nil
	}

	if c.dev {
		return nil
	}

	embeddings, embeddingErr := c.openAIClient.CreateEmbedding(items)
	if embeddingErr != nil {
		return &errors.FailedToUpsertToPinecone
	}

	vectors := []Vector{}
	for i, item := range items {
		vectors = append(vectors, Vector{
			ID:     item.SearchId(),
			Values: embeddings[i].Embedding,
		})
	}

	upsertBody, err := json.Marshal(
		PineconeUpsertRequestBody{
			Vectors:   vectors,
			Namespace: items[0].Namespace(),
		})
	if err != nil {
		return &errors.FailedToUpsertToPinecone
	}

	req, err := http.NewRequest(fiber.MethodPost,
		fmt.Sprintf("%s/vectors/upsert", c.Settings.IndexHost.Expose()),
		bytes.NewBuffer(upsertBody))
	if err != nil {
		return &errors.FailedToUpsertToPinecone
	}

	req = pineconeRequest(c.Settings, req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &errors.FailedToUpsertToPinecone
	}
	defer resp.Body.Close()

	if resp.StatusCode != fiber.StatusOK {
		return &errors.FailedToUpsertToPinecone
	}

	return nil
}

type PineconeDeleteRequestBody struct {
	IDs       []string `json:"ids"`
	Namespace string   `json:"namespace"`
	DeleteAll bool     `json:"deleteAll"`
}

// Deletes the given list of searchables from the Pinecone index.
func (c *PineconeClient) Delete(items []Searchable) *errors.Error {
	if len(items) == 0 {
		return nil
	}

	if c.dev {
		return nil
	}

	// Ensure all items are in the same namespace
	namespace := items[0].Namespace()
	for _, item := range items {
		if item.Namespace() != namespace {
			return &errors.ItemsMustHaveSameNamespace
		}
	}

	itemIds := make([]string, len(items))
	for i, item := range items {
		itemIds[i] = item.SearchId()
	}

	deleteBody, err := json.Marshal(
		PineconeDeleteRequestBody{
			IDs:       itemIds,
			Namespace: namespace,
			DeleteAll: false,
		})
	if err != nil {
		return &errors.FailedToDeleteToPinecone
	}

	req, err := http.NewRequest(fiber.MethodPost,
		fmt.Sprintf("%s/vectors/delete", c.Settings.IndexHost.Expose()),
		bytes.NewBuffer(deleteBody))
	if err != nil {
		return &errors.FailedToDeleteToPinecone
	}

	req = pineconeRequest(c.Settings, req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &errors.FailedToDeleteToPinecone
	}

	if resp.StatusCode != fiber.StatusOK {
		return &errors.FailedToDeleteToPinecone
	}

	return nil
}

type PineconeSearchRequestBody struct {
	IncludeValues   bool      `json:"includeValues"`
	IncludeMetadata bool      `json:"includeMetadata"`
	TopK            int       `json:"topK"`
	Vector          []float32 `json:"vector"`
	Namespace       string    `json:"namespace"`
}

type Match struct {
	Id     string    `json:"id"`
	Score  float32   `json:"score"`
	Values []float32 `json:"values"`
}

type PineconeSearchResponseBody struct {
	Matches   []Match `json:"matches"`
	Namespace string  `json:"namespace"`
}

// Runs a search on the Pinecone index given a searchable item, and returns the topK most similar
// elements' ids.
func (c *PineconeClient) Search(item Searchable, topK int) ([]string, *errors.Error) {
	if c.dev {
		return nil, nil
	}

	values, embeddingErr := c.openAIClient.CreateEmbedding([]Searchable{item})
	if embeddingErr != nil {
		return []string{}, embeddingErr
	}

	searchBody, err := json.Marshal(
		PineconeSearchRequestBody{
			IncludeValues:   false,
			IncludeMetadata: false,
			TopK:            topK,
			// Only 1 item was passed to CreateEmbedding, so grab that 1 item and use its embedding.
			Vector:    values[0].Embedding,
			Namespace: item.Namespace(),
		})
	if err != nil {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	req, err := http.NewRequest(fiber.MethodPost,
		fmt.Sprintf("%s/query", c.Settings.IndexHost.Expose()),
		bytes.NewBuffer(searchBody))
	if err != nil {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	req = pineconeRequest(c.Settings, req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	defer resp.Body.Close()

	if resp.StatusCode != fiber.StatusOK {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	var results PineconeSearchResponseBody
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	resultsToReturn := make([]string, len(results.Matches))

	for i, match := range results.Matches {
		resultsToReturn[i] = match.Id
	}

	return resultsToReturn, nil
}
