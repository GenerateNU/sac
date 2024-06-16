package utilities

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrDuplicate        = errors.New("duplicate")
	ErrExpectedPageInfo = errors.New("expected page info. make sure to use the fiberpaginate middleware")
)

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func IsDuplicate(err error) bool {
	return errors.Is(err, ErrDuplicate)
}

func IsExpectedPageInfo(err error) bool {
	return errors.Is(err, ErrExpectedPageInfo)
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
	return NewAPIError(http.StatusBadRequest, errors.New("invalid JSON request data"))
}

func Unauthorized() APIError {
	return NewAPIError(http.StatusUnauthorized, errors.New("unauthorized"))
}

func Forbidden() APIError {
	return NewAPIError(http.StatusForbidden, errors.New("forbidden"))
}

func ForbiddenReason(err error) APIError {
	return NewAPIError(http.StatusForbidden, err)
}

func Conflict() APIError {
	return NewAPIError(http.StatusConflict, errors.New("conflict"))
}

func InternalServerError() APIError {
	return NewAPIError(http.StatusInternalServerError, errors.New("internal server error"))
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	var apiErr APIError

	switch err {
	case ErrNotFound:
		apiErr = NewAPIError(http.StatusNotFound, err)
	case ErrDuplicate:
		apiErr = NewAPIError(http.StatusConflict, err)
	case ErrExpectedPageInfo:
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
