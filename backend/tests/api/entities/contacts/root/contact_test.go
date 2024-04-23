package root

import (
	stdliberrors "errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestCreateManyContactsWorks(t *testing.T) {
	existingAppAssert, _, _ := CreateManyContacts(h.InitTest(t))
	existingAppAssert.Close()
}

func TestCreateContactWorks(t *testing.T) {
	existingAppAssert, _, _ := CreateSampleContact(h.InitTest(t))
	existingAppAssert.Close()
}

func TestGetContactByIdWorks(t *testing.T) {
	appAssert, _, contactUUID := CreateSampleContact(h.InitTest(t))

	appAssert.TestOnStatusAndTester(h.TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/contacts/%s", contactUUID),
		Role:   &models.Super,
	},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var respContact models.Contact

				err := json.NewDecoder(resp.Body).Decode(&respContact)

				eaa.Assert.NilError(err)

				var dbContacts []models.Contact

				err = eaa.App.Conn.Order("created_at desc").Find(&dbContacts).Error

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(dbContacts[0].ID, respContact.ID)
				eaa.Assert.Equal(dbContacts[0].Type, respContact.Type)
				eaa.Assert.Equal(dbContacts[0].Content, respContact.Content)
			},
		},
	).Close()
}

func TestGetContactFailsOnContactIdNotExist(t *testing.T) {
	appAssert, _, _ := CreateSampleContact(h.InitTest(t))

	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(h.TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/contacts/%s", uuid),
		Role:   &models.Super,
	},
		h.ErrorWithTester{
			Error: errors.ContactNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var contact models.Contact

				err := eaa.App.Conn.Where("id = ?", uuid).First(&contact).Error

				eaa.Assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestDeleteContactWorks(t *testing.T) {
	appAssert, _, contactUUID := CreateSampleContact(h.InitTest(t))

	appAssert.TestOnStatusAndTester(h.TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/contacts/%s", contactUUID),
		Role:   &models.Super,
	},
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var contact models.Contact

				err := eaa.App.Conn.Where("id = ?", contactUUID).First(&contact).Error

				eaa.Assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestDeleteContactFailsOnContactIdNotExist(t *testing.T) {
	appAssert, _, _ := CreateSampleContact(h.InitTest(t))
	uuid := uuid.New()

	appAssert.TestOnErrorAndTester(h.TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/contacts/%s", uuid),
		Role:   &models.Super,
	},
		h.ErrorWithTester{
			Error: errors.ContactNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var contact models.Contact
				err := eaa.App.Conn.Where("id = ?", uuid).First(&contact).Error
				eaa.Assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))

				AssertNumContactsRemainsAtN(eaa, resp, 1)
			},
		},
	).Close()
}

// test that the request returns paginated contacts
func TestGetContactsWorks(t *testing.T) {
	appAssert, _, _ := CreateManyContacts(h.InitTest(t))

	appAssert.TestOnStatus(h.TestRequest{
		Method: fiber.MethodGet,
		Path:   "/api/v1/contacts",
		Role:   &models.Super,
	}, fiber.StatusOK,
	).Close()
}
