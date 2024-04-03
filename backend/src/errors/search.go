package errors

import "github.com/gofiber/fiber/v2"

var (
	ClubSeedingFailed = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to seed pinecone with clubs",
	}
	FailedToCreateEmbedding = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create embedding from string",
	}
	FailedToUpsertToPinecone = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to upsert to pinecone",
	}
	FailedToDeleteToPinecone = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete from pinecone",
	}
	ItemsMustHaveSameNamespace = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "items being deleted have differing namespaces",
	}
	FailedToSearchToPinecone = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to search on pinecone",
	}
)
