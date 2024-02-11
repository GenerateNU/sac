package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/models"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/gofiber/fiber/v2"
)

func TestCreateFollowingWorks(t *testing.T) {
	appAssert, userUUID, clubUUID := CreateSampleClub(h.InitTest(t))

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUID, clubUUID),
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var user models.User

				err := eaa.App.Conn.Where("id = ?", userUUID).Preload("Follower").First(&user)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(1, len(user.Follower))

				eaa.Assert.Equal(clubUUID, user.Follower[0].ID)

				var club models.Club

				err = eaa.App.Conn.Where("id = ?", clubUUID).Preload("Follower").First(&club)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(1, len(club.Follower))

				eaa.Assert.Equal(userUUID, club.Follower[0].ID)
			},
		},
	).Close()
}
