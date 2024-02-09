package errors

import "github.com/gofiber/fiber/v2"

var (
	FailedToUpsertPinecone = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to upsert to pinecone",
	}
	FailedToDeletePinecone = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete from pinecone",
	}
	FailedToSearchPinecone = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to search on pinecone",
	}
)
