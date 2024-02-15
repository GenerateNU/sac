package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/models"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func TestClubFollowerWorks(t *testing.T) {
	t.Parallel()
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
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/followers", clubUUID),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var followers []models.User

				err := json.NewDecoder(resp.Body).Decode(&followers)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(1, len(followers))

				var club models.Club

				err = eaa.App.Conn.Where("id = ?", clubUUID).First(&club).Error

				eaa.Assert.NilError(err)

				var dbFollowers []models.User

				err = eaa.App.Conn.Model(&club).Association("Follower").Find(&dbFollowers)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(len(dbFollowers), len(followers))
			},
		},
	).Close()
}
