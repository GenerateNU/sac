package search

import (
	"fmt"
	"net/http"

	"github.com/garrettladley/mattress"
	"gorm.io/gorm"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2/log"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/GenerateNU/sac/backend/utilities"
)

type SearchClientInterface interface {
	Seed(db *gorm.DB) error
	Upsert(items []Searchable) error
	Delete(items []Searchable) error
	Search(item Searchable) ([]string, error)
}

type PineconeClient struct {
	PineconeSettings config.PineconeSettings
	IndexName        *mattress.Secret[string]
	aiClient         AIClientInterface
}

// Connects to an existing Pinecone index, using the host and keys provided in settings.
func NewPineconeClient(aiClient AIClientInterface, pineconeSettings config.PineconeSettings) *PineconeClient {
	return &PineconeClient{
		PineconeSettings: pineconeSettings,
		aiClient:         aiClient,
	}
}

// Seeds the pinecone index with the clubs currently in the database.
func (c *PineconeClient) Seed(db *gorm.DB) error {
	var clubs []models.Club

	if err := db.Find(&clubs).Error; err != nil {
		return fmt.Errorf("failed to get clubs from database: %w", err)
	}

	searchables := make([]Searchable, len(clubs))
	for i := range clubs {
		searchables[i] = &clubs[i]
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
			return fmt.Errorf("failed to seed pinecone index: %w", err)
		}
	}

	return nil
}

func pineconeModifiers(settings config.PineconeSettings) []utilities.RequestModifier {
	return []utilities.RequestModifier{
		utilities.HeaderKV("Api-Key", settings.APIKey.Expose()),
		utilities.AcceptJSON(),
		utilities.JSON(),
	}
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
func (c *PineconeClient) Upsert(items []Searchable) error {
	if len(items) == 0 {
		return nil
	}

	embeddings, embeddingErr := c.aiClient.CreateEmbedding(items)
	if embeddingErr != nil {
		return fmt.Errorf("failed to create embeddings: %w", embeddingErr)
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
		return fmt.Errorf("failed to marshal upsert body: %w", err)
	}

	resp, err := utilities.Request(http.MethodPost, fmt.Sprintf("%s/vectors/upsert", c.PineconeSettings.IndexHost.Expose()), upsertBody, pineconeModifiers(c.PineconeSettings)...)
	if err != nil {
		return fmt.Errorf("upsert request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upsert to pinecone: %s", resp.Status)
	}

	return nil
}

type PineconeDeleteRequestBody struct {
	IDs       []string `json:"ids"`
	Namespace string   `json:"namespace"`
	DeleteAll bool     `json:"deleteAll"`
}

// Deletes the given list of searchables from the Pinecone index.
func (c *PineconeClient) Delete(items []Searchable) error {
	if len(items) == 0 {
		return nil
	}

	// Ensure all items are in the same namespace
	namespace := items[0].Namespace()
	for _, item := range items {
		if item.Namespace() != namespace {
			return fmt.Errorf("all items must be in the same namespace. item %s is in namespace %s, but expected %s", item.SearchId(), item.Namespace(), namespace)
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
		return fmt.Errorf("failed to marshal delete body: %w", err)
	}

	resp, err := utilities.Request(http.MethodPost, fmt.Sprintf("%s/vectors/delete", c.PineconeSettings.IndexHost.Expose()), deleteBody, pineconeModifiers(c.PineconeSettings)...)
	if err != nil {
		return fmt.Errorf("delete request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete from pinecone: %s", resp.Status)
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
func (c *PineconeClient) Search(item Searchable) ([]string, error) {
	moderation, err := c.aiClient.CreateModeration([]Searchable{item})
	if err != nil {
		return []string{}, err
	}

	if moderation[0].Flagged {
		return []string{}, utilities.BadRequest(fmt.Errorf("potentially harmful content detected in item %s", item))
	}

	values, err := c.aiClient.CreateEmbedding([]Searchable{item})
	if err != nil {
		return []string{}, err
	}

	searchBody, err := json.Marshal(
		PineconeSearchRequestBody{
			IncludeValues:   false,
			IncludeMetadata: false,
			TopK:            constants.TOP_K_RESULTS,
			// Only 1 item was passed to CreateEmbedding, so grab that 1 item and use its embedding.
			Vector:    values[0].Embedding,
			Namespace: item.Namespace(),
		})
	if err != nil {
		return []string{}, fmt.Errorf("failed to marshal search body: %w", err)
	}

	resp, err := utilities.Request(http.MethodPost, fmt.Sprintf("%s/query", c.PineconeSettings.IndexHost.Expose()), searchBody, pineconeModifiers(c.PineconeSettings)...)
	if err != nil {
		return []string{}, fmt.Errorf("search request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return []string{}, fmt.Errorf("failed to search to pinecone: %w", err)
	}

	var results PineconeSearchResponseBody
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return []string{}, fmt.Errorf("failed to decode search response: %w", err)
	}

	resultsToReturn := []string{}

	for _, match := range results.Matches {
		if match.Score > constants.SCORE_THRESHOLD {
			resultsToReturn = append(resultsToReturn, match.Id)
		}
	}

	return resultsToReturn, nil
}
