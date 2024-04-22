package search

// Searchable Represents a value that can be searched (i.e, able to create embeddings and upload them to vector db)
type Searchable interface {
	// SearchId Returns the id this searchable value should be associated with.
	SearchId() string
	// Namespace Returns the namespace this searchable value should be associated with.
	Namespace() string
	// EmbeddingString Returns the string that should be used to create an embedding of this searchable value.
	EmbeddingString() string
}
