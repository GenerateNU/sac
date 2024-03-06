package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func CreateSetOfTags(appAssert h.ExistingAppAssert) ([]uuid.UUID, h.ExistingAppAssert) {
	categories := SampleCategoriesFactory()

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

func TestCreateUserTagsFailsOnInvalidDataType(t *testing.T) {
	appAssert := h.InitTest(t)

	invalidTags := []interface{}{
		[]string{"1", "2", "34"},
		[]models.Tag{{Name: "Test", CategoryID: uuid.UUID{}}, {Name: "Test2", CategoryID: uuid.UUID{}}},
		[]float32{1.32, 23.5, 35.1},
	}

	for _, tag := range invalidTags {
		malformedTag := *SampleTagIDsFactory(nil)
		malformedTag["tags"] = tag

		appAssert = appAssert.TestOnError(
			h.TestRequest{
				Method:             fiber.MethodPost,
				Path:               "/api/v1/users/:userID/tags/",
				Body:               &malformedTag,
				Role:               &models.Student,
				TestUserIDReplaces: h.StringToPointer(":userID"),
			},
			errors.FailedToParseRequestBody,
		)
	}

	appAssert.Close()
}

func TestCreateUserTagsFailsOnInvalidUserID(t *testing.T) {
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
				Method: fiber.MethodPost,
				Path:   fmt.Sprintf("/api/v1/users/%s/tags", badRequest),
				Body:   SampleTagIDsFactory(nil),
				Role:   &models.Student,
			},
			errors.FailedToValidateID,
		)
	}

	appAssert.Close()
}

type UUIDSlice []uuid.UUID

func TestCreateUserTagsFailsOnInvalidKey(t *testing.T) {
	appAssert := h.InitTest(t)

	invalidBody := []map[string]interface{}{
		{
			"tag": UUIDSlice{uuid.New(), uuid.New()},
		},
		{
			"tagIDs": []uint{1, 2, 3},
		},
	}

	for _, body := range invalidBody {
		body := body

		appAssert = appAssert.TestOnError(
			h.TestRequest{
				Method:             fiber.MethodPost,
				Path:               "/api/v1/users/:userID/tags/",
				Body:               &body,
				Role:               &models.Student,
				TestUserIDReplaces: h.StringToPointer(":userID"),
			},
			errors.FailedToValidateUserTags,
		)
	}

	appAssert.Close()
}

func TestCreateUserTagsFailsOnNonExistentUser(t *testing.T) {
	uuid := uuid.New()

	h.InitTest(t).TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   fmt.Sprintf("/api/v1/users/%s/tags/", uuid),
			Body:   SampleTagIDsFactory(nil),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error: errors.UserNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var dbUser models.User
				err := eaa.App.Conn.First(&dbUser, uuid).Error

				eaa.Assert.Assert(err != nil)
			},
		},
	).Close()
}

func TestCreateUserTagsWorks(t *testing.T) {
	tagUUIDs, appAssert := CreateSetOfTags(h.InitTest(t))

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               "/api/v1/users/:userID/tags/",
			Body:               SampleTagIDsFactory(&tagUUIDs),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertSampleUserTagsRespDB(eaa, resp, eaa.App.TestUser.UUID)
			},
		},
	)

	appAssert.Close()
}

func TestCreateUserTagsNoneAddedIfInvalid(t *testing.T) {
	h.InitTest(t).TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               "/api/v1/users/:userID/tags/",
			Body:               SampleTagIDsFactory(nil),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var respTags []models.Tag

				err := json.NewDecoder(resp.Body).Decode(&respTags)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(len(respTags), 0)
			},
		},
	).Close()
}

func TestGetUserTagsFailsOnNonExistentUser(t *testing.T) {
	uuid := uuid.New()

	h.InitTest(t).TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/users/%s/tags/", uuid),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error: errors.UserNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var dbUser models.User

				err := eaa.App.Conn.First(&dbUser, uuid).Error

				eaa.Assert.Assert(err != nil)
			},
		},
	).Close()
}

func TestGetUserTagsReturnsEmptyListWhenNoneAdded(t *testing.T) {
	h.InitTest(t).TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodGet,
			Path:               "/api/v1/users/:userID/tags/",
			Role:               &models.Student,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var respTags []models.Tag

				err := json.NewDecoder(resp.Body).Decode(&respTags)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(len(respTags), 0)
			},
		},
	).Close()
}

func TestGetUserTagsReturnsCorrectList(t *testing.T) {
	tagUUIDs, appAssert := CreateSetOfTags(h.InitTest(t))

	appAssert.TestOnStatus(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               "/api/v1/users/:userID/tags/",
			Body:               SampleTagIDsFactory(&tagUUIDs),
			Role:               &models.Student,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		fiber.StatusCreated,
	).TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodGet,
			Path:               "/api/v1/users/:userID/tags/",
			Role:               &models.Student,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertSampleUserTagsRespDB(eaa, resp, eaa.App.TestUser.UUID)
			},
		},
	).Close()
}

func TestDeleteUserTagFailsOnNonExistentUser(t *testing.T) {
	userID := uuid.New()
	tagID := uuid.New()

	h.InitTest(t).TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/users/%s/tags/%s/", userID, tagID),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error: errors.UserNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var dbUser models.User

				err := eaa.App.Conn.First(&dbUser, userID).Error

				eaa.Assert.Assert(err != nil)
			},
		},
	).Close()
}

func TestDeleteUserTagFailsOnNonExistentTag(t *testing.T) {
	tagID := uuid.New()

	h.InitTest(t).TestOnErrorAndTester(
		h.TestRequest{
			Method:             fiber.MethodDelete,
			Path:               fmt.Sprintf("/api/v1/users/:userID/tags/%s/", tagID),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.ErrorWithTester{
			Error: errors.TagNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var dbTag models.Tag

				err := eaa.App.Conn.First(&dbTag, tagID).Error
				eaa.Assert.Assert(err != nil)
			},
		},
	).Close()
}

func TestDeleteUserTagFailsOnInvalidUserUUID(t *testing.T) {
	appAssert := h.InitTest(t)

	badUserUUIDs := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badUserUUID := range badUserUUIDs {
		appAssert = appAssert.TestOnError(
			h.TestRequest{
				Method: fiber.MethodDelete,
				Path:   fmt.Sprintf("/api/v1/users/%s/tags/%s/", badUserUUID, uuid.New()),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}

	appAssert.Close()
}

func TestDeleteUserTagFailsOnInvalidTagUUID(t *testing.T) {
	appAssert := h.InitTest(t)

	badTagUUIDs := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badTagUUID := range badTagUUIDs {
		appAssert = appAssert.TestOnError(
			h.TestRequest{
				Method:             fiber.MethodDelete,
				Path:               fmt.Sprintf("/api/v1/users/:userID/tags/%s/", badTagUUID),
				Role:               &models.Super,
				TestUserIDReplaces: h.StringToPointer(":userID"),
			},
			errors.FailedToValidateID,
		)
	}

	appAssert.Close()

}

func TestDeleteUserTagDoesNotAlterTagListOnNonAssociation(t *testing.T) {
	tagUUIDs, appAssert := CreateSetOfTags(h.InitTest(t))
	appAssert.Assert.Assert(len(tagUUIDs) > 1)

	// Tag to be queried:
	tagID := tagUUIDs[0]

	// Tags to be added to the user:
	tagUUIDs = tagUUIDs[1:]

	appAssert.TestOnStatus(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               "/api/v1/users/:userID/tags/",
			Body:               SampleTagIDsFactory(&tagUUIDs),
			Role:               &models.Student,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		fiber.StatusCreated,
	)

	userTagsBeforeDeletion, err := transactions.GetUserTags(appAssert.App.Conn, appAssert.App.TestUser.UUID)
	appAssert.Assert.NilError(&err)

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodDelete,
			Path:               fmt.Sprintf("/api/v1/users/:userID/tags/%s/", tagID),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var dbUser models.User

				err := eaa.App.Conn.Where("id = ?", appAssert.App.TestUser.UUID).Preload("Tag").First(&dbUser).Error
				eaa.Assert.NilError(err)

				eaa.Assert.Equal(dbUser.Tag, userTagsBeforeDeletion)
			},
		},
	).Close()
}

func TestDeleteUserTagRemovesTagFromUser(t *testing.T) {
	tagUUIDs, appAssert := CreateSetOfTags(h.InitTest(t))
	appAssert.Assert.Assert(len(tagUUIDs) > 1)

	tagID := tagUUIDs[0]

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               "/api/v1/users/:userID/tags/",
			Body:               SampleTagIDsFactory(&tagUUIDs),
			Role:               &models.Student,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var dbUser models.User

				err := eaa.App.Conn.Where("id = ?", eaa.App.TestUser.UUID).Preload("Tag").First(&dbUser)
				eaa.Assert.NilError(err)

				eaa.Assert.Equal(len(dbUser.Tag), len(tagUUIDs))

				var dbTag models.Tag

				err = eaa.App.Conn.Where("id = ?", tagID).Preload("User").First(&dbTag)
				eaa.Assert.NilError(err)

				eaa.Assert.Equal(len(dbTag.User), 1)
			},
		},
	).TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodDelete,
			Path:               fmt.Sprintf("/api/v1/users/:userID/tags/%s/", tagID),
			Role:               &models.Student,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var dbUser models.User

				err := eaa.App.Conn.Where("id = ?", eaa.App.TestUser.UUID).Preload("Tag").First(&dbUser)
				eaa.Assert.NilError(err)

				eaa.Assert.Equal(len(dbUser.Tag), len(tagUUIDs)-1)

				var dbTag models.Tag

				err = eaa.App.Conn.Where("id = ?", tagID).Preload("User").First(&dbTag)
				eaa.Assert.NilError(err)

				eaa.Assert.Equal(len(dbTag.User), 0)
			},
		},
	)
}
