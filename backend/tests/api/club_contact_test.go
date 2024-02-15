package tests

import (
	stdliberrors "errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AssertCreateBadContactDataFails(t *testing.T, jsonKey string, badValues []interface{}) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	for _, badValue := range badValues {
		sampleContactPermutation := *SampleContactFactory()
		sampleContactPermutation[jsonKey] = badValue

		appAssert = appAssert.TestOnErrorAndTester(h.TestRequest{
			Method: fiber.MethodPut,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/contacts", clubUUID),
			Body:   &sampleContactPermutation,
			Role:   &models.Super,
		},
			h.ErrorWithTester{
				Error: errors.FailedToValidateContact,
				Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
					AssertNumContactsRemainsAtN(eaa, resp, 0)
				},
			},
		)
	}
	appAssert.Close()
}

func TestCreateContactFailsOnInvalidType(t *testing.T) {
	t.Parallel()
	AssertCreateBadContactDataFails(t,
		"type",
		[]interface{}{
			"Not a valid type",
			"@#139081#$Ad_Axf",
		},
	)
}

func TestCreateContactFailsOnInvalidContent(t *testing.T) {
	t.Parallel()
	AssertCreateBadContactDataFails(t,
		"content",
		[]interface{}{
			"Not a valid url",
			"@#139081#$Ad_Axf",
		},
	)
}

func TestPutContactFailsOnClubIdNotExist(t *testing.T) {
	t.Parallel()
	appAssert, _, _ := CreateSampleClub(h.InitTest(t))

	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(h.TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/contacts", uuid),
		Body:   SampleContactFactory(),
		Role:   &models.Super,
	},
		h.ErrorWithTester{
			Error: errors.ClubNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var club models.Club

				err := eaa.App.Conn.Where("id = ?", uuid).First(&club).Error

				eaa.Assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestPutContactUpdatesExistingContact(t *testing.T) {
	t.Parallel()
	appAssert, clubUUID, contactUUID := CreateSampleContact(h.InitTest(t))

	updatedContact := SampleContactFactory()
	(*updatedContact)["content"] = "nedFlanders@gmail.com"

	appAssert.TestOnStatusAndTester(h.TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/contacts", clubUUID),
		Body:   updatedContact,
		Role:   &models.Super,
	},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var dbContact models.Contact

				err := eaa.App.Conn.Where("id = ?", contactUUID).First(&dbContact).Error

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(dbContact.Content, (*updatedContact)["content"])
			},
		},
	).Close()
}

func TestGetClubContacts(t *testing.T) {
	t.Parallel()
	appAssert, clubUUID, _ := CreateManyContacts(h.InitTest(t))

	appAssert.TestOnStatusAndTester(h.TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/contacts", clubUUID),
		Role:   &models.Super,
	},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var respContacts []models.Contact
				var dbContacts []models.Contact
				err := json.NewDecoder(resp.Body).Decode(&respContacts)
				eaa.Assert.NilError(err)

				err = eaa.App.Conn.Where("club_id = ?", clubUUID).Find(&dbContacts).Error
				eaa.Assert.NilError(err)

				eaa.Assert.Equal(len(respContacts), len(dbContacts))
			},
		},
	)

	appAssert.Close()
}
