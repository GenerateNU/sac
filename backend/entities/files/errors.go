package errors

import (
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/gofiber/fiber/v2"
)

var (
	FailedToValidateFileId = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to validate file id",
	}
	InvalidFileSize = errors.Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "file size is greater than 5 MB",
	}
	InvalidFileType = errors.Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "invalid file type",
	}
	FileNotFound = errors.Error{
		StatusCode: fiber.StatusNotFound,
		Message:    "file not found",
	}
	FailedToCreateAWSSession = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create AWS session",
	}
	FailedToUpdateFile = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to update file",
	}
	FailedToUploadFile = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to upload file",
	}
	FailedToCreateFileInDB = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create file in database",
	}
	FailedToCreateFile = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create file",
	}
	FailedToDeleteFile = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete file",
	}
	FailedToReadFile = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to read file",
	}
	FailedToValidateFile = errors.Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate file",
	}
	FailedToGetFiles = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get files",
	}
	FailedToGetFile = errors.Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to get file",
	}
	FailedToProcessRequest = errors.Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to process the request",
	}
	FailedToValidatedData = errors.Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate data",
	}
	FailedToOpenFile = errors.Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to open file",
	}
	InvalidImageFormat = errors.Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "invalid image format",
	}
	FailedToDownloadFile = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to download the file",
	}
	InvalidAssociationType = errors.Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "invalid association type",
	}
	FailedToFindAssociationID = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to find association ID",
	}
	FailedToParseDaysToInt = errors.Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to parse days to int",
	}
	FailedToGetSignedURL = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get signed URL",
	}
	InvalidFileID = errors.Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "invalid file id",
	}
	CannotFindFile = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "unable to find file",
	}
)
