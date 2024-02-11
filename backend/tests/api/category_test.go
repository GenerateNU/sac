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

	"github.com/goccy/go-json"
)

func SampleCategoryFactory() *map[string]interface{} {
	return &map[string]interface{}{
		"name": "Foo",
	}
}

func AssertCategoryBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respCategory models.Category

	err := json.NewDecoder(resp.Body).Decode(&respCategory)

	eaa.Assert.NilError(err)

	var dbCategories []models.Category

	err = eaa.App.Conn.Find(&dbCategories).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(1, len(dbCategories))

	dbCategory := dbCategories[0]

	eaa.Assert.Equal(dbCategory.ID, respCategory.ID)
	eaa.Assert.Equal(dbCategory.Name, respCategory.Name)

	eaa.Assert.Equal((*body)["name"].(string), dbCategory.Name)

	return dbCategory.ID
}

func AssertCategoryWithBodyRespDBMostRecent(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respCategory models.Category

	err := json.NewDecoder(resp.Body).Decode(&respCategory)

	eaa.Assert.NilError(err)

	var dbCategory models.Category

	err = eaa.App.Conn.Order("created_at desc").First(&dbCategory).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(dbCategory.ID, respCategory.ID)
	eaa.Assert.Equal(dbCategory.Name, respCategory.Name)

	eaa.Assert.Equal((*body)["name"].(string), dbCategory.Name)

	return dbCategory.ID
}

func AssertSampleCategoryBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response) uuid.UUID {
	return AssertCategoryBodyRespDB(eaa, resp, SampleCategoryFactory())
}

func CreateSampleCategory(existingAppAssert h.ExistingAppAssert) (h.ExistingAppAssert, uuid.UUID) {
	var sampleCategoryUUID uuid.UUID

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/categories/",
			Body:   SampleCategoryFactory(),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				sampleCategoryUUID = AssertSampleCategoryBodyRespDB(eaa, resp)
			},
		},
	)

	return existingAppAssert, sampleCategoryUUID
}

func TestCreateCategoryWorks(t *testing.T) {
	existingAppAssert, _ := CreateSampleCategory(h.InitTest(t))
	existingAppAssert.Close()
}

func TestCreateCategoryIgnoresid(t *testing.T) {
	h.InitTest(t).TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/categories/",
			Body: &map[string]interface{}{
				"id":   12,
				"name": "Foo",
			},
			Role: &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertSampleCategoryBodyRespDB(eaa, resp)
			},
		},
	).Close()
}

func Assert1Category(eaa h.ExistingAppAssert, resp *http.Response) {
	AssertNumCategoriesRemainsAtN(eaa, resp, 1)
}

func AssertNoCategories(eaa h.ExistingAppAssert, resp *http.Response) {
	AssertNumCategoriesRemainsAtN(eaa, resp, 0)
}

func AssertNumCategoriesRemainsAtN(eaa h.ExistingAppAssert, resp *http.Response, n int) {
	var categories []models.Category

	err := eaa.App.Conn.Find(&categories).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(n, len(categories))
}

func TestCreateCategoryFailsIfNameIsNotString(t *testing.T) {
	h.InitTest(t).TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/categories/",
			Body: &map[string]interface{}{
				"name": 1231,
			},
			Role: &models.Super,
		},
		h.ErrorWithTester{
			Error:  errors.FailedToParseRequestBody,
			Tester: AssertNoCategories,
		},
	).Close()
}

func TestCreateCategoryFailsIfNameIsMissing(t *testing.T) {
	h.InitTest(t).TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/categories/",
			Body:   &map[string]interface{}{},
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error:  errors.FailedToValidateCategory,
			Tester: AssertNoCategories,
		},
	).Close()
}

func TestCreateCategoryFailsIfCategoryWithThatNameAlreadyExists(t *testing.T) {
	existingAppAssert, _ := CreateSampleCategory(h.InitTest(t))

	TestNumCategoriesRemainsAt1 := func(eaa h.ExistingAppAssert, resp *http.Response) {
		AssertNumCategoriesRemainsAtN(eaa, resp, 1)
	}

	for _, permutation := range h.AllCasingPermutations((*SampleCategoryFactory())["name"].(string)) {
		modifiedSampleCategoryBody := *SampleCategoryFactory()
		modifiedSampleCategoryBody["name"] = permutation

		existingAppAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodPost,
				Path:   "/api/v1/categories/",
				Body:   &modifiedSampleCategoryBody,
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error:  errors.CategoryAlreadyExists,
				Tester: TestNumCategoriesRemainsAt1,
			},
		)
	}

	existingAppAssert.Close()
}

func TestGetCategoryWorks(t *testing.T) {
	existingAppAssert, uuid := CreateSampleCategory(h.InitTest(t))

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/categories/%s", uuid),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertSampleCategoryBodyRespDB(eaa, resp)
			},
		},
	).Close()
}

func TestGetCategoryFailsBadRequest(t *testing.T) {
	appAssert := h.InitTest(t)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		appAssert.TestOnError(
			h.TestRequest{
				Method: fiber.MethodGet,
				Path:   fmt.Sprintf("/api/v1/categories/%s", badRequest),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}

	appAssert.Close()
}

func TestGetCategoryFailsNotFound(t *testing.T) {
	h.InitTest(t).TestOnError(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/categories/%s", uuid.New()),
			Role:   &models.Super,
		}, errors.CategoryNotFound,
	).Close()
}

func TestGetCategoriesWorks(t *testing.T) {
	existingAppAssert, _ := CreateSampleCategory(h.InitTest(t))

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   "/api/v1/categories/",
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var categories []models.Category

				err := eaa.App.Conn.Find(&categories).Error

				eaa.Assert.NilError(err)

				var respCategories []models.Category

				err = json.NewDecoder(resp.Body).Decode(&respCategories)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(1, len(respCategories))
				eaa.Assert.Equal(1, len(categories))

				eaa.Assert.Equal(categories[0].ID, respCategories[0].ID)

				eaa.Assert.Equal(categories[0].Name, respCategories[0].Name)

				eaa.Assert.Equal((*SampleCategoryFactory())["name"].(string), categories[0].Name)
			},
		},
	).Close()
}

func AssertUpdatedCategoryBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) {
	var respCategory models.Category

	err := json.NewDecoder(resp.Body).Decode(&respCategory)

	eaa.Assert.NilError(err)

	var dbCategories []models.Category

	err = eaa.App.Conn.Find(&dbCategories).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(1, len(dbCategories))

	dbCategory := dbCategories[0]

	eaa.Assert.Equal(dbCategory.ID, respCategory.ID)
	eaa.Assert.Equal(dbCategory.Name, respCategory.Name)

	eaa.Assert.Equal((*body)["id"].(uuid.UUID), dbCategory.ID)
	eaa.Assert.Equal((*body)["name"].(string), dbCategory.Name)
}

func TestUpdateCategoryWorks(t *testing.T) {
	existingAppAssert, uuid := CreateSampleCategory(h.InitTest(t))

	category := map[string]interface{}{
		"id":   uuid,
		"name": "Arts & Crafts",
	}

	AssertUpdatedCategoryBodyRespDB := func(eaa h.ExistingAppAssert, resp *http.Response) {
		AssertUpdatedCategoryBodyRespDB(eaa, resp, &category)
	}

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/categories/%s", uuid),
			Body:   &category,
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: AssertUpdatedCategoryBodyRespDB,
		},
	).Close()
}

func TestUpdateCategoryWorksWithSameDetails(t *testing.T) {
	existingAppAssert, uuid := CreateSampleCategory(h.InitTest(t))

	category := *SampleCategoryFactory()
	category["id"] = uuid

	AssertSampleCategoryBodyRespDB := func(eaa h.ExistingAppAssert, resp *http.Response) {
		AssertUpdatedCategoryBodyRespDB(eaa, resp, &category)
	}

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/categories/%s", uuid),
			Body:   &category,
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: AssertSampleCategoryBodyRespDB,
		},
	).Close()
}

func TestUpdateCategoryFailsBadRequest(t *testing.T) {
	appAssert := h.InitTest(t)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		appAssert.TestOnError(
			h.TestRequest{
				Method: fiber.MethodPatch,
				Path:   fmt.Sprintf("/api/v1/categories/%s", badRequest),
				Body:   SampleCategoryFactory(),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}

	appAssert.Close()
}

func TestDeleteCategoryWorks(t *testing.T) {
	existingAppAssert, uuid := CreateSampleCategory(h.InitTest(t))

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/categories/%s", uuid),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: AssertNoCategories,
		},
	).Close()
}

func TestDeleteCategoryFailsBadRequest(t *testing.T) {
	existingAppAssert, _ := CreateSampleCategory(h.InitTest(t))

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		existingAppAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodDelete,
				Path:   fmt.Sprintf("/api/v1/categories/%s", badRequest),
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error:  errors.FailedToValidateID,
				Tester: Assert1Category,
			},
		)
	}

	existingAppAssert.Close()
}

func TestDeleteCategoryFailsNotFound(t *testing.T) {
	existingAppAssert, _ := CreateSampleCategory(h.InitTest(t))

	existingAppAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/categories/%s", uuid.New()),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error:  errors.CategoryNotFound,
			Tester: Assert1Category,
		},
	).Close()
}
