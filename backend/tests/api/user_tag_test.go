package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
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

func CreateSetOfTags(t *testing.T, appAssert *h.ExistingAppAssert) ([]uuid.UUID, *h.ExistingAppAssert) {
	if appAssert == nil {
		app, assert := h.InitTest(t)
		appAssert = &h.ExistingAppAssert{App: app, Assert: assert}
	}

	categories := SampleCategoriesFactory()

	categoryIDs := []uuid.UUID{}
	for _, category := range *categories {
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/categories/",
			Body:   &category,
			Role:   &models.Super,
		}.TestOnStatusAndDB(t, appAssert,
			h.TesterWithStatus{
				Status: fiber.StatusCreated,
				Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
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
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/tags/",
			Body:   &tag,
			Role:   &models.Super,
		}.TestOnStatusAndDB(t, appAssert,
			h.TesterWithStatus{
				Status: fiber.StatusCreated,
				Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
					var respTag models.Tag

					err := json.NewDecoder(resp.Body).Decode(&respTag)

					assert.NilError(err)

					tagIDs = append(tagIDs, respTag.ID)
				},
			},
		)
	}

	return tagIDs, appAssert
}

func AssertUserTagsRespDB(app h.TestApp, assert *assert.A, resp *http.Response, id uuid.UUID) {
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

func AssertSampleUserTagsRespDB(app h.TestApp, assert *assert.A, resp *http.Response, uuid uuid.UUID) {
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

		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               "/api/v1/users/:userID/tags/",
			Body:               &malformedTag,
			Role:               &models.Student,
			TestUserIDRequired: h.BoolToPointer(true),
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
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   fmt.Sprintf("/api/v1/users/%s/tags", badRequest),
			Body:   SampleTagIDsFactory(nil),
			Role:   &models.Student,
		}.TestOnError(t, nil, errors.FailedToParseUUID).Close()
	}
}

type UUIDSlice []uuid.UUID

var testUUID = uuid.New()

func TestCreateUserTagsFailsOnInvalidKey(t *testing.T) {
	invalidBody := []map[string]interface{}{
		{
			"tag": UUIDSlice{testUUID, testUUID},
		},
		{
			"tagIDs": []uint{1, 2, 3},
		},
	}

	for _, body := range invalidBody {
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               "/api/v1/users/:userID/tags/",
			Body:               &body,
			Role:               &models.Student,
			TestUserIDRequired: h.BoolToPointer(true),
		}.TestOnError(t, nil, errors.FailedToValidateUserTags).Close()
	}
}

// TODO: should this be unauthorized or not found?
func TestCreateUserTagsFailsOnNonExistentUser(t *testing.T) {
	uuid := uuid.New()

	h.TestRequest{
		Method: fiber.MethodPost,
		Path:   fmt.Sprintf("/api/v1/users/%s/tags/", uuid),
		Body:   SampleTagIDsFactory(nil),
		Role:   &models.Super,
	}.TestOnErrorAndDB(t, nil,
		h.ErrorWithTester{
			Error: errors.UserNotFound,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				var dbUser models.User
				err := app.Conn.First(&dbUser, uuid).Error

				assert.Error(err)
			},
		},
	)
}

func TestCreateUserTagsWorks(t *testing.T) {
	// Create a set of tags:
	tagUUIDs, appAssert := CreateSetOfTags(t, nil)

	// Confirm adding real tags adds them to the user:
	h.TestRequest{
		Method:             fiber.MethodPost,
		Path:               "/api/v1/users/:userID/tags/",
		Body:               SampleTagIDsFactory(&tagUUIDs),
		Role:               &models.Super,
		TestUserIDRequired: h.BoolToPointer(true),
	}.TestOnStatusAndDB(t, appAssert,
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				AssertSampleUserTagsRespDB(app, assert, resp, app.TestUser.UUID)
			},
		},
	)

	appAssert.Close()
}

func TestCreateUserTagsNoneAddedIfInvalid(t *testing.T) {
	h.TestRequest{
		Method:             fiber.MethodPost,
		Path:               "/api/v1/users/:userID/tags/",
		Body:               SampleTagIDsFactory(nil),
		Role:               &models.Super,
		TestUserIDRequired: h.BoolToPointer(true),
	}.TestOnStatusAndDB(t, nil,
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				var respTags []models.Tag

				err := json.NewDecoder(resp.Body).Decode(&respTags)

				assert.NilError(err)

				assert.Equal(len(respTags), 0)
			},
		},
	).Close()
}

// TODO: should this be unauthorized or not found?
func TestGetUserTagsFailsOnNonExistentUser(t *testing.T) {
	h.TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/users/%s/tags/", uuid.New()),
		Role:   &models.Super,
	}.TestOnError(t, nil, errors.UserNotFound).Close()
}

func TestGetUserTagsReturnsEmptyListWhenNoneAdded(t *testing.T) {
	h.TestRequest{
		Method:             fiber.MethodGet,
		Path:               "/api/v1/users/:userID/tags/",
		Role:               &models.Student,
		TestUserIDRequired: h.BoolToPointer(true),
	}.TestOnStatusAndDB(t, nil,
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				var respTags []models.Tag

				err := json.NewDecoder(resp.Body).Decode(&respTags)

				assert.NilError(err)

				assert.Equal(len(respTags), 0)
			},
		},
	).Close()
}

func TestGetUserTagsReturnsCorrectList(t *testing.T) {
	// Create a set of tags:
	tagUUIDs, appAssert := CreateSetOfTags(t, nil)

	// Add the tags:
	h.TestRequest{
		Method:             fiber.MethodPost,
		Path:               "/api/v1/users/:userID/tags/",
		Body:               SampleTagIDsFactory(&tagUUIDs),
		Role:               &models.Student,
		TestUserIDRequired: h.BoolToPointer(true),
	}.TestOnStatus(t, appAssert, fiber.StatusCreated)

	// Get the tags:
	h.TestRequest{
		Method:             fiber.MethodGet,
		Path:               "/api/v1/users/:userID/tags/",
		Role:               &models.Student,
		TestUserIDRequired: h.BoolToPointer(true),
	}.TestOnStatusAndDB(t, appAssert,
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				AssertSampleUserTagsRespDB(app, assert, resp, app.TestUser.UUID)
			},
		},
	)

	appAssert.Close()
}
