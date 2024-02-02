package tests

import (
	stdliberrors "errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/huandu/go-assert"
	"gorm.io/gorm"
)

func TestCreateUserFollowingWorks(t *testing.T) {
	appAssert, userUUID, clubUUID := CreateSampleClub(t, nil)

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUID, clubUUID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusCreated,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var dbUser models.User
				err := app.Conn.Preload("Follower").First(&dbUser, userUUID).Error
				assert.NilError(err)

				assert.Equal(len(dbUser.Follower), 1)
			},
		},
	)
	appAssert.Close()
}

func TestCreateUserFollowingFailsClubIdBadRequest(t *testing.T) {
	appAssert, userUUID := CreateSampleUser(t, nil)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		TestRequest{
			Method: fiber.MethodPut,
			Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUID, badRequest),
		}.TestOnError(t, &appAssert, errors.FailedToValidateID).Close()
	}
}

func TestCreateUserFollowingFailsUserIdBadRequest(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(t, nil)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		TestRequest{
			Method: fiber.MethodPut,
			Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", badRequest, clubUUID),
		}.TestOnError(t, &appAssert, errors.FailedToValidateID).Close()
	}
}

func TestCreateUserFollowingFailsUserNotExist(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(t, nil)
	userUUIDNotExist := uuid.New()

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUIDNotExist, clubUUID),
	}.TestOnErrorAndDB(t, &appAssert,
		ErrorWithDBTester{
			Error: errors.UserNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var user models.User
				err := app.Conn.Where("id = ?", userUUIDNotExist).First(&user).Error
				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestCreateUserFollowingFailsClubNotExist(t *testing.T) {
	appAssert, userUUID := CreateSampleUser(t, nil)
	clubUUIDNotExist := uuid.New()

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUID, clubUUIDNotExist),
	}.TestOnErrorAndDB(t, &appAssert,
		ErrorWithDBTester{
			Error: errors.ClubNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var club models.Club
				err := app.Conn.Where("id = ?", clubUUIDNotExist).First(&club).Error
				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}
