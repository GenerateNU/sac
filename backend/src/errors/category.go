package errors

import "github.com/gofiber/fiber/v2"

var (
	FailedToValidateCategory = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate category",
	}
	FailedToCreateCategory = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create category",
	}
	FailedToGetCategory = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get category",
	}
	FailedToGetCategories = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get categories",
	}
	FailedToUpdateCategory = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to update category",
	}
	FailedToDeleteCategory = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete category",
	}
	CategoryAlreadyExists = Error{
		StatusCode: fiber.StatusConflict,
		Message:    "category already exists",
	}
	CategoryNotFound = Error{
		StatusCode: fiber.StatusNotFound,
		Message:    "category not found",
	}
)
