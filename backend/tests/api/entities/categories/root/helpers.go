package root

import (
	"encoding/json"
	"net/http"

	"github.com/GenerateNU/sac/backend/entities/models"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
