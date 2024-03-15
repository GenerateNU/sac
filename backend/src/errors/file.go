package errors

import "github.com/gofiber/fiber/v2"

var (
	FailedToValidateFileId = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to validate file id",
	}
	InvalidFileSize = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "file size is greater than 5 MB",
	}
	FailedToCreateAWSSession = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create AWS session",
	}
	FailedToUpdateFile = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to update file",
	}
	FailedToUploadFile = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to upload file",
	}
	FailedToCreateFileInDB = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create file in database",
	}
	FailedToDeleteFile = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete file",
	}
	FailedToReadFile = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to read file",
	}
	FailedToValidateFile = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate file",
	}
	FailedToGetFile = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to get file",
	}
	FailedToProcessRequest = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to process the request",
	}
	FailedToValidatedData = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate data",
	}
	FailedToOpenFile = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to open file",
	}
	InvalidImageFormat = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "invalid image format",
	}
	FailedToDownloadFile = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to download the file",
	}
	InvalidAssociationType = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "invalid association type",
	}
	FailedToFindAssociationID = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to find association ID",
	}
	FailedToParseDaysToInt = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to parse days to int",
	}
	FailedToGetSignedURL = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get signed URL",
	}
	InvalidFileID = Error{
		StatusCode: fiber.StatusBadRequest, 
		Message: "invalid file id", 
	}
	CannotFindFile = Error{
		StatusCode: fiber.StatusInternalServerError, 
		Message: "unable to find file", 
	}
)
