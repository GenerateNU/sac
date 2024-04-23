package root

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	categoriesRoot "github.com/GenerateNU/sac/backend/tests/api/entities/categories/root"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func TestCreateTagWorks(t *testing.T) {
	appAssert, _, _ := CreateSampleTag(h.InitTest(t))
	appAssert.Close()
}

func TestCreateTagFailsBadRequest(t *testing.T) {
	appAssert := h.InitTest(t)

	badBodys := []map[string]interface{}{
		{
			"name":        "Generate",
			"category_id": "1",
		},
		{
			"name":        1,
			"category_id": 1,
		},
	}

	for _, badBody := range badBodys {
		badBody := badBody
		appAssert = appAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodPost,
				Path:   "/api/v1/tags/",
				Body:   &badBody,
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error:  errors.FailedToParseRequestBody,
				Tester: AssertNoTags,
			},
		)
	}

	appAssert.Close()
}

func TestCreateTagFailsValidation(t *testing.T) {
	appAssert := h.InitTest(t)

	badBodys := []map[string]interface{}{
		{
			"name": "Generate",
		},
		{
			"category_id": uuid.New(),
		},
		{},
	}

	for _, badBody := range badBodys {
		badBody := badBody
		appAssert = appAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodPost,
				Path:   "/api/v1/tags/",
				Body:   &badBody,
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error:  errors.FailedToValidateTag,
				Tester: AssertNoTags,
			},
		)
	}

	appAssert.Close()
}

func TestGetTagWorks(t *testing.T) {
	existingAppAssert, categoryUUID, tagUUID := CreateSampleTag(h.InitTest(t))

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
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

func TestGetTagFailsBadRequest(t *testing.T) {
	appAssert := h.InitTest(t)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		appAssert = appAssert.TestOnError(
			h.TestRequest{
				Method: fiber.MethodGet,
				Path:   fmt.Sprintf("/api/v1/tags/%s", badRequest),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}

	appAssert.Close()
}

func TestGetTagFailsNotFound(t *testing.T) {
	h.InitTest(t).TestOnError(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/tags/%s", uuid.New()),
			Role:   &models.Super,
		},
		errors.TagNotFound,
	).Close()
}

func TestUpdateTagWorksUpdateName(t *testing.T) {
	existingAppAssert, categoryUUID, tagUUID := CreateSampleTag(h.InitTest(t))

	generateNUTag := *SampleTagFactory(categoryUUID)
	generateNUTag["name"] = "GenerateNU"

	AssertUpdatedTagBodyRespDB := func(eaa h.ExistingAppAssert, resp *http.Response) {
		tagUUID = AssertTagWithBodyRespDB(eaa, resp, &generateNUTag)
	}

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
			Body:   &generateNUTag,
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: AssertUpdatedTagBodyRespDB,
		},
	).Close()
}

func TestUpdateTagWorksUpdateCategory(t *testing.T) {
	existingAppAssert, _, tagUUID := CreateSampleTag(h.InitTest(t))

	technologyCategory := *categoriesRoot.SampleCategoryFactory()
	technologyCategory["name"] = "Technology"

	var technologyCategoryUUID uuid.UUID

	AssertNewCategoryBodyRespDB := func(eaa h.ExistingAppAssert, resp *http.Response) {
		technologyCategoryUUID = categoriesRoot.AssertCategoryWithBodyRespDBMostRecent(eaa, resp, &technologyCategory)
	}

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/categories/",
			Body:   &technologyCategory,
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: AssertNewCategoryBodyRespDB,
		},
	)

	technologyTag := *SampleTagFactory(technologyCategoryUUID)

	AssertUpdatedTagBodyRespDB := func(eaa h.ExistingAppAssert, resp *http.Response) {
		AssertTagWithBodyRespDB(eaa, resp, &technologyTag)
	}

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
			Body:   &technologyTag,
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: AssertUpdatedTagBodyRespDB,
		},
	).Close()
}

func TestUpdateTagWorksWithSameDetails(t *testing.T) {
	existingAppAssert, categoryUUID, tagUUID := CreateSampleTag(h.InitTest(t))

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
			Body:   SampleTagFactory(categoryUUID),
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

func TestUpdateTagFailsBadRequest(t *testing.T) {
	appAssert, uuid := categoriesRoot.CreateSampleCategory(h.InitTest(t))

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		appAssert = appAssert.TestOnError(
			h.TestRequest{
				Method: fiber.MethodPatch,
				Path:   fmt.Sprintf("/api/v1/tags/%s", badRequest),
				Body:   SampleTagFactory(uuid),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}

	appAssert.Close()
}

func TestDeleteTagWorks(t *testing.T) {
	existingAppAssert, _, tagUUID := CreateSampleTag(h.InitTest(t))

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: AssertNoTags,
		},
	).Close()
}

func TestDeleteTagFailsBadRequest(t *testing.T) {
	appAssert, _, _ := CreateSampleTag(h.InitTest(t))

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		appAssert = appAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodDelete,
				Path:   fmt.Sprintf("/api/v1/tags/%s", badRequest),
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error:  errors.FailedToValidateID,
				Tester: Assert1Tag,
			},
		)
	}

	appAssert.Close()
}

func TestDeleteTagFailsNotFound(t *testing.T) {
	appAssert, _, _ := CreateSampleTag(h.InitTest(t))

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/tags/%s", uuid.New()),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error:  errors.TagNotFound,
			Tester: Assert1Tag,
		},
	).Close()
}
