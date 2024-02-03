package tests

import (
	stdliberrors "errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
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

func TestDeleteUserFollowingWorks(t *testing.T) {
	appAssert, userUUID, clubUUID := CreateSampleClub(t, nil)

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUID, clubUUID),
	}.TestOnStatus(t, &appAssert, fiber.StatusCreated)
	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUID, clubUUID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusCreated,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var dbUser models.User
				err := app.Conn.Preload("Follower").First(&dbUser, userUUID).Error
				assert.NilError(err)

				assert.Equal(len(dbUser.Follower), 0)
			},
		},
	)
	appAssert.Close()
}

func TestDeleteUserFollowingFailsClubIdBadRequest(t *testing.T) {
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
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUID, badRequest),
		}.TestOnError(t, &appAssert, errors.FailedToValidateID).Close()
	}
}

func TestDeleteUserFollowingFailsUserIdBadRequest(t *testing.T) {
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
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", badRequest, clubUUID),
		}.TestOnError(t, &appAssert, errors.FailedToValidateID).Close()
	}
}


func TestDeleteUserFollowingFailsUserNotExist(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(t, nil)
	userUUIDNotExist := uuid.New()

	TestRequest{
		Method: fiber.MethodDelete,
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

func TestDeleteUserFollowingFailsClubNotExist(t *testing.T) {
	appAssert, userUUID := CreateSampleUser(t, nil)
	clubUUIDNotExist := uuid.New()

	TestRequest{
		Method: fiber.MethodDelete,
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

func TestGetUserFollowingWorks(t *testing.T) {
	appAssert, userUUID, clubUUID := CreateSampleClub(t, nil)

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUID, clubUUID),
	}.TestOnStatus(t, &appAssert, fiber.StatusCreated)
	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/users/%s/follower", userUUID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusCreated,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var dbUser *models.User
				err := app.Conn.Preload("Follower").First(dbUser, userUUID).Error
				assert.NilError(err)
				assert.Equal(len(dbUser.Follower), 1)

				var club *models.Club
				err = app.Conn.First(club, clubUUID).Error
				assert.NilError(err)

				club, _ = transactions.GetClub(app.Conn, clubUUID)
				clubFollowed := &dbUser.Follower[0]
				assert.Equal(clubFollowed, club)
			},
		},
	)
	appAssert.Close()
}


func TestGetUserFailsUserNotExist(t *testing.T) {
	appAssert, userUUID, clubUUID := CreateSampleClub(t, nil)
	userUUIDNotExist := uuid.New()

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUID, clubUUID),
	}.TestOnStatus(t, &appAssert, fiber.StatusCreated)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/users/%s/follower", userUUIDNotExist),
	}.TestOnErrorAndDB(t, &appAssert,
		ErrorWithDBTester{
			Error: errors.ClubNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var user models.User
				err := app.Conn.Where("id = ?", userUUIDNotExist).First(&user).Error
				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestGetUserFailsUserIdBadRequest(t *testing.T) {
	appAssert, userUUID, clubUUID := CreateSampleClub(t, nil)

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUID, clubUUID),
	}.TestOnStatus(t, &appAssert, fiber.StatusCreated)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}
	for _, badRequest := range badRequests {
		TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/users/%s/follower", badRequest),
		}.TestOnError(t, &appAssert, errors.FailedToValidateID).Close()
	}
}