package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/huandu/go-assert"

	"github.com/goccy/go-json"
)

func SampleCategoryFactory() *map[string]interface{} {
	return &map[string]interface{}{
		"name": "Foo",
	}
}

func AssertCategoryBodyRespDB(app h.TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respCategory models.Category

	err := json.NewDecoder(resp.Body).Decode(&respCategory)

	assert.NilError(err)

	var dbCategories []models.Category

	err = app.Conn.Find(&dbCategories).Error

	assert.NilError(err)

	assert.Equal(1, len(dbCategories))

	dbCategory := dbCategories[0]

	assert.Equal(dbCategory.ID, respCategory.ID)
	assert.Equal(dbCategory.Name, respCategory.Name)

	assert.Equal((*body)["name"].(string), dbCategory.Name)

	return dbCategory.ID
}

func AssertCategoryWithBodyRespDBMostRecent(app h.TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respCategory models.Category

	err := json.NewDecoder(resp.Body).Decode(&respCategory)

	assert.NilError(err)

	var dbCategory models.Category

	err = app.Conn.Order("created_at desc").First(&dbCategory).Error

	assert.NilError(err)

	assert.Equal(dbCategory.ID, respCategory.ID)
	assert.Equal(dbCategory.Name, respCategory.Name)

	assert.Equal((*body)["name"].(string), dbCategory.Name)

	return dbCategory.ID
}

func AssertSampleCategoryBodyRespDB(app h.TestApp, assert *assert.A, resp *http.Response) uuid.UUID {
	return AssertCategoryBodyRespDB(app, assert, resp, SampleCategoryFactory())
}

func CreateSampleCategory(t *testing.T, existingAppAssert *h.ExistingAppAssert) (h.ExistingAppAssert, uuid.UUID) {
	var sampleCategoryUUID uuid.UUID

	newAppAssert := h.TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/categories/",
		Body:   SampleCategoryFactory(),
		Role:   &models.Super,
	}.TestOnStatusAndDB(t, existingAppAssert,
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				sampleCategoryUUID = AssertSampleCategoryBodyRespDB(app, assert, resp)
			},
		},
	)

	if existingAppAssert == nil {
		return newAppAssert, sampleCategoryUUID
	} else {
		return *existingAppAssert, sampleCategoryUUID
	}
}

func TestCreateCategoryWorks(t *testing.T) {
	existingAppAssert, _ := CreateSampleCategory(t, nil)
	existingAppAssert.Close()
}

func TestCreateCategoryIgnoresid(t *testing.T) {
	h.TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/categories/",
		Body: &map[string]interface{}{
			"id":   12,
			"name": "Foo",
		},
		Role: &models.Super,
	}.TestOnStatusAndDB(t, nil,
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				AssertSampleCategoryBodyRespDB(app, assert, resp)
			},
		},
	).Close()
}

func Assert1Category(app h.TestApp, assert *assert.A, resp *http.Response) {
	AssertNumCategoriesRemainsAtN(app, assert, resp, 1)
}

func AssertNoCategories(app h.TestApp, assert *assert.A, resp *http.Response) {
	AssertNumCategoriesRemainsAtN(app, assert, resp, 0)
}

func AssertNumCategoriesRemainsAtN(app h.TestApp, assert *assert.A, resp *http.Response, n int) {
	var categories []models.Category

	err := app.Conn.Find(&categories).Error

	assert.NilError(err)

	assert.Equal(n, len(categories))
}

func TestCreateCategoryFailsIfNameIsNotString(t *testing.T) {
	h.TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/categories/",
		Body: &map[string]interface{}{
			"name": 1231,
		},
		Role: &models.Super,
	}.TestOnErrorAndDB(t, nil,
		h.ErrorWithTester{
			Error:  errors.FailedToParseRequestBody,
			Tester: AssertNoCategories,
		},
	).Close()
}

func TestCreateCategoryFailsIfNameIsMissing(t *testing.T) {
	h.TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/categories/",
		Body:   &map[string]interface{}{},
		Role:   &models.Super,
	}.TestOnErrorAndDB(t, nil,
		h.ErrorWithTester{
			Error:  errors.FailedToValidateCategory,
			Tester: AssertNoCategories,
		},
	).Close()
}

func TestCreateCategoryFailsIfCategoryWithThatNameAlreadyExists(t *testing.T) {
	existingAppAssert, _ := CreateSampleCategory(t, nil)

	TestNumCategoriesRemainsAt1 := func(app h.TestApp, assert *assert.A, resp *http.Response) {
		AssertNumCategoriesRemainsAtN(app, assert, resp, 1)
	}

	for _, permutation := range h.AllCasingPermutations((*SampleCategoryFactory())["name"].(string)) {
		modifiedSampleCategoryBody := *SampleCategoryFactory()
		modifiedSampleCategoryBody["name"] = permutation

		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/categories/",
			Body:   &modifiedSampleCategoryBody,
			Role:   &models.Super,
		}.TestOnErrorAndDB(t, &existingAppAssert,
			h.ErrorWithTester{
				Error:  errors.CategoryAlreadyExists,
				Tester: TestNumCategoriesRemainsAt1,
			},
		)
	}

	existingAppAssert.Close()
}

func TestGetCategoryWorks(t *testing.T) {
	existingAppAssert, uuid := CreateSampleCategory(t, nil)

	h.TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/categories/%s", uuid),
		Role:   &models.Super,
	}.TestOnStatusAndDB(t, &existingAppAssert,
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				AssertSampleCategoryBodyRespDB(app, assert, resp)
			},
		},
	).Close()
}

func TestGetCategoryFailsBadRequest(t *testing.T) {
	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/categories/%s", badRequest),
			Role:   &models.Super,
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

func TestGetCategoryFailsNotFound(t *testing.T) {
	h.TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/categories/%s", uuid.New()),
		Role:   &models.Super,
	}.TestOnError(t, nil, errors.CategoryNotFound).Close()
}

func TestGetCategoriesWorks(t *testing.T) {
	existingAppAssert, _ := CreateSampleCategory(t, nil)

	h.TestRequest{
		Method: fiber.MethodGet,
		Path:   "/api/v1/categories/",
		Role:   &models.Super,
	}.TestOnStatusAndDB(t, &existingAppAssert,
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				var categories []models.Category

				err := app.Conn.Find(&categories).Error

				assert.NilError(err)

				var respCategories []models.Category

				err = json.NewDecoder(resp.Body).Decode(&respCategories)

				assert.NilError(err)

				assert.Equal(1, len(respCategories))
				assert.Equal(1, len(categories))

				assert.Equal(categories[0].ID, respCategories[0].ID)

				assert.Equal(categories[0].Name, respCategories[0].Name)

				assert.Equal((*SampleCategoryFactory())["name"].(string), categories[0].Name)
			},
		},
	).Close()
}

func AssertUpdatedCategoryBodyRespDB(app h.TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) {
	var respCategory models.Category

	err := json.NewDecoder(resp.Body).Decode(&respCategory)

	assert.NilError(err)

	var dbCategories []models.Category

	err = app.Conn.Find(&dbCategories).Error

	assert.NilError(err)

	assert.Equal(1, len(dbCategories))

	dbCategory := dbCategories[0]

	assert.Equal(dbCategory.ID, respCategory.ID)
	assert.Equal(dbCategory.Name, respCategory.Name)

	assert.Equal((*body)["id"].(uuid.UUID), dbCategory.ID)
	assert.Equal((*body)["name"].(string), dbCategory.Name)
}

func TestUpdateCategoryWorks(t *testing.T) {
	existingAppAssert, uuid := CreateSampleCategory(t, nil)

	category := map[string]interface{}{
		"id":   uuid,
		"name": "Arts & Crafts",
	}

	AssertUpdatedCategoryBodyRespDB := func(app h.TestApp, assert *assert.A, resp *http.Response) {
		AssertUpdatedCategoryBodyRespDB(app, assert, resp, &category)
	}

	h.TestRequest{
		Method: fiber.MethodPatch,
		Path:   fmt.Sprintf("/api/v1/categories/%s", uuid),
		Body:   &category,
		Role:   &models.Super,
	}.TestOnStatusAndDB(t, &existingAppAssert,
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: AssertUpdatedCategoryBodyRespDB,
		},
	).Close()
}

func TestUpdateCategoryWorksWithSameDetails(t *testing.T) {
	existingAppAssert, uuid := CreateSampleCategory(t, nil)

	category := *SampleCategoryFactory()
	category["id"] = uuid

	AssertSampleCategoryBodyRespDB := func(app h.TestApp, assert *assert.A, resp *http.Response) {
		AssertUpdatedCategoryBodyRespDB(app, assert, resp, &category)
	}

	h.TestRequest{
		Method: fiber.MethodPatch,
		Path:   fmt.Sprintf("/api/v1/categories/%s", uuid),
		Body:   &category,
		Role:   &models.Super,
	}.TestOnStatusAndDB(t, &existingAppAssert,
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: AssertSampleCategoryBodyRespDB,
		},
	).Close()
}

func TestUpdateCategoryFailsBadRequest(t *testing.T) {
	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/categories/%s", badRequest),
			Body:   SampleCategoryFactory(),
			Role:   &models.Super,
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

func TestDeleteCategoryWorks(t *testing.T) {
	existingAppAssert, uuid := CreateSampleCategory(t, nil)

	h.TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/categories/%s", uuid),
		Role:   &models.Super,
	}.TestOnStatusAndDB(t, &existingAppAssert,
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: AssertNoCategories,
		},
	).Close()
}

func TestDeleteCategoryFailsBadRequest(t *testing.T) {
	existingAppAssert, _ := CreateSampleCategory(t, nil)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/categories/%s", badRequest),
			Role:   &models.Super,
		}.TestOnErrorAndDB(t, &existingAppAssert,
			h.ErrorWithTester{
				Error:  errors.FailedToValidateID,
				Tester: Assert1Category,
			},
		)
	}

	existingAppAssert.Close()
}

func TestDeleteCategoryFailsNotFound(t *testing.T) {
	existingAppAssert, _ := CreateSampleCategory(t, nil)

	h.TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/categories/%s", uuid.New()),
		Role:   &models.Super,
	}.TestOnErrorAndDB(t, &existingAppAssert,
		h.ErrorWithTester{
			Error:  errors.CategoryNotFound,
			Tester: Assert1Category,
		},
	).Close()
}
