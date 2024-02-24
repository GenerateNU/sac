package search

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/garrettladley/mattress"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/utilities"

	stdliberrors "errors"
)

type PineconeClientInterface interface {
	Upsert(item Searchable) *errors.Error
	Delete(item Searchable) *errors.Error
	Search(item Searchable, topK int) ([]string, *errors.Error)
}

type PineconeClient struct {
	Settings config.PineconeSettings
	// FIXME: this is needed to delete an index for dev client, nicer way to do this though?
	IndexName    *mattress.Secret[string]
	openAIClient *OpenAIClient
}

func NewPineconeClient(openAIClient *OpenAIClient, settings config.PineconeSettings) *PineconeClient {
	return &PineconeClient{
		Settings:     settings,
		openAIClient: openAIClient,
	}
}

type PineconePodRequest struct {
	Environment string `json:"environment"`
	PodType     string `json:"pod_type"`
}

type PineconeSpecRequest struct {
	Pod PineconePodRequest `json:"pod"`
}

type PineconeCreateIndexRequestBody struct {
	Name      string              `json:"name"`
	Dimension int32               `json:"dimension"`
	Cosine    string              `json:"metric"`
	Spec      PineconeSpecRequest `json:"spec"`
}

type PineconeCreateIndexResponseBody struct {
	Host string `json:"host"`
}

// FIXME: eeuuuuhhhh
func NewPineconeDevelopmentClient(openAIClient *OpenAIClient, settings config.PineconeSettings) (*PineconeClient, error) {
	// Create a new index
	newIndexUUID, err := uuid.NewUUID()
	if err != nil {
		print("UUID failed")
		return nil, err
	}
	newIndexName := fmt.Sprintf("dev-%s", newIndexUUID.String())
	createIndexBody, err := json.Marshal(
		PineconeCreateIndexRequestBody{
			Name:      newIndexName,
			Dimension: 1536,
			Cosine:    "cosine",
			Spec: PineconeSpecRequest{
				Pod: PineconePodRequest{
					Environment: "gcp-starter",
					PodType:     "p1.x1",
				},
			},
		})
	if err != nil {
		print("Json marshaling failed")
		return nil, err
	}

	req, err := http.NewRequest(fiber.MethodPost,
		"https://api.pinecone.io/indexes",
		bytes.NewBuffer(createIndexBody))
	if err != nil {
		print("Request creation failed")
		return nil, err
	}

	req = utilities.ApplyModifiers(req,
		utilities.HeaderKV("Api-Key", settings.APIKey.Expose()),
		utilities.AcceptJSON(),
		utilities.JSON())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		print("Sending request failed")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != fiber.StatusCreated {
		print("Creating pinecone index did not return right status code.")
		return nil, nil
	}

	var body PineconeCreateIndexResponseBody
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		print("JSON decoding failed")
		return nil, err
	}

	indexHostSecret, err := mattress.NewSecret(fmt.Sprintf("https://%s", body.Host))
	indexNameSecret, err := mattress.NewSecret(newIndexName)
	return &PineconeClient{
		Settings: config.PineconeSettings{
			IndexHost: indexHostSecret,
			APIKey:    settings.APIKey,
		},
		IndexName:    indexNameSecret,
		openAIClient: openAIClient,
	}, nil
}

type PineconeDeleteIndexRequestBody struct {
	IndexName string `json:"index_name"`
}

// FIXME: euhhhh, needs code review
func (c *PineconeClient) DeletePineconeDevelopmentClient() error {
	req, err := http.NewRequest(fiber.MethodDelete,
		fmt.Sprintf("https://api.pinecone.io/indexes/%s", c.IndexName.Expose()),
		nil)
	if err != nil {
		return err
	}

	req = utilities.ApplyModifiers(req,
		utilities.HeaderKV("Api-Key", c.Settings.APIKey.Expose()),
		utilities.AcceptJSON(),
		utilities.JSON())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// FIXME: what do we do here
	if resp.StatusCode != fiber.StatusAccepted {
		return nil
	}

	return nil
}

// FIXME: euuhhhhh
func (c *PineconeClient) Seed(db *gorm.DB) error {
	var clubs []models.Club

	if err := db.Find(&clubs).Error; err != nil {
		return err
	}

	for _, club := range clubs {
		print(fmt.Sprintf("Uploading club %s to pinecone...\n", club.ID))
		err := c.Upsert(&club)
		if err != nil {
			return stdliberrors.New("Club upsert failed...")
		}
	}

	return nil
}

func (c *PineconeClient) pineconeRequest(req *http.Request) *http.Request {
	return utilities.ApplyModifiers(req,
		utilities.HeaderKV("Api-Key", c.Settings.APIKey.Expose()),
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

// TODO: get rid of print debug
func (c *PineconeClient) Upsert(item Searchable) *errors.Error {
	values, embeddingErr := c.openAIClient.CreateEmbedding(item.EmbeddingString())
	if embeddingErr != nil {
		print("OpenAI embedding failed")
		return &errors.FailedToUpsertToPinecone
	}

	upsertBody, err := json.Marshal(
		PineconeUpsertRequestBody{
			Vectors: []Vector{
				{
					ID:     item.SearchId(),
					Values: values,
				},
			},
			Namespace: item.Namespace(),
		})
	if err != nil {
		print("JSON marshaling failed")
		return &errors.FailedToUpsertToPinecone
	}

	req, err := http.NewRequest(fiber.MethodPost,
		fmt.Sprintf("%s/vectors/upsert", c.Settings.IndexHost.Expose()),
		bytes.NewBuffer(upsertBody))
	if err != nil {
		print("Failed to create request")
		return &errors.FailedToUpsertToPinecone
	}

	req = c.pineconeRequest(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		print("Failed to send request")
		return &errors.FailedToUpsertToPinecone
	}
	defer resp.Body.Close()

	if resp.StatusCode != fiber.StatusOK {
		println("Pinecone returned a", resp.StatusCode, "status code")
		respBodyBytes, _ := io.ReadAll(resp.Body)
		respBody := string(respBodyBytes)
		println("Response body:\n", respBody)

		return &errors.FailedToUpsertToPinecone
	}

	return nil
}

type PineconeDeleteRequestBody struct {
	IDs       []string `json:"ids"`
	Namespace string   `json:"namespace"`
	DeleteAll bool     `json:"deleteAll"`
}

func NewPineconeDeleteRequestBody(ids []string, namespace string, deleteAll bool) *PineconeDeleteRequestBody {
	return &PineconeDeleteRequestBody{
		IDs:       ids,
		Namespace: namespace,
		DeleteAll: deleteAll,
	}
}

func (c *PineconeClient) Delete(item Searchable) *errors.Error {
	deleteBody, err := json.Marshal(
		PineconeDeleteRequestBody{
			IDs:       []string{item.SearchId()},
			Namespace: item.Namespace(),
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

	req = c.pineconeRequest(req)

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

func (c *PineconeClient) Search(item Searchable, topK int) ([]string, *errors.Error) {
	values, embeddingErr := c.openAIClient.CreateEmbedding(item.EmbeddingString())
	if embeddingErr != nil {
		return []string{}, embeddingErr
	}

	searchBody, err := json.Marshal(
		PineconeSearchRequestBody{
			IncludeValues:   false,
			IncludeMetadata: false,
			TopK:            topK,
			Vector:          values,
			Namespace:       item.Namespace(),
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

	req = c.pineconeRequest(req)

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
