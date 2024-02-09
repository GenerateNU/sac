package tests

import (
	"fmt"
	"github.com/GenerateNU/sac/backend/src/embeddings"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/h2non/gock"
	"github.com/huandu/go-assert"
	"os"
	"testing"
)

type MockEmbeddingStruct struct{}

func (e *MockEmbeddingStruct) EmbeddingId() string {
	return "testing_uuid"
}

func (e *MockEmbeddingStruct) Namespace() string {
	return "testing"
}

func (e *MockEmbeddingStruct) Embed() (*types.Embedding, *errors.Error) {
	return &types.Embedding{
		Id:     "testing_uuid",
		Values: []float32{1.0, 1.0, 1.0, 1.0},
	}, nil
}

func TestPineconeUpsertWorks(t *testing.T) {
	assert := assert.New(t)
	indexHost := os.Getenv("SAC_PINECONE_INDEX_HOST")
	apiKey := os.Getenv("SAC_PINECONE_API_KEY")

	dummyEmbed, _ := (&MockEmbeddingStruct{}).Embed()

	defer gock.Off()

	gock.New(fmt.Sprintf("%s", indexHost)).
		Post("/vectors/upsert").
		MatchHeader("Api-Key", apiKey).
		MatchHeader("accept", "application/json").
		MatchHeader("content-type", "application/json").
		MatchType("json").
		JSON(map[string]interface{}{
			"vectors": []interface{}{
				*dummyEmbed,
			},
			"namespace": (&MockEmbeddingStruct{}).Namespace(),
		}).
		Reply(200)

	err := embeddings.UpsertToPinecone(&MockEmbeddingStruct{})
	assert.Equal(err, nil)
}

func TestPineconeDeleteWorks(t *testing.T) {
	assert := assert.New(t)

	indexHost := os.Getenv("SAC_PINECONE_INDEX_HOST")
	apiKey := os.Getenv("SAC_PINECONE_API_KEY")

	defer gock.Off()

	gock.New(fmt.Sprintf("%s", indexHost)).
		Post("/vectors/delete").
		MatchHeader("Api-Key", apiKey).
		MatchHeader("accept", "application/json").
		MatchHeader("content-type", "application/json").
		MatchType("json").
		JSON(map[string]interface{}{
			"deleteAll": false,
			"ids": []string{
				(&MockEmbeddingStruct{}).EmbeddingId(),
			},
			"namespace": (&MockEmbeddingStruct{}).Namespace(),
		}).
		Reply(200)

	err := embeddings.DeleteFromPinecone(&MockEmbeddingStruct{})
	assert.Equal(err, nil)
}

func TestPineconeSearchWorks(t *testing.T) {
	assert := assert.New(t)

	indexHost := os.Getenv("SAC_PINECONE_INDEX_HOST")
	apiKey := os.Getenv("SAC_PINECONE_API_KEY")

	dummyEmbed, _ := (&MockEmbeddingStruct{}).Embed()
	topK := 5
	defer gock.Off()

	gock.New(fmt.Sprintf("%s", indexHost)).
		Post("/query").
		MatchHeader("Api-Key", apiKey).
		MatchHeader("accept", "application/json").
		MatchHeader("content-type", "application/json").
		MatchType("json").
		JSON(map[string]interface{}{
			"includeValues":   false,
			"includeMetadata": false,
			"topK":            topK,
			"vector":          dummyEmbed.Values,
			"namespace":       (&MockEmbeddingStruct{}).Namespace(),
		}).
		Reply(200).
		JSON(map[string]interface{}{
			"matches": []map[string]interface{}{
				{
					"id":     (&MockEmbeddingStruct{}).EmbeddingId(),
					"score":  1.0,
					"values": []float32{1.0, 1.0, 1.0, 1.0},
				},
			},
			"namespace": (&MockEmbeddingStruct{}).Namespace(),
		})

	ids, err := embeddings.SearchPinecone(&MockEmbeddingStruct{}, 5)
	assert.Equal(err, nil)
	assert.Equal(len(ids), 1)
	assert.Equal(ids[0], (&MockEmbeddingStruct{}).EmbeddingId())
}

func TestOpenAIEmbeddingWorks(T *testing.T) {
	assert := assert.New(T)

	apiKey := os.Getenv("SAC_OPENAI_API_KEY")
	testString := "test string"

	defer gock.Off()

	gock.New("https://api.openai.com").
		Post("/v1/embeddings").
		MatchHeader("Authorization", fmt.Sprintf("Bearer %s", apiKey)).
		MatchHeader("content-type", "application/json").
		MatchType("json").
		JSON(map[string]interface{}{
			"input": testString,
			"model": "text-embedding-ada-002",
		}).
		Reply(200).
		JSON(map[string]interface{}{
			"data": []map[string]interface{}{
				{
					"embedding": []float32{1.0, 1.0, 1.0, 1.0},
				},
			},
		})

	vector, err := embeddings.CreateEmbeddingVector(testString)
	assert.Equal(err, nil)
	assert.Equal(vector, []float32{1.0, 1.0, 1.0, 1.0})
}
