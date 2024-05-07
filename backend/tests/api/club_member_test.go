package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func TestCreateMembershipWorks(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               fmt.Sprintf("/api/v1/clubs/%s/members/:userID", clubUUID),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var user models.User

				err := eaa.App.Conn.Where("id = ?", eaa.App.TestUser.UUID).Preload("Member").First(&user)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(2, len(user.Member)) // SAC Super Club and the one just added

				eaa.Assert.Equal(clubUUID, user.Member[1].ID) // second club AKA the one just added

				var club models.Club

				err = eaa.App.Conn.Where("id = ?", clubUUID).Preload("Member").First(&club)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(1, len(club.Member))

				eaa.Assert.Equal(eaa.App.TestUser.UUID, club.Member[0].ID)
			},
		},
	).Close()
}

func TestClubMemberWorks(t *testing.T) {
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
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/members", clubUUID),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var members []models.User

				err := json.NewDecoder(resp.Body).Decode(&members)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(1, len(members))

				var club models.Club

				err = eaa.App.Conn.Where("id = ?", clubUUID).First(&club).Error

				eaa.Assert.NilError(err)

				var dbMembers []models.User

				err = eaa.App.Conn.Model(&club).Association("Member").Find(&dbMembers)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(len(dbMembers), len(members))
			},
		},
	).Close()
}

func TestCreateMembershipFailsClubIdNotExists(t *testing.T) {
	appAssert, _, _ := CreateSampleClub(h.InitTest(t))

	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               fmt.Sprintf("/api/v1/clubs/%s/members/:userID", uuid),
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

func TestCreateMembershipFailsUserIdNotExists(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/members/%s", clubUUID, uuid),
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

func TestDeleteMembershipWorks(t *testing.T) {
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
			Method:             fiber.MethodDelete,
			Path:               fmt.Sprintf("/api/v1/clubs/%s/members/:userID", clubUUID),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var user models.User

				err := eaa.App.Conn.Where("id = ?", eaa.App.TestUser.UUID).Preload("Member").First(&user)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(1, len(user.Member)) // SAC Super Club

				var club models.Club

				err = eaa.App.Conn.Where("id = ?", clubUUID).Preload("Member").First(&club)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(0, len(club.Member))
			},
		},
	).Close()
}

func TestDeleteMembershipFailsClubIdNotExists(t *testing.T) {
	appAssert, _, _ := CreateSampleClub(h.InitTest(t))

	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method:             fiber.MethodDelete,
			Path:               fmt.Sprintf("/api/v1/clubs/%s/members/:userID", uuid),
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

func TestDeleteMembershipFailsUserIdNotExists(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/members/%s", clubUUID, uuid),
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
