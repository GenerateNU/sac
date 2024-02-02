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

func SampleCategoriesFactory() *[]map[string]interface{} {
	return &[]map[string]interface{}{
		{
			"name": "Business",
		},
		{
			"name": "STEM",
		},
	}
}

func SampleTagsFactory(categoryIDs []uuid.UUID) *[]map[string]interface{} {
	lenOfIDs := len(categoryIDs)

	return &[]map[string]interface{}{
		{
			"name":        "Computer Science",
			"category_id": categoryIDs[1%lenOfIDs],
		},
		{
			"name":        "Mechanical Engineering",
			"category_id": categoryIDs[1%lenOfIDs],
		},
		{
			"name":        "Finance",
			"category_id": categoryIDs[0%lenOfIDs],
		},
	}
}

func SampleTagIDsFactory(tagIDs *[]uuid.UUID) *map[string]interface{} {
	tags := tagIDs

	if tags == nil {
		tags = &[]uuid.UUID{uuid.New()}
	}

	return &map[string]interface{}{
		"tags": tags,
	}
}

func CreateSetOfTags(t *testing.T, appAssert ExistingAppAssert) []uuid.UUID {
	categories := SampleCategoriesFactory()

	categoryIDs := []uuid.UUID{}
	for _, category := range *categories {
		TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/categories/",
			Body:   &category,
		}.TestOnStatusAndDB(t, &appAssert,
			DBTesterWithStatus{
				Status: fiber.StatusCreated,
				DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
					var respCategory models.Category

					err := json.NewDecoder(resp.Body).Decode(&respCategory)

					assert.NilError(err)

					categoryIDs = append(categoryIDs, respCategory.ID)
				},
			},
		)
	}

	tags := SampleTagsFactory(categoryIDs)

	tagIDs := []uuid.UUID{}
	for _, tag := range *tags {
		TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/tags/",
			Body:   &tag,
		}.TestOnStatusAndDB(t, &appAssert,
			DBTesterWithStatus{
				Status: fiber.StatusCreated,
				DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
					var respTag models.Tag

					err := json.NewDecoder(resp.Body).Decode(&respTag)

					assert.NilError(err)

					tagIDs = append(tagIDs, respTag.ID)
				},
			},
		)
	}

	return tagIDs
}

func AssertUserTagsRespDB(app TestApp, assert *assert.A, resp *http.Response, id uuid.UUID) {
	var respTags []models.Tag

	// Retrieve the tags from the response:
	err := json.NewDecoder(resp.Body).Decode(&respTags)

	assert.NilError(err)

	// Retrieve the user connected to the tags:
	var dbUser models.User
	err = app.Conn.First(&dbUser, id).Error

	assert.NilError(err)

	// Retrieve the tags in the bridge table associated with the user:
	var dbTags []models.Tag
	err = app.Conn.Model(&dbUser).Association("Tag").Find(&dbTags)

	assert.NilError(err)

	// Confirm all the resp tags are equal to the db tags:
	for i, respTag := range respTags {
		assert.Equal(respTag.ID, dbTags[i].ID)
		assert.Equal(respTag.Name, dbTags[i].Name)
		assert.Equal(respTag.CategoryID, dbTags[i].CategoryID)
	}
}

func AssertSampleUserTagsRespDB(app TestApp, assert *assert.A, resp *http.Response, uuid uuid.UUID) {
	AssertUserTagsRespDB(app, assert, resp, uuid)
}

func TestCreateUserTagsFailsOnInvalidDataType(t *testing.T) {
	// Invalid tag data types:
	invalidTags := []interface{}{
		[]string{"1", "2", "34"},
		[]models.Tag{{Name: "Test", CategoryID: uuid.UUID{}}, {Name: "Test2", CategoryID: uuid.UUID{}}},
		[]float32{1.32, 23.5, 35.1},
	}

	// Test each of the invalid tags:
	for _, tag := range invalidTags {
		malformedTag := *SampleTagIDsFactory(nil)
		malformedTag["tags"] = tag

		TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/users/1/tags/",
			Body:   &malformedTag,
		}.TestOnError(t, nil, errors.FailedToParseRequestBody).Close()
	}
}

func TestCreateUserTagsFailsOnInvalidUserID(t *testing.T) {
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
			Path:   fmt.Sprintf("/api/v1/users/%s/tags", badRequest),
			Body:   SampleTagIDsFactory(nil),
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

type UUIDSlice []uuid.UUID

var testUUID = uuid.New()

func TestCreateUserTagsFailsOnInvalidKey(t *testing.T) {
	appAssert, uuid := CreateSampleUser(t, nil)

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
			Path:   fmt.Sprintf("/api/v1/users/%s/tags/", uuid),
			Body:   &body,
		}.TestOnError(t, &appAssert, errors.FailedToValidateUserTags)
	}

	appAssert.Close()
}

func TestCreateUserTagsFailsOnNonExistentUser(t *testing.T) {
	TestRequest{
		Method: fiber.MethodPost,
		Path:   fmt.Sprintf("/api/v1/users/%s/tags", uuid.New()),
		Body:   SampleTagIDsFactory(nil),
	}.TestOnError(t, nil, errors.UserNotFound).Close()
}

func TestCreateUserTagsWorks(t *testing.T) {
	appAssert, uuid := CreateSampleUser(t, nil)

	// Create a set of tags:
	tagUUIDs := CreateSetOfTags(t, appAssert)

	// Confirm adding real tags adds them to the user:
	TestRequest{
		Method: fiber.MethodPost,
		Path:   fmt.Sprintf("/api/v1/users/%s/tags/", uuid),
		Body:   SampleTagIDsFactory(&tagUUIDs),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusCreated,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				AssertSampleUserTagsRespDB(app, assert, resp, uuid)
			},
		},
	)

	appAssert.Close()
}

func TestCreateUserTagsNoneAddedIfInvalid(t *testing.T) {
	appAssert, uuid := CreateSampleUser(t, nil)

	TestRequest{
		Method: fiber.MethodPost,
		Path:   fmt.Sprintf("/api/v1/users/%s/tags/", uuid),
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

func TestGetUserTagsFailsOnNonExistentUser(t *testing.T) {
	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/users/%s/tags/", uuid.New()),
	}.TestOnError(t, nil, errors.UserNotFound).Close()
}

func TestGetUserTagsReturnsEmptyListWhenNoneAdded(t *testing.T) {
	appAssert, uuid := CreateSampleUser(t, nil)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/users/%s/tags/", uuid),
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

func TestGetUserTagsReturnsCorrectList(t *testing.T) {
	appAssert, uuid := CreateSampleUser(t, nil)

	// Create a set of tags:
	tagUUIDs := CreateSetOfTags(t, appAssert)

	// Add the tags:
	TestRequest{
		Method: fiber.MethodPost,
		Path:   fmt.Sprintf("/api/v1/users/%s/tags/", uuid),
		Body:   SampleTagIDsFactory(&tagUUIDs),
	}.TestOnStatus(t, &appAssert, fiber.StatusCreated)

	// Get the tags:
	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/users/%s/tags/", uuid),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				AssertSampleUserTagsRespDB(app, assert, resp, uuid)
			},
		},
	)

	appAssert.Close()
}
