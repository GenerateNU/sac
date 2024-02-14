package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func TestCreateFollowingWorks(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               fmt.Sprintf("/api/v1/users/:userID/follower/%s", clubUUID),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var user models.User

				err := eaa.App.Conn.Where("id = ?", eaa.App.TestUser.UUID).Preload("Follower").First(&user)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(1, len(user.Follower))

				eaa.Assert.Equal(clubUUID, user.Follower[0].ID)

				var club models.Club

				err = eaa.App.Conn.Where("id = ?", clubUUID).Preload("Follower").First(&club)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(1, len(club.Follower))

				eaa.Assert.Equal(eaa.App.TestUser.UUID, club.Follower[0].ID)
			},
		},
	).Close()
}

func TestCreateFollowingFailsClubIdNotExists(t *testing.T) {
	appAssert, _, _ := CreateSampleClub(h.InitTest(t))

	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               fmt.Sprintf("/api/v1/users/:userID/follower/%s", uuid),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.ErrorWithTester{
			Error: errors.ClubNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var club models.Club

				err := eaa.App.Conn.Where("id = ?", uuid).First(&club).Error

				eaa.Assert.Assert(err != nil)
			},
		},
	).Close()
}

func TestCreateFollowingFailsUserIdNotExists(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", uuid, clubUUID),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error: errors.UserNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var user models.User

				err := eaa.App.Conn.Where("id = ?", uuid).First(&user).Error

				eaa.Assert.Assert(err != nil)
			},
		},
	).Close()
}

func TestDeleteFollowingWorks(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	appAssert.TestOnStatus(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               fmt.Sprintf("/api/v1/users/:userID/follower/%s", clubUUID),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		fiber.StatusCreated,
	).TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodDelete,
			Path:               fmt.Sprintf("/api/v1/users/:userID/follower/%s", clubUUID),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var user models.User

				err := eaa.App.Conn.Where("id = ?", eaa.App.TestUser.UUID).Preload("Follower").First(&user)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(0, len(user.Follower))

				var club models.Club

				err = eaa.App.Conn.Where("id = ?", clubUUID).Preload("Follower").First(&club)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(0, len(club.Follower))
			},
		},
	).Close()
}

func TestDeleteFollowingNotFollowing(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodDelete,
			Path:               fmt.Sprintf("/api/v1/users/:userID/follower/%s", clubUUID),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var user models.User

				err := eaa.App.Conn.Where("id = ?", eaa.App.TestUser.UUID).Preload("Follower").First(&user)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(0, len(user.Follower))

				var club models.Club

				err = eaa.App.Conn.Where("id = ?", clubUUID).Preload("Follower").First(&club)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(0, len(club.Follower))
			},
		},
	).Close()
}

func TestDeleteFollowingFailsClubIdNotExists(t *testing.T) {
	appAssert, _, _ := CreateSampleClub(h.InitTest(t))

	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method:             fiber.MethodDelete,
			Path:               fmt.Sprintf("/api/v1/users/:userID/follower/%s", uuid),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.ErrorWithTester{
			Error: errors.ClubNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var club models.Club

				err := eaa.App.Conn.Where("id = ?", uuid).First(&club).Error

				eaa.Assert.Assert(err != nil)
			},
		},
	).Close()
}

func TestDeleteFollowingFailsUserIdNotExists(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", uuid, clubUUID),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error: errors.UserNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var user models.User

				err := eaa.App.Conn.Where("id = ?", uuid).First(&user).Error

				eaa.Assert.Assert(err != nil)
			},
		},
	).Close()
}

func TestGetFollowingWorks(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	appAssert.TestOnStatus(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               fmt.Sprintf("/api/v1/users/:userID/follower/%s", clubUUID),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		fiber.StatusCreated,
	).TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodGet,
			Path:               "/api/v1/users/:userID/follower",
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var clubs []models.Club

				err := json.NewDecoder(resp.Body).Decode(&clubs)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(1, len(clubs))

				var dbClubs []models.Club

				err = eaa.App.Conn.Where("id = ?", clubUUID).Preload("Follower").First(&dbClubs).Error

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(1, len(clubs))
			},
		},
	).Close()
}

func TestGetFollowingFailsUserIdNotExists(t *testing.T) {
	appAssert, _, _ := CreateSampleClub(h.InitTest(t))

	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/users/%s/follower", uuid),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error: errors.UserNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var user models.User

				err := eaa.App.Conn.Where("id = ?", uuid).First(&user).Error

				eaa.Assert.Assert(err != nil)
			},
		},
	).Close()
}
