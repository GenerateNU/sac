package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/goccy/go-json"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func TestCreateMembershipWorks(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               fmt.Sprintf("/api/v1/users/:userID/member/%s", clubUUID),
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

func TestCreateMembershipFailsClubIdNotExists(t *testing.T) {
	appAssert, _, _ := CreateSampleClub(h.InitTest(t))

	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               fmt.Sprintf("/api/v1/users/:userID/member/%s", uuid),
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
			Path:   fmt.Sprintf("/api/v1/users/%s/member/%s", uuid, clubUUID),
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
			Path:               fmt.Sprintf("/api/v1/users/:userID/member/%s", clubUUID),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		fiber.StatusCreated,
	).TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodDelete,
			Path:               fmt.Sprintf("/api/v1/users/:userID/member/%s", clubUUID),
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

// TODO: test can't work because you become a member when you create a club
// func TestDeleteMembershipNotMembership(t *testing.T) {
// 	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

// 	userClubsMemberBefore, err := transactions.GetClubMembership(appAssert.App.Conn, appAssert.App.TestUser.UUID)

// 	appAssert.Assert.Assert(err == nil)

// 	clubUsersMemberBefore, err := transactions.GetClubMembers(appAssert.App.Conn, clubUUID, 10, 0)

// 	appAssert.Assert.Assert(err == nil)

// 	appAssert.TestOnErrorAndTester(
// 		h.TestRequest{
// 			Method:             fiber.MethodDelete,
// 			Path:               fmt.Sprintf("/api/v1/users/:userID/member/%s", clubUUID),
// 			Role:               &models.Super,
// 			TestUserIDReplaces: h.StringToPointer(":userID"),
// 		},
// 		h.ErrorWithTester{
// 			Error: errors.UserNotMemberOfClub,
// 			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
// 				var user models.User

// 				err := eaa.App.Conn.Where("id = ?", eaa.App.TestUser.UUID).Preload("Member").First(&user)

// 				eaa.Assert.NilError(err)

// 				eaa.Assert.Equal(userClubsMemberBefore, user.Member)

// 				var club models.Club

// 				err = eaa.App.Conn.Where("id = ?", clubUUID).Preload("Member").First(&club)

// 				eaa.Assert.NilError(err)

// 				eaa.Assert.Equal(clubUsersMemberBefore, club.Member)
// 			},
// 		},
// 	).Close()
// }

func TestDeleteMembershipFailsClubIdNotExists(t *testing.T) {
	appAssert, _, _ := CreateSampleClub(h.InitTest(t))

	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method:             fiber.MethodDelete,
			Path:               fmt.Sprintf("/api/v1/users/:userID/member/%s", uuid),
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
			Path:   fmt.Sprintf("/api/v1/users/%s/member/%s", uuid, clubUUID),
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

func TestGetMembershipWorks(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	appAssert.TestOnStatus(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               fmt.Sprintf("/api/v1/users/:userID/member/%s", clubUUID),
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
