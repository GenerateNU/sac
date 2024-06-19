package search

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreateEmbeddingRequestBody struct {
	Input      []string `json:"input"`
	Model      string   `json:"model"`
	Dimensions int      `json:"dimensions"`
}

type EmbeddingResponseItem struct {
	Embedding []float32 `json:"embedding"`
}

type CreateEmbeddingResponseBody struct {
	Data []EmbeddingResponseItem `json:"data"`
}

func UpsertClubEmbedding(db *gorm.DB, s *config.SearchSettings, club *models.Club) error {
	embedding, err :=
		createEmbedding(s, fmt.Sprintf("%s %s %s", club.Name, club.Preview, club.Description))
	if err != nil {
		return err
	}

	embeddingStr := floatArrayToSql(embedding)

	queryString := fmt.Sprintf(
		"UPDATE clubs SET embedding = '[%s]' WHERE id = '%s'", embeddingStr, club.ID.String())

	if err := db.Exec(queryString).Error; err != nil {
		return err
	}

	return nil
}

func floatArrayToSql(embedding []float32) string {
	embeddingArr := make([]string, len(embedding))
	for i, f := range embedding {
		embeddingArr[i] = strconv.FormatFloat(float64(f), 'f', -1, 32)
	}
	embeddingStr := strings.Join(embeddingArr, ", ")
	return embeddingStr
}

func createEmbedding(s *config.SearchSettings, item string) ([]float32, error) {
	embeddingBody, err := json.Marshal(
		CreateEmbeddingRequestBody{
			Input:      []string{item},
			Model:      "text-embedding-3-large",
			Dimensions: 512,
		})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(fiber.MethodPost,
		"https://api.openai.com/v1/embeddings",
		bytes.NewBuffer(embeddingBody))
	if err != nil {
		return nil, err
	}

	req = utilities.ApplyModifiers(req,
		utilities.Authorization(s.OpenAIApiKey.Expose()),
		utilities.JSON(),
	)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var embeddingResultBody CreateEmbeddingResponseBody

	err = json.NewDecoder(resp.Body).Decode(&embeddingResultBody)
	if err != nil {
		return nil, err
	}

	if len(embeddingResultBody.Data) < 1 {
		return nil, err
	}

	return embeddingResultBody.Data[0].Embedding, nil
}
