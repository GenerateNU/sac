package utilities

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrNotFound           = errors.New("not found")
	ErrDuplicate          = errors.New("duplicate")
	ErrExpectedPagination = errors.New("expected pagination. make sure to use the fiberpaginate middleware")
)

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func IsDuplicate(err error) bool {
	return errors.Is(err, ErrDuplicate)
}

func IsExpectedPagination(err error) bool {
	return errors.Is(err, ErrExpectedPagination)
}

type APIError struct {
	StatusCode int `json:"statusCode"`
	Message    any `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

func InvalidRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    errors,
	}
}

func BadRequest(err error) APIError {
	return NewAPIError(http.StatusBadRequest, err)
}

func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid JSON request data"))
}

func InvalidCookies() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid cookies"))
}

func Unauthorized() APIError {
	return NewAPIError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
}

func Forbidden() APIError {
	return NewAPIError(http.StatusForbidden, fmt.Errorf("forbidden"))
}

func Conflict() APIError {
	return NewAPIError(http.StatusConflict, fmt.Errorf("conflict"))
}

func InternalServerError() APIError {
	return NewAPIError(http.StatusInternalServerError, fmt.Errorf("internal server error"))
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	var apiErr APIError

	switch err {
	case ErrNotFound:
		apiErr = NewAPIError(http.StatusNotFound, err)
	case ErrDuplicate:
		apiErr = NewAPIError(http.StatusConflict, err)
	case ErrExpectedPagination:
		apiErr = NewAPIError(http.StatusInternalServerError, err)
	default:
		if castedErr, ok := err.(APIError); ok {
			apiErr = castedErr
		} else {
			apiErr = InternalServerError()
		}
	}

	slog.Error("HTTP API error", "err", err.Error(), "method", c.Method(), "path", c.Path())

	return c.Status(apiErr.StatusCode).JSON(apiErr)
}
