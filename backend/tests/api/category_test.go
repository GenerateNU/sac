package tests

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/huandu/go-assert"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/goccy/go-json"
)

func SampleCategoryFactory() *map[string]interface{} {
	return &map[string]interface{}{
		"name": "Foo",
	}
}

func AssertCategoryWithIDBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, id uint, body *map[string]interface{}) {
	var respCategory models.Category

	err := json.NewDecoder(resp.Body).Decode(&respCategory)

	assert.NilError(err)

	var dbCategory models.Category

	err = app.Conn.First(&dbCategory, id).Error

	assert.NilError(err)
	x, _ := io.ReadAll(resp.Body)
	fmt.Println(string(x))
	assert.Equal(dbCategory.ID, respCategory.ID)
	assert.Equal(dbCategory.Name, respCategory.Name)

	assert.Equal((*body)["name"].(string), dbCategory.Name)

}

func AssertSampleCategoryBodyRespDB(app TestApp, assert *assert.A, resp *http.Response) {
	AssertCategoryWithIDBodyRespDB(app, assert, resp, 1, SampleCategoryFactory())
}

func CreateSampleCategory(t *testing.T) ExistingAppAssert {
	return TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/categories/",
		Body:   SampleCategoryFactory(),
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status:   fiber.StatusCreated,
			DBTester: AssertSampleCategoryBodyRespDB,
		},
	)
}

func TestCreateCategoryWorks(t *testing.T) {
	CreateSampleCategory(t).Close()
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
			Status:   fiber.StatusCreated,
			DBTester: AssertSampleCategoryBodyRespDB,
		},
	).Close()
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
	}.TestOnStatusMessageAndDB(t, nil,
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
	}.TestOnStatusMessageAndDB(t, nil,
		ErrorWithDBTester{
			Error:    errors.FailedToValidateCategory,
			DBTester: AssertNoCategories,
		},
	).Close()
}

func TestCreateCategoryFailsIfCategoryWithThatNameAlreadyExists(t *testing.T) {

	existingAppAssert := CreateSampleCategory(t)

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
		}.TestOnStatusMessageAndDB(t, &existingAppAssert,
			ErrorWithDBTester{
				Error:    errors.CategoryAlreadyExists,
				DBTester: TestNumCategoriesRemainsAt1,
			},
		)
	}

	existingAppAssert.Close()
}

func TestGetCategoryWorks(t *testing.T) {
	existingAppAssert := CreateSampleCategory(t)

	TestRequest{
		Method: "GET",
		Path:   "/api/v1/categories/1",
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusOK,
			DBTester: AssertSampleCategoryBodyRespDB,
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
		Path:   "/api/v1/categories/1",
	}.TestOnError(t, nil, errors.CategoryNotFound).Close()
}

func TestGetCategoriesWorks(t *testing.T) {
	existingAppAssert := CreateSampleCategory(t)

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

func TestUpdateCategoryWorks(t *testing.T) {
	existingAppAssert := CreateSampleCategory(t)

	generateNUCategory := *SampleCategoryFactory()
	generateNUCategory["name"] = cases.Title(language.English).String("GenerateNU")

	var AssertUpdatedCategoryBodyRespDB = func(app TestApp, assert *assert.A, resp *http.Response) {
		AssertCategoryWithIDBodyRespDB(app, assert, resp, 1, &generateNUCategory)
	}

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   "/api/v1/categories/1",
		Body:   &generateNUCategory,
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusOK,
			DBTester: AssertUpdatedCategoryBodyRespDB,
		},
	).Close()
}

func TestUpdateCategoryWorksWithSameDetails(t *testing.T) {
	existingAppAssert := CreateSampleCategory(t)

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   "/api/v1/categories/1",
		Body:   SampleCategoryFactory(),
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
			Path:   fmt.Sprintf("/api/v1/tags/%s", badRequest),
			Body:   SampleTagFactory(),
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

func TestDeleteCategoryWorks(t *testing.T) {
	existingAppAssert := CreateSampleCategory(t)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   "/api/v1/categories/1",
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusNoContent,
			DBTester: AssertNoCategories,
		},
	).Close()
}

func TestDeleteCategoryFailsBadRequest(t *testing.T) {
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
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

func TestDeleteCategoryFailsNotFound(t *testing.T) {
	TestRequest{
		Method: fiber.MethodDelete,
		Path:   "/api/v1/categories/1",
	}.TestOnError(t, nil, errors.CategoryNotFound).Close()
}
