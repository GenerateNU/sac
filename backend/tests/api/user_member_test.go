package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/goccy/go-json"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func TestGetMembershipWorks(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	appAssert.TestOnStatus(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               fmt.Sprintf("/api/v1/clubs/%s/members/:userID", clubUUID),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		fiber.StatusCreated,
	).TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodGet,
			Path:               "/api/v1/users/:userID/member",
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var clubs []models.Club

				err := json.NewDecoder(resp.Body).Decode(&clubs)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(2, len(clubs)) // SAC Super Club and the one just added

				var dbClubs []models.Club

				err = eaa.App.Conn.Where("id = ?", clubUUID).Preload("Member").First(&dbClubs).Error

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(2, len(clubs)) // SAC Super Club and the one just added
			},
		},
	).Close()
}

func TestGetMembershipFailsUserIdNotExists(t *testing.T) {
	appAssert, _, _ := CreateSampleClub(h.InitTest(t))

	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/users/%s/member", uuid),
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
