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
)

func AssertTagsWithBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, body *[]map[string]interface{}) []uuid.UUID {
	var respTags []models.Tag

	err := json.NewDecoder(resp.Body).Decode(&respTags)

	eaa.Assert.NilError(err)

	var dbTags []models.Tag

	err = eaa.App.Conn.Find(&dbTags).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(len(dbTags), len(respTags))

	for i, dbTag := range dbTags {
		eaa.Assert.Equal(dbTag.ID, respTags[i].ID)
		eaa.Assert.Equal(dbTag.Name, respTags[i].Name)
		eaa.Assert.Equal(dbTag.CategoryID, respTags[i].CategoryID)
	}

	tagIDs := make([]uuid.UUID, len(dbTags))
	for i, dbTag := range dbTags {
		tagIDs[i] = dbTag.ID
	}

	return tagIDs
}

func TestGetCategoryTagsWorks(t *testing.T) {
	t.Parallel()
	appAssert, categoryUUID, tagID := CreateSampleTag(h.InitTest(t))

	body := SampleTagFactory(categoryUUID)
	(*body)["id"] = tagID

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/categories/%s/tags", categoryUUID),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertTagsWithBodyRespDB(eaa, resp, &[]map[string]interface{}{*body})
			},
		},
	).Close()
}

func TestGetCategoryTagsFailsCategoryBadRequest(t *testing.T) {
	t.Parallel()
	appAssert, _ := CreateSampleCategory(h.InitTest(t))

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
				Path:   fmt.Sprintf("/api/v1/categories/%s/tags", badRequest),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}

	appAssert.Close()
}

func TestGetCategoryTagsFailsCategoryNotFound(t *testing.T) {
	t.Parallel()
	appAssert, _ := CreateSampleCategory(h.InitTest(t))

	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/categories/%s/tags", uuid),
			Role:   &models.Super,
		}, h.ErrorWithTester{
			Error: errors.CategoryNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var category models.Category
				err := eaa.App.Conn.Where("id = ?", uuid).First(&category).Error
				eaa.Assert.Assert(err != nil)
			},
		},
	).Close()
}

func TestGetCategoryTagWorks(t *testing.T) {
	t.Parallel()
	existingAppAssert, categoryUUID, tagUUID := CreateSampleTag(h.InitTest(t))

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/categories/%s/tags/%s", categoryUUID, tagUUID),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertTagWithBodyRespDB(eaa, resp, SampleTagFactory(categoryUUID))
			},
		},
	).Close()
}

func TestGetCategoryTagFailsCategoryBadRequest(t *testing.T) {
	t.Parallel()
	appAssert, _, tagUUID := CreateSampleTag(h.InitTest(t))

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
				Path:   fmt.Sprintf("/api/v1/categories/%s/tags/%s", badRequest, tagUUID),
				Role:   &models.Super,
			}, errors.FailedToValidateID,
		)
	}

	appAssert.Close()
}

func TestGetCategoryTagFailsTagBadRequest(t *testing.T) {
	t.Parallel()
	appAssert, categoryUUID := CreateSampleCategory(h.InitTest(t))

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
				Path:   fmt.Sprintf("/api/v1/categories/%s/tags/%s", categoryUUID, badRequest),
				Role:   &models.Super,
			},
			errors.FailedToValidateID)
	}

	appAssert.Close()
}

func TestGetCategoryTagFailsCategoryNotFound(t *testing.T) {
	t.Parallel()
	appAssert, _, tagUUID := CreateSampleTag(h.InitTest(t))

	appAssert.TestOnError(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/categories/%s/tags/%s", uuid.New(), tagUUID),
			Role:   &models.Super,
		},
		errors.TagNotFound,
	).Close()
}

func TestGetCategoryTagFailsTagNotFound(t *testing.T) {
	t.Parallel()
	appAssert, categoryUUID := CreateSampleCategory(h.InitTest(t))

	appAssert.TestOnError(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/categories/%s/tags/%s", categoryUUID, uuid.New()),
			Role:   &models.Super,
		},
		errors.TagNotFound,
	).Close()
}
