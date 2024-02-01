package tests

import (
	"fmt"

	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/huandu/go-assert"
)

func AssertClubTagsRespDB(app TestApp, assert *assert.A, resp *http.Response, id uuid.UUID) {
	var respTags []models.Tag

	// Retrieve the tags from the response:
	err := json.NewDecoder(resp.Body).Decode(&respTags)

	assert.NilError(err)

	// Retrieve the club connected to the tags:
	var dbClub models.Club
	err = app.Conn.First(&dbClub, id).Error

	assert.NilError(err)

	// Retrieve the tags in the bridge table associated with the club:
	var dbTags []models.Tag
	err = app.Conn.Model(&dbClub).Association("Tag").Find(&dbTags)

	assert.NilError(err)

	// Confirm all the resp tags are equal to the db tags:
	for i, respTag := range respTags {
		assert.Equal(respTag.ID, dbTags[i].ID)
		assert.Equal(respTag.Name, dbTags[i].Name)
		assert.Equal(respTag.CategoryID, dbTags[i].CategoryID)
	}
}

func TestCreateClubTagsWorks(t *testing.T) {
	appAssert, _, uuid := CreateSampleClub(t, nil)

	// Create a set of tags:
	tagUUIDs := CreateSetOfTags(t, appAssert)

	// Confirm adding real tags adds them to the club:
	TestRequest{
		Method: fiber.MethodPost,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", uuid),
		Body:   SampleTagIDsFactory(&tagUUIDs),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusCreated,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				AssertClubTagsRespDB(app, assert, resp, uuid)
			},
		},
	)

	appAssert.Close()
}

// func TestCreateClubTagsFailsOnInvalidDataType(t *testing.T) {
// 	_, _, uuid := CreateSampleClub(t, nil)
	
// 	// Invalid tag data types:
// 	invalidTags := []interface{}{
// 		[]string{"1", "2", "34"},
// 		[]models.Tag{{Name: "Test", CategoryID: uuid.UUID{}}, {Name: "Test2", CategoryID: uuid.UUID{}}},
// 		[]float32{1.32, 23.5, 35.1},
// 	}

// 	// Test each of the invalid tags:
// 	for _, tag := range invalidTags {
// 		malformedTag := *SampleTagIDsFactory(nil)
// 		malformedTag["tags"] = tag

// 		TestRequest{
// 			Method: fiber.MethodPost,
// 			Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", uuid),
// 			Body:   &malformedTag,
// 		}.TestOnError(t, nil, errors.FailedToParseRequestBody).Close()
// 	}
// }

func TestCreateClubTagsFailsOnInvalidClubID(t *testing.T) {
	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		TestRequest{
			Method: fiber.MethodPost,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/tags", badRequest),
			Body:   SampleTagIDsFactory(nil),
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}


func TestCreateClubTagsFailsOnInvalidKey(t *testing.T) {
	appAssert, _, uuid := CreateSampleClub(t, nil)

	invalidBody := []map[string]interface{}{
		{
			"tag": UUIDSlice{testUUID, testUUID},
		},
		{
			"tagIDs": []uint{1, 2, 3},
		},
	}

	for _, body := range invalidBody {
		TestRequest{
			Method: fiber.MethodPost,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", uuid),
			Body:   &body,
		}.TestOnError(t, &appAssert, errors.FailedToValidateClubTags)
	}

	appAssert.Close()
}

func TestCreateClubTagsNoneAddedIfInvalid(t *testing.T) {
	appAssert, _, uuid := CreateSampleClub(t, nil)
	
	TestRequest{
		Method: fiber.MethodPost,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", uuid),
		Body:   SampleTagIDsFactory(nil),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusCreated,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respTags []models.Tag

				err := json.NewDecoder(resp.Body).Decode(&respTags)

				assert.NilError(err)

				assert.Equal(len(respTags), 0)
			},
		},
	)

	appAssert.Close()
}

func TestGetClubTagsFailsOnNonExistentClub(t *testing.T) {
	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", uuid.New()),
	}.TestOnError(t, nil, errors.ClubNotFound).Close()
}

func TestGetClubTagsReturnsEmptyListWhenNoneAdded(t *testing.T) {
	appAssert, _, uuid := CreateSampleClub(t, nil)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", uuid),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: 200,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respTags []models.Tag

				err := json.NewDecoder(resp.Body).Decode(&respTags)

				assert.NilError(err)

				assert.Equal(len(respTags), 0)
			},
		},
	)

	appAssert.Close()
}

func TestGetClubTagsReturnsCorrectList(t *testing.T) {
	appAssert, _, uuid := CreateSampleClub(t, nil)

	// Create a set of tags:
	tagUUIDs := CreateSetOfTags(t, appAssert)

	// Add the tags:
	TestRequest{
		Method: fiber.MethodPost,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", uuid),
		Body:   SampleTagIDsFactory(&tagUUIDs),
	}.TestOnStatus(t, &appAssert, fiber.StatusCreated)

	// Get the tags:
	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", uuid),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				AssertClubTagsRespDB(app, assert, resp, uuid)
			},
		},
	)

	appAssert.Close()
}
