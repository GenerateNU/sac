package tests

import (
	"fmt"
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/huandu/go-assert"
	"net/http"
	"testing"
)

func CreateSampleUser2(t *testing.T, existingAppAssert *ExistingAppAssert) (ExistingAppAssert, uuid.UUID) {
	var uuid uuid.UUID

	body := SampleUserFactory()
	(*body)["nuid"] = "012820050"
	(*body)["email"] = "brennan.mic@northeastern.edu"

	newAppAssert := TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/users/",
		Body:   body,
	}.TestOnStatusAndDB(t, existingAppAssert,
		DBTesterWithStatus{
			Status: fiber.StatusCreated,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respUser models.User

				err := json.NewDecoder(resp.Body).Decode(&respUser)

				assert.NilError(err)

				var dbUser models.User

				err = app.Conn.Where("nuid = ?", "012820050").First(&dbUser).Error

				assert.NilError(err)

				assert.Equal(dbUser.FirstName, respUser.FirstName)
				assert.Equal(dbUser.LastName, respUser.LastName)
				assert.Equal(dbUser.Email, respUser.Email)
				assert.Equal(dbUser.NUID, respUser.NUID)
				assert.Equal(dbUser.College, respUser.College)
				assert.Equal(dbUser.Year, respUser.Year)

				match, err := auth.ComparePasswordAndHash((*body)["password"].(string), dbUser.PasswordHash)

				assert.NilError(err)

				assert.Assert(match)

				assert.Equal((*body)["first_name"].(string), dbUser.FirstName)
				assert.Equal((*body)["last_name"].(string), dbUser.LastName)
				assert.Equal((*body)["email"].(string), dbUser.Email)
				assert.Equal((*body)["nuid"].(string), dbUser.NUID)
				assert.Equal(models.College((*body)["college"].(string)), dbUser.College)
				assert.Equal(models.Year((*body)["year"].(int)), dbUser.Year)

				uuid = dbUser.ID
			},
		},
	)

	if existingAppAssert == nil {
		return newAppAssert, uuid
	} else {
		return *existingAppAssert, uuid
	}

}

// Creates a single club with a single member for testing
func CreateSampleClubWithMembership(t *testing.T, existingAppAssert *ExistingAppAssert) (eaa ExistingAppAssert, userUUID uuid.UUID, clubUUID uuid.UUID) {
	appAssert, _, clubID := CreateSampleClub(t, existingAppAssert)
	appAssert2, userID := CreateSampleUser2(t, &appAssert)

	newAppAssert := TestRequest{
		Method: fiber.MethodPost,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/membership/%s", clubID, userID),
		Body:   &map[string]interface{}{},
	}.TestOnStatusAndDB(t, &appAssert2,
		DBTesterWithStatus{
			Status: fiber.StatusNoContent,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				club := models.Club{}
				err := app.Conn.Where("id = ?", clubID).First(&club).Error

				assert.NilError(err)

				var members []models.User
				err = app.Conn.Model(&club).Association("Member").Find(&members)

				assert.NilError(err)
				assert.Equal(len(members), 1)
				assert.Equal(members[0].ID, userID)
			},
		})

	if existingAppAssert == nil {
		return newAppAssert, userUUID, clubUUID
	} else {
		return *existingAppAssert, userUUID, clubUUID
	}
}

func AssertNumClubMembersRemainsAtN(app TestApp, assert *assert.A, clubID uuid.UUID, n int) {
	var club models.Club
	var members []models.User
	err := app.Conn.Where("id = ?", clubID).First(&club).Association("Members").Find(&members).Error

	assert.NilError(err)
	assert.Equal(n, len(members))
}

func AssertNumClubMembersRemainsAt0(app TestApp, assert *assert.A, clubID uuid.UUID) {
	AssertNumClubMembersRemainsAtN(app, assert, clubID, 0)
}

func AssertNumClubMembersRemainsAt1(app TestApp, assert *assert.A, clubID uuid.UUID) {
	AssertNumClubMembersRemainsAtN(app, assert, clubID, 1)
}

// Create membership by user id - 201 OK
func TestCreateMembershipWorks(t *testing.T) {
	appAssert, _, _ := CreateSampleClubWithMembership(t, nil)
	appAssert.Close()
}

// Create membership by user id - 404 not found (bad club id)
func TestCreateMembershipFailsOnInvalidClubId(t *testing.T) {
	appAssert, userID := CreateSampleUser(t, nil)

	TestRequest{
		Method: fiber.MethodPost,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/membership/%s", "gobbledygook", userID),
		Body:   &map[string]interface{}{},
	}.TestOnError(t, &appAssert, errors.FailedToValidateID)
}

// Create membership by user id - 404 not found (bad user id)
func TestCreateMembershipFailsOnInvalidUserId(t *testing.T) {
	appAssert, _, clubID := CreateSampleClub(t, nil)

	TestRequest{
		Method: fiber.MethodPost,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/membership/%s", clubID, "gobbledygook"),
		Body:   &map[string]interface{}{},
	}.TestOnError(t, &appAssert, errors.FailedToValidateID)
}

// Create membership by email lists - 201 OK
func TestCreateMembershipByEmailListWorks(t *testing.T) {
	appAssert, _, clubID := CreateSampleClub(t, nil)
	appAssert2, userID := CreateSampleUser2(t, &appAssert)

	TestRequest{
		Method: fiber.MethodPost,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/membership", clubID),
		Body: &map[string]interface{}{
			"emails": []string{
				"brennan.mic@northeastern.edu",
			},
		},
	}.TestOnStatusAndDB(t, &appAssert2,
		DBTesterWithStatus{
			Status: fiber.StatusNoContent,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var club models.Club
				var members []models.User
				err := app.Conn.Where("id = ?", clubID).First(&club).Association("Member").Find(&members).Error

				assert.NilError(err)
				assert.Equal(1, len(members))
				assert.Equal(members[0].ID, userID)
			},
		})
}

// Create membership by email lists - 400 invalid body (bad json)
func TestCreateMembershipByEmailListFailsOnInvalidBody(t *testing.T) {
	appAssert, _, clubID := CreateSampleClub(t, nil)

	badBodies := []map[string]interface{}{
		{
			"foo":   "bar",
			"alice": "bob",
		},
		{
			"x": false,
		},
	}

	for i := 0; i < len(badBodies); i += 1 {
		TestRequest{
			Method: fiber.MethodPost,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/membership", clubID),
			Body:   &badBodies[i],
		}.TestOnErrorAndDB(t, &appAssert,
			ErrorWithDBTester{
				Error: errors.FailedToParseRequestBody,
				DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
					AssertNumClubMembersRemainsAt0(app, assert, clubID)
				},
			})
	}

}

// Create membership by email lists - 404 not found (bad club id)
func TestCreateMembershipByEmailListFailsOnInvalidClubId(t *testing.T) {
	TestRequest{
		Method: fiber.MethodPost,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/membership", "gobbledygook"),
		Body: &map[string]interface{}{
			"emails": []string{
				"doe.jane@northeastern.edu",
			},
		},
	}.TestOnError(t, nil, errors.FailedToValidateID)
}

// Create membership by email lists - 404 not found (no user with email)
func TestCreateMembershipByEmailListFailsOnUserNotFound(t *testing.T) {
	appAssert, _, clubID := CreateSampleClub(t, nil)

	TestRequest{
		Method: fiber.MethodPost,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/membership", clubID),
		Body: &map[string]interface{}{
			"emails": []string{
				"doe.jane@northeastern.edu",
			},
		},
	}.TestOnErrorAndDB(t, &appAssert, ErrorWithDBTester{
		Error: errors.UserNotFound,
		DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
			AssertNumClubMembersRemainsAt0(app, assert, clubID)
		},
	})
}

// Delete membership by user id - 200 OK
func TestDeleteMembershipWorks(t *testing.T) {
	appAssert, userID, clubID := CreateSampleClubWithMembership(t, nil)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/membership/%s", clubID, userID),
		Body:   &map[string]interface{}{},
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				AssertNumClubMembersRemainsAt0(app, assert, clubID)
			},
		})
}

// Delete membership by user id - 404 not found (bad club id)
func TestDeleteMembershipFailsOnInvalidClubId(t *testing.T) {
	appAssert, userID := CreateSampleUser(t, nil)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/membership/%s", "gobbledygook", userID),
		Body:   &map[string]interface{}{},
	}.TestOnError(t, &appAssert, errors.ClubNotFound)
}

// Delete membership by user id - 404 not found (bad user id)
func TestDeleteMembershipFailsOnInvalidUserId(t *testing.T) {
	appAssert, _, clubID := CreateSampleClub(t, nil)

	TestRequest{
		Method: fiber.MethodPost,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/membership/%s", clubID, "gobbledygook"),
		Body:   &map[string]interface{}{},
	}.TestOnError(t, &appAssert, errors.UserNotFound)
}

// Delete membership by user ids - 200 OK
func TestDeleteMembershipByUserIdsWorks(t *testing.T) {
	appAssert, userID, clubID := CreateSampleClubWithMembership(t, nil)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/membership", clubID),
		Body: &map[string]interface{}{
			"ids": []string{
				userID.String(),
			},
		},
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				AssertNumClubMembersRemainsAt0(app, assert, clubID)
			},
		})
}

// Delete membership by user ids - 400 invalid body (bad json)
func TestDeleteMembershipByUserIdsFailsOnInvalidBody(t *testing.T) {
	appAssert, _, clubID := CreateSampleClub(t, nil)

	badBodies := []map[string]interface{}{
		{
			"ids": []int{
				1,
				2,
				3,
				4,
				890,
			},
		},
		{
			"foo":   "bar",
			"alice": "bob",
		},
		{
			"x": false,
		},
	}

	for i := 0; i < len(badBodies); i += 1 {
		TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/membership", clubID),
			Body:   &badBodies[i],
		}.TestOnErrorAndDB(t, &appAssert,
			ErrorWithDBTester{
				Error: errors.FailedToParseRequestBody,
				DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
					AssertNumClubMembersRemainsAt1(app, assert, clubID)
				},
			})
	}
}

// Delete membership by user ids - 404 not found (bad club id)
func TestDeleteMembershipByUserIdsFailsOnInvalidClubId(t *testing.T) {
	appAssert, userID, clubID := CreateSampleClubWithMembership(t, nil)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/membership", "gobbledygook"),
		Body: &map[string]interface{}{
			"ids": []string{
				userID.String(),
			},
		},
	}.TestOnErrorAndDB(t, &appAssert,
		ErrorWithDBTester{
			Error: errors.ClubNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				AssertNumClubMembersRemainsAt1(app, assert, clubID)
			},
		})
}

// Delete membership by user ids - 404 not found (bad user id)
func TestDeleteMembershipByUserIdsFailsOnUserNotFound(t *testing.T) {
	appAssert, _, clubID := CreateSampleClubWithMembership(t, nil)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/membership", clubID),
		Body: &map[string]interface{}{
			"ids": []string{
				"gobbledygook",
			},
		},
	}.TestOnErrorAndDB(t, &appAssert,
		ErrorWithDBTester{
			Error: errors.UserNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				AssertNumClubMembersRemainsAt1(app, assert, clubID)
			},
		})
}

// Get all club memberships for 1 user - 200 OK
func TestGetMembershipsForUserWorks(t *testing.T) {
	appAssert, userID, clubID := CreateSampleClubWithMembership(t, nil)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/users/%s/membership", userID),
		Body:   &map[string]interface{}{},
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respClubs []models.Club
				err := json.NewDecoder(resp.Body).Decode(&respClubs)
				assert.NilError(err)

				var respClub = respClubs[0]

				var dbClub models.Club
				err = app.Conn.Where("id = ?", clubID).First(&dbClub).Error
				assert.NilError(err)

				assert.Equal(dbClub.ID, respClub.ID)
				assert.Equal(dbClub.Name, respClub.Name)
				assert.Equal(dbClub.Preview, respClub.Preview)
				assert.Equal(dbClub.Description, respClub.Description)
				assert.Equal(dbClub.NumMembers, respClub.NumMembers)
				assert.Equal(dbClub.IsRecruiting, respClub.IsRecruiting)
				assert.Equal(dbClub.RecruitmentCycle, respClub.RecruitmentCycle)
				assert.Equal(dbClub.RecruitmentType, respClub.RecruitmentType)
				assert.Equal(dbClub.ApplicationLink, respClub.ApplicationLink)
				assert.Equal(dbClub.Logo, respClub.Logo)
			},
		})
}

// Get all club memberships for 1 user - 404 not found (bad user id)
func TestGetMembershipsForUserFailsOnInvalidUserId(t *testing.T) {
	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/users/%s/membership", "gobbledygook"),
		Body:   &map[string]interface{}{},
	}.TestOnError(t, nil, errors.UserNotFound)
}

// Get all club memberships for 1 club - 200 OK
func TestGetMembershipsForClubWorks(t *testing.T) {
	appAssert, userID, clubID := CreateSampleClubWithMembership(t, nil)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/membership", clubID),
		Body:   &map[string]interface{}{},
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respUsers []models.User
				err := json.NewDecoder(resp.Body).Decode(&respUsers)
				assert.NilError(err)

				respUser := respUsers[0]

				var dbUser models.User
				err = app.Conn.Where("id = ?", userID).First(&dbUser).Error
				assert.NilError(err)

				assert.Equal(dbUser.FirstName, respUser.FirstName)
				assert.Equal(dbUser.LastName, respUser.LastName)
				assert.Equal(dbUser.Email, respUser.Email)
				assert.Equal(dbUser.NUID, respUser.NUID)
				assert.Equal(dbUser.College, respUser.College)
				assert.Equal(dbUser.Year, respUser.Year)
			},
		})
}

// Get all club memberships for 1 club - 404 not found (bad club id)
func TestGetMembershipsForClubFailsOnInvalidClubId(t *testing.T) {
	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/membership", "gobbledygook"),
		Body:   &map[string]interface{}{},
	}.TestOnError(t, nil, errors.FailedToValidateID)
}
