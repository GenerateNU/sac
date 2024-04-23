package root

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/entities/models"
	categoriesRoot "github.com/GenerateNU/sac/backend/tests/api/entities/categories/root"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SampleTagFactory(categoryID uuid.UUID) *map[string]interface{} {
	return &map[string]interface{}{
		"name":        "Generate",
		"category_id": categoryID,
	}
}

func AssertTagWithBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respTag models.Tag

	err := json.NewDecoder(resp.Body).Decode(&respTag)

	eaa.Assert.NilError(err)

	var dbTag models.Tag

	err = eaa.App.Conn.First(&dbTag).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(dbTag.ID, respTag.ID)
	eaa.Assert.Equal(dbTag.Name, respTag.Name)
	eaa.Assert.Equal(dbTag.CategoryID, respTag.CategoryID)

	eaa.Assert.Equal((*body)["name"].(string), dbTag.Name)
	eaa.Assert.Equal((*body)["category_id"].(uuid.UUID), dbTag.CategoryID)

	return dbTag.ID
}

func AssertSampleTagBodyRespDB(t *testing.T, eaa h.ExistingAppAssert, resp *http.Response) uuid.UUID {
	appAssert, uuid := categoriesRoot.CreateSampleCategory(eaa)
	return AssertTagWithBodyRespDB(appAssert, resp, SampleTagFactory(uuid))
}

func CreateSampleTag(appAssert h.ExistingAppAssert) (existingAppAssert h.ExistingAppAssert, categoryUUID uuid.UUID, tagUUID uuid.UUID) {
	appAssert, categoryUUID = categoriesRoot.CreateSampleCategory(appAssert)

	AssertSampleTagBodyRespDB := func(eaa h.ExistingAppAssert, resp *http.Response) {
		tagUUID = AssertTagWithBodyRespDB(appAssert, resp, SampleTagFactory(categoryUUID))
	}

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/tags/",
			Body:   SampleTagFactory(categoryUUID),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: AssertSampleTagBodyRespDB,
		},
	)

	return appAssert, categoryUUID, tagUUID
}

func AssertNumTagsRemainsAtN(eaa h.ExistingAppAssert, resp *http.Response, n int) {
	var tags []models.Tag

	err := eaa.App.Conn.Find(&tags).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(n, len(tags))
}

func AssertNoTags(eaa h.ExistingAppAssert, resp *http.Response) {
	AssertNumTagsRemainsAtN(eaa, resp, 0)
}

func Assert1Tag(eaa h.ExistingAppAssert, resp *http.Response) {
	AssertNumTagsRemainsAtN(eaa, resp, 1)
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

func CreateSetOfTags(appAssert h.ExistingAppAssert) ([]uuid.UUID, h.ExistingAppAssert) {
	categories := categoriesRoot.SampleCategoriesFactory()

	categoryIDs := []uuid.UUID{}
	for _, category := range *categories {
		category := category
		appAssert = appAssert.TestOnStatusAndTester(
			h.TestRequest{
				Method: fiber.MethodPost,
				Path:   "/api/v1/categories/",
				Body:   &category,
				Role:   &models.Super,
			},
			h.TesterWithStatus{
				Status: fiber.StatusCreated,
				Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
					var respCategory models.Category

					err := json.NewDecoder(resp.Body).Decode(&respCategory)

					eaa.Assert.NilError(err)

					categoryIDs = append(categoryIDs, respCategory.ID)
				},
			},
		)
	}

	tags := SampleTagsFactory(categoryIDs)

	tagIDs := []uuid.UUID{}
	for _, tag := range *tags {
		tag := tag
		appAssert = appAssert.TestOnStatusAndTester(
			h.TestRequest{
				Method: fiber.MethodPost,
				Path:   "/api/v1/tags/",
				Body:   &tag,
				Role:   &models.Super,
			},
			h.TesterWithStatus{
				Status: fiber.StatusCreated,
				Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
					var respTag models.Tag

					err := json.NewDecoder(resp.Body).Decode(&respTag)

					eaa.Assert.NilError(err)

					tagIDs = append(tagIDs, respTag.ID)
				},
			},
		)
	}

	return tagIDs, appAssert
}

func AssertUserTagsRespDB(eaa h.ExistingAppAssert, resp *http.Response, id uuid.UUID) {
	var respTags []models.Tag

	err := json.NewDecoder(resp.Body).Decode(&respTags)

	eaa.Assert.NilError(err)

	var dbUser models.User

	err = eaa.App.Conn.First(&dbUser, id).Error

	eaa.Assert.NilError(err)

	var dbTags []models.Tag
	err = eaa.App.Conn.Model(&dbUser).Association("Tag").Find(&dbTags)

	eaa.Assert.NilError(err)

	for i, respTag := range respTags {
		eaa.Assert.Equal(respTag.ID, dbTags[i].ID)
		eaa.Assert.Equal(respTag.Name, dbTags[i].Name)
		eaa.Assert.Equal(respTag.CategoryID, dbTags[i].CategoryID)
	}
}

func AssertSampleUserTagsRespDB(eaa h.ExistingAppAssert, resp *http.Response, uuid uuid.UUID) {
	AssertUserTagsRespDB(eaa, resp, uuid)
}
