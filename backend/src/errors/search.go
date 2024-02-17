package errors

import "github.com/gofiber/fiber/v2"

var (
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
	FailedToSearchToPinecone = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to search on pinecone",
	}
)
