package errors

import "github.com/gofiber/fiber/v2"

var (
	FailedToValidateUser = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate user",
	}
	FailedToValidateUserTags = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate user tags",
	}
	FailedToCreateUser = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create user",
	}
	FailedToUpdateUser = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to update user",
	}
	FailedToGetUsers = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get users",
	}
	FailedToGetUser = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get user",
	}
	FailedToDeleteUser = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete user",
	}
	UserAlreadyExists = Error{
		StatusCode: fiber.StatusConflict,
		Message:    "user already exists",
	}
	UserNotFound = Error{
		StatusCode: fiber.StatusNotFound,
		Message:    "user not found",
	}
	FailedToComputePasswordHash = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to compute password hash",
	}
	FailedToFindUsersByEmail = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get users by email",
	}
	FailedToGetUserMemberships = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get user memberships",
	}
	UserNotMemberOfClub = Error{
		StatusCode: fiber.StatusNotFound,
		Message:    "user not member of club",
	}
	FailedToGetUserFollowing = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get user following",
	}
	UserNotFollowingClub = Error{
		StatusCode: fiber.StatusNotFound,
		Message:    "user not following club",
	}
	FailedToUpdatePassword = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to update password",
	}
)
