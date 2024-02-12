package embeddings

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GenerateNU/sac/backend/src/errors"
	"net/http"
	"os"
)

func CreateEmbeddingVector(infoForEmbedding string) ([]float32, *errors.Error) {
	apiKey := os.Getenv("SAC_OPENAI_API_KEY")

	InfoPayload := map[string]interface{}{
		"input": infoForEmbedding,
		"model": "text-embedding-ada-002",
	}

	InfoBody, _ := json.Marshal(InfoPayload)
	requestInfoBody := bytes.NewBuffer(InfoBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.openai.com/v1/embeddings"), requestInfoBody)
	if err != nil {
		return nil, &errors.FailedToVectorizeClub
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, &errors.FailedToVectorizeClub
	}

	defer resp.Body.Close()

	if err != nil {
		return nil, &errors.FailedToVectorizeClub
	}

	type ResponseBody struct {
		Data []struct {
			Embedding []float32 `json:"embedding"`
		} `json:"data"`
	}

	embeddingResultBody := ResponseBody{}
	err = json.NewDecoder(resp.Body).Decode(&embeddingResultBody)
	if err != nil {
		return nil, &errors.FailedToVectorizeClub
	}

	if len(embeddingResultBody.Data) < 1 {
		return nil, &errors.FailedToVectorizeClub
	}

	return embeddingResultBody.Data[0].Embedding, nil
}
