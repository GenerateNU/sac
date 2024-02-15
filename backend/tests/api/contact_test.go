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

func SampleContactFactory() *map[string]interface{} {
	return &map[string]interface{}{
		"type":    "email",
		"content": "jermaine@gmail.com",
	}
}

func ManyContactsFactory() map[string](*map[string]interface{}) {
	arr := make(map[string]*map[string]interface{})

	arr["email"] = &map[string]interface{}{
		"type":    "email",
		"content": "cheeseClub@gmail.com",
	}

	arr["youtube"] = &map[string]interface{}{
		"type":    "youtube",
		"content": "https://youtube.com/cheeseClub",
	}

	arr["facebook"] = &map[string]interface{}{
		"type":    "facebook",
		"content": "https://facebook.com/cheeseClub",
	}

	arr["discord"] = &map[string]interface{}{
		"type":    "discord",
		"content": "https://discord.com/cheeseClub",
	}

	arr["instagram"] = &map[string]interface{}{
		"type":    "instagram",
		"content": "https://instagram.com/cheeseClub",
	}
	arr["github"] = &map[string]interface{}{
		"type":    "github",
		"content": "https://github.com/cheeseClub",
	}

	return arr
}

func AssertContactBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respContact models.Contact

	err := json.NewDecoder(resp.Body).Decode(&respContact)

	eaa.Assert.NilError(err)

	var dbContacts []models.Contact

	err = eaa.App.Conn.Order("created_at desc").Find(&dbContacts).Error

	eaa.Assert.NilError(err)

	dbContact := dbContacts[0]

	eaa.Assert.Equal(dbContact.ID, respContact.ID)
	eaa.Assert.Equal(dbContact.Type, respContact.Type)
	eaa.Assert.Equal(dbContact.Content, respContact.Content)

	return dbContact.ID
}

func CreateSampleContact(existingAppAssert h.ExistingAppAssert) (eaa h.ExistingAppAssert, clubUUID uuid.UUID, contactUUID uuid.UUID) {
	appAssert, _, clubUUID := CreateSampleClub(existingAppAssert)

	var sampleContactUUID uuid.UUID

	return appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPut,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/contacts", clubUUID),
			Body:   SampleContactFactory(),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				sampleContactUUID = AssertContactBodyRespDB(eaa, resp, SampleContactFactory())
				AssertNumContactsRemainsAtN(eaa, resp, 1)
			},
		},
	), clubUUID, sampleContactUUID
}

func CreateManyContacts(existingAppAssert h.ExistingAppAssert) (eaa h.ExistingAppAssert, clubUUID uuid.UUID, contactUUIDs map[string]uuid.UUID) {
	existingAppAssert, _, clubUUID = CreateSampleClub(existingAppAssert)

	contactUUIDs = make(map[string]uuid.UUID)

	currentLength := 0
	for key, contact := range ManyContactsFactory() {
		existingAppAssert = existingAppAssert.TestOnStatusAndTester(h.TestRequest{
			Method: fiber.MethodPut,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/contacts", clubUUID),
			Body:   contact,
			Role:   &models.Super,
		},
			h.TesterWithStatus{
				Status: fiber.StatusOK,
				Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
					contactUUIDs[key] = AssertContactBodyRespDB(eaa, resp, contact)
					currentLength++
					AssertNumContactsRemainsAtN(eaa, resp, currentLength)
				},
			},
		)
	}

	return existingAppAssert, clubUUID, contactUUIDs
}

func TestCreateManyContactsWorks(t *testing.T) {
	t.Parallel()
	existingAppAssert, _, _ := CreateManyContacts(h.InitTest(t))
	existingAppAssert.Close()
}

func TestCreateContactWorks(t *testing.T) {
	t.Parallel()
	existingAppAssert, _, _ := CreateSampleContact(h.InitTest(t))
	existingAppAssert.Close()
}

func AssertNumContactsRemainsAtN(eaa h.ExistingAppAssert, resp *http.Response, n int) {
	var dbContacts []models.Contact

	err := eaa.App.Conn.Order("created_at desc").Find(&dbContacts).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(n, len(dbContacts))
}

func TestGetContactByIdWorks(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
	appAssert, _, _ := CreateManyContacts(h.InitTest(t))

	appAssert.TestOnStatus(h.TestRequest{
		Method: fiber.MethodGet,
		Path:   "/api/v1/contacts",
		Role:   &models.Super,
	}, fiber.StatusOK,
	).Close()
}
