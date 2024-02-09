package types

import (
	"github.com/GenerateNU/sac/backend/src/errors"
)

// Embedding the necessary data for an embedding vector. This type is designed to mimic how Pinecone's API handles
// vectors, for easy use with it.
type Embedding struct {
	// The id this embedding should be upserted with. Note: This should be the same value as produced by
	// Embeddable.EmbeddingId(), the reason it is in both places is to simplify the upload to Pinecone code (expects
	// both id and values in the upsert payload).
	Id string `json:"id"`
	// The vector that should be upserted.
	Values []float32 `json:"values"`
}

// Embeddable Represents a value that can be transformed into an embedding vector (i.e for use in a vector database)
type Embeddable interface {
	// EmbeddingId Returns the id this embeddable value should be upserted with.
	EmbeddingId() string
	// Namespace Returns the namespace this embeddable value should be upserted to.
	Namespace() string
	// Embed Returns the embedding vector this embeddable value should be upserted as.
	Embed() (*Embedding, *errors.Error)
}
