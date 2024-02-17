package tests

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/search"
	"github.com/h2non/gock"
	"github.com/huandu/go-assert"
)

type MockSearchableStruct struct{}

func (e *MockSearchableStruct) SearchId() string {
	return "testing_uuid"
}

func (e *MockSearchableStruct) Namespace() string {
	return "testing"
}

func (e *MockSearchableStruct) EmbeddingString() string {
	return "testing testing testing"
}

func TestPineconeUpsertWorks(t *testing.T) {
	assert := assert.New(t)

	configuration, err := config.GetConfiguration(filepath.Join("..", "..", "..", "config"))

	assert.NilError(err)

	mockSearchId := (&MockSearchableStruct{}).SearchId()
	mockSearchString := (&MockSearchableStruct{}).EmbeddingString()
	mockValues := []float32{1.0, 1.0, 1.0, 1.0}
	mockNamespace := (&MockSearchableStruct{}).Namespace()

	defer gock.Off()

	gock.New(fmt.Sprintf("%s", configuration.PineconeSettings.IndexHost.Expose())).
		Post("/vectors/upsert").
		MatchHeader("Api-Key", configuration.PineconeSettings.APIKey.Expose()).
		MatchHeader("accept", "application/json").
		MatchHeader("content-type", "application/json").
		MatchType("json").
		JSON(map[string]interface{}{
			"vectors": []interface{}{
				map[string]interface{}{
					"id":     mockSearchId,
					"values": mockValues,
				},
			},
			"namespace": mockNamespace,
		}).
		Reply(200)

	gock.New("https://api.openai.com").
		Post("/v1/embeddings").
		MatchHeader("Authorization", fmt.Sprintf("Bearer %s", configuration.OpenAISettings.APIKey.Expose())).
		MatchHeader("content-type", "application/json").
		MatchType("json").
		JSON(map[string]interface{}{
			"input": mockSearchString,
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

	client := search.NewPineconeClient(search.NewOpenAIClient(configuration.OpenAISettings), configuration.PineconeSettings)
	err = client.Upsert(&MockSearchableStruct{})
	assert.Equal(err, nil)
}

func TestPineconeDeleteWorks(t *testing.T) {
	assert := assert.New(t)

	configuration, err := config.GetConfiguration(filepath.Join("..", "..", "..", "config"))

	assert.NilError(err)

	mockSearchId := (&MockSearchableStruct{}).SearchId()
	mockNamespace := (&MockSearchableStruct{}).Namespace()

	defer gock.Off()

	gock.New(fmt.Sprintf("%s", configuration.PineconeSettings.IndexHost.Expose())).
		Post("/vectors/delete").
		MatchHeader("Api-Key", configuration.PineconeSettings.APIKey.Expose()).
		MatchHeader("accept", "application/json").
		MatchHeader("content-type", "application/json").
		MatchType("json").
		JSON(map[string]interface{}{
			"deleteAll": false,
			"ids": []string{
				mockSearchId,
			},
			"namespace": mockNamespace,
		}).
		Reply(200)

	client := search.NewPineconeClient(search.NewOpenAIClient(configuration.OpenAISettings), configuration.PineconeSettings)
	err = client.Delete(&MockSearchableStruct{})
	assert.Equal(err, nil)
}

func TestPineconeSearchWorks(t *testing.T) {
	assert := assert.New(t)

	configuration, err := config.GetConfiguration(filepath.Join("..", "..", "..", "config"))

	assert.NilError(err)

	mockSearchId := (&MockSearchableStruct{}).SearchId()
	mockSearchString := (&MockSearchableStruct{}).EmbeddingString()
	mockValues := []float32{1.0, 1.0, 1.0, 1.0}
	mockNamespace := (&MockSearchableStruct{}).Namespace()
	topK := 5

	defer gock.Off()

	gock.New(fmt.Sprintf("%s", configuration.PineconeSettings.IndexHost.Expose())).
		Post("/query").
		MatchHeader("Api-Key", configuration.PineconeSettings.APIKey.Expose()).
		MatchHeader("accept", "application/json").
		MatchHeader("content-type", "application/json").
		MatchType("json").
		JSON(map[string]interface{}{
			"includeValues":   false,
			"includeMetadata": false,
			"topK":            topK,
			"vector":          mockValues,
			"namespace":       mockNamespace,
		}).
		Reply(200).
		JSON(map[string]interface{}{
			"matches": []map[string]interface{}{
				{
					"id":     mockSearchId,
					"score":  1.0,
					"values": []float32{1.0, 1.0, 1.0, 1.0},
				},
			},
			"namespace": mockNamespace,
		})

	gock.New("https://api.openai.com").
		Post("/v1/embeddings").
		MatchHeader("Authorization", "Bearer "+configuration.OpenAISettings.APIKey.Expose()).
		MatchHeader("content-type", "application/json").
		MatchType("json").
		JSON(map[string]interface{}{
			"input": mockSearchString,
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

	client := search.NewPineconeClient(search.NewOpenAIClient(configuration.OpenAISettings), configuration.PineconeSettings)
	ids, err := client.Search(&MockSearchableStruct{}, 5)
	assert.Equal(err, nil)
	assert.Equal(len(ids), 1)
	assert.Equal(ids[0], (&MockSearchableStruct{}).SearchId())
}

func TestOpenAIEmbeddingWorks(T *testing.T) {
	assert := assert.New(T)

	configuration, err := config.GetConfiguration(filepath.Join("..", "..", "..", "config"))

	assert.NilError(err)

	testString := "test string"

	defer gock.Off()

	gock.New("https://api.openai.com").
		Post("/v1/embeddings").
		MatchHeader("Authorization", "Bearer "+configuration.OpenAISettings.APIKey.Expose()).
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

	client := search.NewOpenAIClient(configuration.OpenAISettings)
	vector, err := client.CreateEmbedding(testString)
	assert.Equal(err, nil)
	assert.Equal(vector, []float32{1.0, 1.0, 1.0, 1.0})
}
