package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
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

func AssertCategoryBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) uuid.UUID {
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

func AssertCategoryWithBodyRespDBMostRecent(app TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) uuid.UUID {
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

func AssertSampleCategoryBodyRespDB(app TestApp, assert *assert.A, resp *http.Response) uuid.UUID {
	return AssertCategoryBodyRespDB(app, assert, resp, SampleCategoryFactory())
}

func CreateSampleCategory(t *testing.T, existingAppAssert *ExistingAppAssert) (ExistingAppAssert, uuid.UUID) {
	var sampleCategoryUUID uuid.UUID

	newAppAssert := TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/categories/",
		Body:   SampleCategoryFactory(),
	}.TestOnStatusAndDB(t, existingAppAssert,
		DBTesterWithStatus{
			Status: fiber.StatusCreated,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
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
	TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/categories/",
		Body: &map[string]interface{}{
			"id":   12,
			"name": "Foo",
		},
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status: fiber.StatusCreated,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				AssertSampleCategoryBodyRespDB(app, assert, resp)
			},
		},
	).Close()
}

func Assert1Category(app TestApp, assert *assert.A, resp *http.Response) {
	AssertNumCategoriesRemainsAtN(app, assert, resp, 1)
}

func AssertNoCategories(app TestApp, assert *assert.A, resp *http.Response) {
	AssertNumCategoriesRemainsAtN(app, assert, resp, 0)
}

func AssertNumCategoriesRemainsAtN(app TestApp, assert *assert.A, resp *http.Response, n int) {
	var categories []models.Category

	err := app.Conn.Find(&categories).Error

	assert.NilError(err)

	assert.Equal(n, len(categories))
}

func TestCreateCategoryFailsIfNameIsNotString(t *testing.T) {
	TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/categories/",
		Body: &map[string]interface{}{
			"name": 1231,
		},
	}.TestOnErrorAndDB(t, nil,
		ErrorWithDBTester{
			Error:    errors.FailedToParseRequestBody,
			DBTester: AssertNoCategories,
		},
	).Close()
}

func TestCreateCategoryFailsIfNameIsMissing(t *testing.T) {
	TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/categories/",
		Body:   &map[string]interface{}{},
	}.TestOnErrorAndDB(t, nil,
		ErrorWithDBTester{
			Error:    errors.FailedToValidateCategory,
			DBTester: AssertNoCategories,
		},
	).Close()
}

func TestCreateCategoryFailsIfCategoryWithThatNameAlreadyExists(t *testing.T) {
	existingAppAssert, _ := CreateSampleCategory(t, nil)

	var TestNumCategoriesRemainsAt1 = func(app TestApp, assert *assert.A, resp *http.Response) {
		AssertNumCategoriesRemainsAtN(app, assert, resp, 1)
	}

	for _, permutation := range AllCasingPermutations((*SampleCategoryFactory())["name"].(string)) {
		modifiedSampleCategoryBody := *SampleCategoryFactory()
		modifiedSampleCategoryBody["name"] = permutation

		TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/categories/",
			Body:   &modifiedSampleCategoryBody,
		}.TestOnErrorAndDB(t, &existingAppAssert,
			ErrorWithDBTester{
				Error:    errors.CategoryAlreadyExists,
				DBTester: TestNumCategoriesRemainsAt1,
			},
		)
	}

	existingAppAssert.Close()
}

func TestGetCategoryWorks(t *testing.T) {
	existingAppAssert, uuid := CreateSampleCategory(t, nil)

	TestRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/api/v1/categories/%s", uuid),
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
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
		TestRequest{
			Method: "GET",
			Path:   fmt.Sprintf("/api/v1/categories/%s", badRequest),
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

func TestGetCategoryFailsNotFound(t *testing.T) {
	TestRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/api/v1/categories/%s", uuid.New()),
	}.TestOnError(t, nil, errors.CategoryNotFound).Close()
}

func TestGetCategoriesWorks(t *testing.T) {
	existingAppAssert, _ := CreateSampleCategory(t, nil)

	TestRequest{
		Method: "GET",
		Path:   "/api/v1/categories/",
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
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

func AssertUpdatedCategoryBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) {
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

	var AssertUpdatedCategoryBodyRespDB = func(app TestApp, assert *assert.A, resp *http.Response) {
		AssertUpdatedCategoryBodyRespDB(app, assert, resp, &category)
	}

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   fmt.Sprintf("/api/v1/categories/%s", uuid),
		Body:   &category,
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusOK,
			DBTester: AssertUpdatedCategoryBodyRespDB,
		},
	).Close()
}

func TestUpdateCategoryWorksWithSameDetails(t *testing.T) {
	existingAppAssert, uuid := CreateSampleCategory(t, nil)

	category := *SampleCategoryFactory()
	category["id"] = uuid

	var AssertSampleCategoryBodyRespDB = func(app TestApp, assert *assert.A, resp *http.Response) {
		AssertUpdatedCategoryBodyRespDB(app, assert, resp, &category)
	}

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   fmt.Sprintf("/api/v1/categories/%s", uuid),
		Body:   &category,
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusOK,
			DBTester: AssertSampleCategoryBodyRespDB,
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
		TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/categories/%s", badRequest),
			Body:   SampleCategoryFactory(),
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

func TestDeleteCategoryWorks(t *testing.T) {
	existingAppAssert, uuid := CreateSampleCategory(t, nil)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/categories/%s", uuid),
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusNoContent,
			DBTester: AssertNoCategories,
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
		TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/categories/%s", badRequest),
		}.TestOnErrorAndDB(t, &existingAppAssert,
			ErrorWithDBTester{
				Error:    errors.FailedToValidateID,
				DBTester: Assert1Category,
			},
		)
	}

	existingAppAssert.Close()
}

func TestDeleteCategoryFailsNotFound(t *testing.T) {
	existingAppAssert, _ := CreateSampleCategory(t, nil)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/categories/%s", uuid.New()),
	}.TestOnErrorAndDB(t, &existingAppAssert,
		ErrorWithDBTester{
			Error:    errors.CategoryNotFound,
			DBTester: Assert1Category,
		},
	).Close()
}
