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


func TestGetClubFollowersWorks(t *testing.T) {
	appAssert, userUUID, clubUUID := CreateSampleClub(t, nil)

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUID, clubUUID),
	}.TestOnStatus(t, &appAssert, fiber.StatusCreated)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/follower", clubUUID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusCreated,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var dbClub *models.Club
				err := app.Conn.Preload("Follower").First(dbClub, clubUUID).Error
				assert.NilError(err)
				assert.Equal(len(dbClub.Follower), 1)

				var user *models.User
				err = app.Conn.First(user, userUUID).Error
				assert.NilError(err)

				user, _ = transactions.GetUser(app.Conn, userUUID)
				userFollower := &dbClub.Follower[0]
				assert.Equal(userFollower, user)
			},
		},
	)
	appAssert.Close()
}


func TestGetClubFollowersFailsClubNotExist(t *testing.T) {
	appAssert, userUUID, clubUUID := CreateSampleClub(t, nil)
	clubUUIDNotExist := uuid.New()

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUID, clubUUID),
	}.TestOnStatus(t, &appAssert, fiber.StatusCreated)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/follower", clubUUIDNotExist),
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

func TestGetClubFollowersFailsClubIdBadRequest(t *testing.T) {
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
			Path:   fmt.Sprintf("/api/v1/clubs/%s/follower", badRequest),
		}.TestOnError(t, &appAssert, errors.FailedToValidateID).Close()
	}
}