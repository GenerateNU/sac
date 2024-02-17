package tests

import (
	"testing"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/search"
	m "github.com/garrettladley/mattress"
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

type mockConfig struct {
	Pinecone *config.PineconeSettings
	OpenAI   *config.OpenAISettings
}

func newMockConfig() *mockConfig {
	pineconeIndexHost := "https://indexHost.com"
	pineconeApiKey := "pinecone"

	pineconeIndexHostSecret, err := m.NewSecret(pineconeIndexHost)
	if err != nil {
		panic(err)
	}

	pineconeApiKeySecret, err := m.NewSecret(pineconeApiKey)
	if err != nil {
		panic(err)
	}

	openAIApiKey := "openAI"

	openAIApiKeySecret, err := m.NewSecret(openAIApiKey)
	if err != nil {
		panic(err)
	}

	return &mockConfig{
		Pinecone: &config.PineconeSettings{
			IndexHost: pineconeIndexHostSecret,
			APIKey:    pineconeApiKeySecret,
		},
		OpenAI: &config.OpenAISettings{
			APIKey: openAIApiKeySecret,
		},
	}
}

func TestPineconeUpsertWorks(t *testing.T) {
	assert := assert.New(t)

	mockConfig := newMockConfig()

	mockSearchId := (&MockSearchableStruct{}).SearchId()
	mockSearchString := (&MockSearchableStruct{}).EmbeddingString()
	mockValues := []float32{1.0, 1.0, 1.0, 1.0}
	mockNamespace := (&MockSearchableStruct{}).Namespace()

	defer gock.Off()

	gock.New(mockConfig.Pinecone.IndexHost.Expose()).
		Post("/vectors/upsert").
		MatchHeader("Api-Key", mockConfig.Pinecone.APIKey.Expose()).
		MatchHeader("accept", "application/json").
		MatchHeader("content-type", "application/json").
		MatchType("json").
		JSON(search.PineconeUpsertRequestBody{
			Vectors: []search.Vector{
				{
					ID:     mockSearchId,
					Values: mockValues,
				},
			},
			Namespace: mockNamespace,
		}).
		Reply(200)

	gock.New("https://api.openai.com").
		Post("/v1/embeddings").
		MatchHeader("Authorization", "Bearer "+mockConfig.OpenAI.APIKey.Expose()).
		MatchHeader("content-type", "application/json").
		MatchType("json").
		JSON(search.CreateEmbeddingRequestBody{
			Input: mockSearchString,
			Model: "text-embedding-ada-002",
		}).
		Reply(200).
		JSON(search.CreateEmbeddingResponseBody{
			Data: []search.Embedding{
				{
					Embedding: mockValues,
				},
			},
		})

	client := search.NewPineconeClient(search.NewOpenAIClient(*mockConfig.OpenAI), *mockConfig.Pinecone)
	err := client.Upsert(&MockSearchableStruct{})
	assert.Equal(err, nil)
}

func TestPineconeDeleteWorks(t *testing.T) {
	assert := assert.New(t)

	mockConfig := newMockConfig()

	mockSearchId := (&MockSearchableStruct{}).SearchId()
	mockNamespace := (&MockSearchableStruct{}).Namespace()

	defer gock.Off()

	gock.New(mockConfig.Pinecone.IndexHost.Expose()).
		Post("/vectors/delete").
		MatchHeader("Api-Key", mockConfig.Pinecone.APIKey.Expose()).
		MatchHeader("accept", "application/json").
		MatchHeader("content-type", "application/json").
		MatchType("json").
		JSON(search.PineconeDeleteRequestBody{
			Namespace: mockNamespace,
			IDs:       []string{mockSearchId},
			DeleteAll: false,
		}).
		Reply(200)

	client := search.NewPineconeClient(search.NewOpenAIClient(*mockConfig.OpenAI), *mockConfig.Pinecone)
	err := client.Delete(&MockSearchableStruct{})
	assert.Equal(err, nil)
}

func TestPineconeSearchWorks(t *testing.T) {
	assert := assert.New(t)

	mockConfig := newMockConfig()

	mockSearchId := (&MockSearchableStruct{}).SearchId()
	mockSearchString := (&MockSearchableStruct{}).EmbeddingString()
	mockValues := []float32{1.0, 1.0, 1.0, 1.0}
	mockNamespace := (&MockSearchableStruct{}).Namespace()
	topK := 5

	defer gock.Off()

	gock.New(mockConfig.Pinecone.IndexHost.Expose()).
		Post("/query").
		MatchHeader("Api-Key", mockConfig.Pinecone.APIKey.Expose()).
		MatchHeader("accept", "application/json").
		MatchHeader("content-type", "application/json").
		MatchType("json").
		JSON(search.PineconeSearchRequestBody{
			IncludeValues:   false,
			IncludeMetadata: false,
			TopK:            topK,
			Vector:          mockValues,
			Namespace:       mockNamespace,
		}).
		Reply(200).
		JSON(search.PineconeSearchResponseBody{
			Matches: []search.Match{
				{
					Id:     mockSearchId,
					Score:  1.0,
					Values: mockValues,
				},
			},
			Namespace: mockNamespace,
		})

	gock.New("https://api.openai.com").
		Post("/v1/embeddings").
		MatchHeader("Authorization", "Bearer "+mockConfig.OpenAI.APIKey.Expose()).
		MatchHeader("content-type", "application/json").
		MatchType("json").
		JSON(search.CreateEmbeddingRequestBody{
			Input: mockSearchString,
			Model: "text-embedding-ada-002",
		}).
		Reply(200).
		JSON(search.CreateEmbeddingResponseBody{
			Data: []search.Embedding{
				{
					Embedding: []float32{1.0, 1.0, 1.0, 1.0},
				},
			},
		})

	client := search.NewPineconeClient(search.NewOpenAIClient(*mockConfig.OpenAI), *mockConfig.Pinecone)
	ids, err := client.Search(&MockSearchableStruct{}, 5)
	assert.Equal(err, nil)
	assert.Equal(len(ids), 1)
	assert.Equal(ids[0], (&MockSearchableStruct{}).SearchId())
}

func TestOpenAIEmbeddingWorks(t *testing.T) {
	assert := assert.New(t)

	mockConfig := newMockConfig()

	testString := "test string"

	defer gock.Off()

	gock.New("https://api.openai.com").
		Post("/v1/embeddings").
		MatchHeader("Authorization", "Bearer "+mockConfig.OpenAI.APIKey.Expose()).
		MatchHeader("content-type", "application/json").
		MatchType("json").
		JSON(search.CreateEmbeddingRequestBody{
			Input: testString,
			Model: "text-embedding-ada-002",
		}).
		Reply(200).
		JSON(search.CreateEmbeddingResponseBody{
			Data: []search.Embedding{
				{
					Embedding: []float32{1.0, 1.0, 1.0, 1.0},
				},
			},
		})

	client := search.NewOpenAIClient(*mockConfig.OpenAI)
	vector, err := client.CreateEmbedding(testString)
	assert.Equal(err, nil)
	assert.Equal(vector, []float32{1.0, 1.0, 1.0, 1.0})
}
