package types

import (
	"github.com/GenerateNU/sac/backend/src/errors"
)

type EmbeddingResult struct {
	// The id this embedding should be upserted with.
	Id string `json:"id"`
	// The embedding vector that should be upserted.
	Embedding []float32 `json:"values"`
}

type Vectorizable interface {
	ID() string
	Namespace() string
	Vectorize() (EmbeddingResult, *errors.Error)
}
