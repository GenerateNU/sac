package tests

import (
	stdliberrors "errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/huandu/go-assert"
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

func AssertContactBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respContact models.Contact

	// decode the response body into a respContact
	err := json.NewDecoder(resp.Body).Decode(&respContact)

	assert.NilError(err)

	var dbContacts []models.Contact

	// get all contacts from the database ordered by created_at and store them in dbContacts
	err = app.Conn.Order("created_at desc").Find(&dbContacts).Error

	assert.NilError(err)

	dbContact := dbContacts[0]

	assert.Equal(dbContact.ID, respContact.ID)
	assert.Equal(dbContact.Type, respContact.Type)
	assert.Equal(dbContact.Content, respContact.Content)

	return dbContact.ID
}

func CreateSampleContact(t *testing.T, existingAppAssert *ExistingAppAssert) (eaa ExistingAppAssert, clubUUID uuid.UUID, contactUUID uuid.UUID) {
	appAssert, _, clubUUID := CreateSampleClub(t, existingAppAssert)

	var sampleContactUUID uuid.UUID


	appAssert = TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/contacts", clubUUID),
		Body:   SampleContactFactory(),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				sampleContactUUID = AssertContactBodyRespDB(app, assert, resp, SampleContactFactory())
				AssertNumContactsRemainsAtN(app, assert, resp, 1)
			},
		},
	)

	return appAssert, clubUUID, sampleContactUUID

}

func CreateManyContacts(t *testing.T, existingAppAssert *ExistingAppAssert) (eaa ExistingAppAssert, clubUUID uuid.UUID, contactUUIDs map[string]uuid.UUID) {
	appAssert, _, clubUUID := CreateSampleClub(t, existingAppAssert)

	contactUUIDs = make(map[string]uuid.UUID)

	currentLength := 0
	for key, contact := range ManyContactsFactory() {
		TestRequest{
			Method: fiber.MethodPut,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/contacts", clubUUID),
			Body:   contact,
		}.TestOnStatusAndDB(t, &appAssert,
			DBTesterWithStatus{
				Status: fiber.StatusOK,
				DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
					contactUUIDs[key] = AssertContactBodyRespDB(app, assert, resp, contact)
					currentLength++
					AssertNumContactsRemainsAtN(app, assert, resp, currentLength)
				},
			},
		)

	}

	return appAssert, clubUUID, contactUUIDs
}

func TestCreateManyContactsWorks(t *testing.T) {
	existingAppAssert, _, _ := CreateManyContacts(t, nil)
	existingAppAssert.Close()
}

func TestCreateContactWorks(t *testing.T) {
	existingAppAssert, _, _ := CreateSampleContact(t, nil)
	existingAppAssert.Close()
}

func AssertNumContactsRemainsAtN(app TestApp, assert *assert.A, resp *http.Response, n int) {
	var dbContacts []models.Contact

	err := app.Conn.Order("created_at desc").Find(&dbContacts).Error

	assert.NilError(err)

	assert.Equal(n, len(dbContacts))
}


func AssertCreateBadContactDataFails(t *testing.T, jsonKey string, badValues []interface{}) {
	appAssert, _, clubUUID := CreateSampleClub(t, nil)

	for _, badValue := range badValues {
		sampleContactPermutation := *SampleContactFactory()
		sampleContactPermutation[jsonKey] = badValue

		TestRequest{
			Method: fiber.MethodPut,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/contacts", clubUUID),
			Body:   &sampleContactPermutation,
		}.TestOnErrorAndDB(t, &appAssert,
			ErrorWithDBTester{
				Error:    errors.FailedToValidateContact,
				DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
					AssertNumContactsRemainsAtN(app, assert, resp, 0)
				},
			},
		)
	}
	appAssert.Close()
}

// if an invalid type is given, the request should fail
func TestCreateContactFailsOnInvalidType(t *testing.T) {
	AssertCreateBadContactDataFails(t,
		"type",
		[]interface{}{
			"Not a valid type",
			"@#139081#$Ad_Axf",
		},
	)
}

// if a bad link/email is given, the request should fail
func TestCreateContactFailsOnInvalidContent(t *testing.T) {
	AssertCreateBadContactDataFails(t,
		"content",
		[]interface{}{
			"Not a valid url",
			"@#139081#$Ad_Axf",
		},
	)
}

// if given an invalid club ID, the request should fail
func TestPutContactFailsOnClubIdNotExist(t *testing.T) {
	appAssert, _, _ := CreateSampleClub(t, nil)

	uuid := uuid.New()

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/contacts", uuid),
		Body:   SampleContactFactory(),
	}.TestOnErrorAndDB(t, &appAssert,
		ErrorWithDBTester{
			Error: errors.ClubNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var club models.Club

				err := app.Conn.Where("id = ?", uuid).First(&club).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

// if a club already has a contact of the same type, the new contact should replace the old one
func TestPutContactUpdatesExistingContact(t *testing.T){
	appAssert, clubUUID, contactUUID := CreateSampleContact(t, nil)

	updatedContact := SampleContactFactory()
	(*updatedContact)["content"] = "nedFlanders@gmail.com"

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/contacts", clubUUID),
		Body:   updatedContact,
	}.TestOnStatusAndDB(t, &appAssert, 
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var dbContact models.Contact

				err := app.Conn.Where("id = ?", contactUUID).First(&dbContact).Error

				assert.NilError(err)

				assert.Equal(dbContact.Content, (*updatedContact)["content"])
			},
		},
	).Close()
}

// given a valid contactID the request should return the contact
func TestGetContactByIdWorks(t *testing.T) {
	appAssert, _, contactUUID := CreateSampleContact(t, nil)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/contacts/%s", contactUUID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respContact models.Contact

				err := json.NewDecoder(resp.Body).Decode(&respContact)

				assert.NilError(err)

				var dbContacts []models.Contact

				err = app.Conn.Order("created_at desc").Find(&dbContacts).Error

				assert.NilError(err)

				assert.Equal(dbContacts[0].ID, respContact.ID)
				assert.Equal(dbContacts[0].Type, respContact.Type)
				assert.Equal(dbContacts[0].Content, respContact.Content)
			},
		},
	).Close()
}

// if a contactID does not exist, request should fail
func TestGetContactFailsOnContactIdNotExist(t *testing.T) {
	appAssert, _, _ := CreateSampleContact(t, nil)

	uuid := uuid.New()

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/contacts/%s", uuid),
	}.TestOnErrorAndDB(t, &appAssert,
		ErrorWithDBTester{
			Error: errors.ContactNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var contact models.Contact

				err := app.Conn.Where("id = ?", uuid).First(&contact).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

// given a valid contactID the request should delete the contact
func TestDeleteContactWorks(t *testing.T) {
	appAssert, _, contactUUID := CreateSampleContact(t, nil)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/contacts/%s", contactUUID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusNoContent,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var contact models.Contact

				err := app.Conn.Where("id = ?", contactUUID).First(&contact).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

// if a contactID does not exist, request should fail
func TestDeleteContactFailsOnContactIdNotExist(t *testing.T) {
	appAssert, _, _ := CreateSampleContact(t, nil)
	uuid := uuid.New()
	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/contacts/%s", uuid),
	}.TestOnErrorAndDB(t, &appAssert,
		ErrorWithDBTester{
			Error: errors.ContactNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var contact models.Contact
				err := app.Conn.Where("id = ?", uuid).First(&contact).Error
				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))

				AssertNumContactsRemainsAtN(app, assert, resp, 1)
			},
		},
	).Close()
}

// test that the request returns paginated contacts
func TestGetContactsWorks(t *testing.T) {
	appAssert, _, _ := CreateManyContacts(t, nil)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   "/api/v1/contacts",
	}.TestOnStatus(t, &appAssert, fiber.StatusOK)

	appAssert.Close()
}

// test that the request returns contacts that belong to a club
func TestGetClubContacts(t *testing.T) {
	appAssert, clubUUID, _ := CreateManyContacts(t, nil)
	// TODO create another contact that does not belong to the current club 
	// and assert that it is not returned in the response
	// using createSampleClub twice returns a 409 ID conflict error
	// so I may have to create a club manually

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/contacts", clubUUID),
	}.TestOnStatusAndDB(t, &appAssert, 
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respContacts []models.Contact
				var dbContacts []models.Contact
				err := json.NewDecoder(resp.Body).Decode(&respContacts)
				assert.NilError(err)

				err = app.Conn.Where("club_id = ?", clubUUID).Find(&dbContacts).Error
				assert.NilError(err)

				assert.Equal(len(respContacts), len(dbContacts))
			},
		},
	)

	appAssert.Close()
}