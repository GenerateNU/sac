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

func SampleClubFactory(userID uuid.UUID) *map[string]interface{} {
	return &map[string]interface{}{
		"user_id":           userID,
		"name":              "Generate",
		"preview":           "Generate is Northeastern's premier student-led product development studio.",
		"description":       "https://mongodb.com",
		"num_members":       1,
		"is_recruiting":     true,
		"recruitment_cycle": "always",
		"recruitment_type":  "application",
		"application_link":  "https://generatenu.com/apply",
		"logo":              "https://aws.amazon.com/s3/",
	}
}

func AssertClubBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respClub models.Club

	err := json.NewDecoder(resp.Body).Decode(&respClub)

	assert.NilError(err)

	var dbClubs []models.Club

	err = app.Conn.Order("created_at desc").Find(&dbClubs).Error

	assert.NilError(err)

	assert.Equal(2, len(dbClubs))

	dbClub := dbClubs[0]

	assert.Equal(dbClub.ID, respClub.ID)
	assert.Equal(dbClub.Name, respClub.Name)
	assert.Equal(dbClub.Preview, respClub.Preview)
	assert.Equal(dbClub.Description, respClub.Description)
	assert.Equal(dbClub.NumMembers, respClub.NumMembers)
	assert.Equal(dbClub.IsRecruiting, respClub.IsRecruiting)
	assert.Equal(dbClub.RecruitmentCycle, respClub.RecruitmentCycle)
	assert.Equal(dbClub.RecruitmentType, respClub.RecruitmentType)
	assert.Equal(dbClub.ApplicationLink, respClub.ApplicationLink)
	assert.Equal(dbClub.Logo, respClub.Logo)

	var dbAdmins []models.User

	err = app.Conn.Model(&dbClub).Association("Admin").Find(&dbAdmins)

	assert.NilError(err)

	assert.Equal(1, len(dbAdmins))

	assert.Equal((*body)["user_id"].(uuid.UUID), dbAdmins[0].ID)
	assert.Equal((*body)["name"].(string), dbClub.Name)
	assert.Equal((*body)["preview"].(string), dbClub.Preview)
	assert.Equal((*body)["description"].(string), dbClub.Description)
	assert.Equal((*body)["num_members"].(int), dbClub.NumMembers)
	assert.Equal((*body)["is_recruiting"].(bool), dbClub.IsRecruiting)
	assert.Equal(models.RecruitmentCycle((*body)["recruitment_cycle"].(string)), dbClub.RecruitmentCycle)
	assert.Equal(models.RecruitmentType((*body)["recruitment_type"].(string)), dbClub.RecruitmentType)
	assert.Equal((*body)["application_link"].(string), dbClub.ApplicationLink)
	assert.Equal((*body)["logo"].(string), dbClub.Logo)

	return dbClub.ID
}

func AssertClubWithBodyRespDBMostRecent(app TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respClub models.Club

	err := json.NewDecoder(resp.Body).Decode(&respClub)

	assert.NilError(err)

	var dbClub models.Club

	err = app.Conn.Order("created_at desc").First(&dbClub).Error

	assert.NilError(err)

	assert.Equal(dbClub.ID, respClub.ID)
	assert.Equal(dbClub.Name, respClub.Name)
	assert.Equal(dbClub.Preview, respClub.Preview)
	assert.Equal(dbClub.Description, respClub.Description)
	assert.Equal(dbClub.NumMembers, respClub.NumMembers)
	assert.Equal(dbClub.IsRecruiting, respClub.IsRecruiting)
	assert.Equal(dbClub.RecruitmentCycle, respClub.RecruitmentCycle)
	assert.Equal(dbClub.RecruitmentType, respClub.RecruitmentType)
	assert.Equal(dbClub.ApplicationLink, respClub.ApplicationLink)
	assert.Equal(dbClub.Logo, respClub.Logo)

	var dbAdmins []models.User

	err = app.Conn.Model(&dbClub).Association("Admins").Find(&dbAdmins)

	assert.NilError(err)

	assert.Equal(1, len(dbAdmins))

	dbAdmin := dbAdmins[0]

	assert.Equal((*body)["user_id"].(uuid.UUID), dbAdmin.ID)
	assert.Equal((*body)["name"].(string), dbClub.Name)
	assert.Equal((*body)["preview"].(string), dbClub.Preview)
	assert.Equal((*body)["description"].(string), dbClub.Description)
	assert.Equal((*body)["num_members"].(int), dbClub.NumMembers)
	assert.Equal((*body)["is_recruiting"].(bool), dbClub.IsRecruiting)
	assert.Equal((*body)["recruitment_cycle"].(string), dbClub.RecruitmentCycle)
	assert.Equal((*body)["recruitment_type"].(string), dbClub.RecruitmentType)
	assert.Equal((*body)["application_link"].(string), dbClub.ApplicationLink)
	assert.Equal((*body)["logo"].(string), dbClub.Logo)

	return dbClub.ID
}

func AssertSampleClubBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, userID uuid.UUID) uuid.UUID {
	return AssertClubBodyRespDB(app, assert, resp, SampleClubFactory(userID))
}

func CreateSampleClub(t *testing.T, existingAppAssert *ExistingAppAssert) (eaa ExistingAppAssert, userUUID uuid.UUID, clubUUID uuid.UUID) {
	appAssert, userID := CreateSampleUser(t, existingAppAssert)

	var sampleClubUUID uuid.UUID

	newAppAssert := TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/clubs/",
		Body:   SampleClubFactory(userID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusCreated,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				sampleClubUUID = AssertSampleClubBodyRespDB(app, assert, resp, userID)
			},
		},
	)

	if existingAppAssert == nil {
		return newAppAssert, userID, sampleClubUUID
	} else {
		return *existingAppAssert, userID, sampleClubUUID
	}
}

func TestCreateClubWorks(t *testing.T) {
	existingAppAssert, _, _ := CreateSampleClub(t, nil)
	existingAppAssert.Close()
}

func TestGetClubsWorks(t *testing.T) {
	TestRequest{
		Method: fiber.MethodGet,
		Path:   "/api/v1/clubs/",
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respClubs []models.Club

				err := json.NewDecoder(resp.Body).Decode(&respClubs)

				assert.NilError(err)

				assert.Equal(1, len(respClubs))

				respClub := respClubs[0]

				var dbClubs []models.Club

				err = app.Conn.Order("created_at desc").Find(&dbClubs).Error

				assert.NilError(err)

				assert.Equal(1, len(dbClubs))

				dbClub := dbClubs[0]

				assert.Equal(dbClub.ID, respClub.ID)
				assert.Equal(dbClub.Name, respClub.Name)
				assert.Equal(dbClub.Preview, respClub.Preview)
				assert.Equal(dbClub.Description, respClub.Description)
				assert.Equal(dbClub.NumMembers, respClub.NumMembers)
				assert.Equal(dbClub.IsRecruiting, respClub.IsRecruiting)
				assert.Equal(dbClub.RecruitmentCycle, respClub.RecruitmentCycle)
				assert.Equal(dbClub.RecruitmentType, respClub.RecruitmentType)
				assert.Equal(dbClub.ApplicationLink, respClub.ApplicationLink)
				assert.Equal(dbClub.Logo, respClub.Logo)

				assert.Equal("SAC", dbClub.Name)
				assert.Equal("SAC", dbClub.Preview)
				assert.Equal("SAC", dbClub.Description)
				assert.Equal(1, dbClub.NumMembers)
				assert.Equal(true, dbClub.IsRecruiting)
				assert.Equal(models.RecruitmentCycle(models.Always), dbClub.RecruitmentCycle)
				assert.Equal(models.RecruitmentType(models.Application), dbClub.RecruitmentType)
				assert.Equal("https://generatenu.com/apply", dbClub.ApplicationLink)
				assert.Equal("https://aws.amazon.com/s3", dbClub.Logo)
			},
		},
	).Close()
}

func AssertNumClubsRemainsAtN(app TestApp, assert *assert.A, resp *http.Response, n int) {
	var dbClubs []models.Club

	err := app.Conn.Order("created_at desc").Find(&dbClubs).Error

	assert.NilError(err)

	assert.Equal(n, len(dbClubs))
}

var TestNumClubsRemainsAt1 = func(app TestApp, assert *assert.A, resp *http.Response) {
	AssertNumClubsRemainsAtN(app, assert, resp, 1)
}

func AssertCreateBadClubDataFails(t *testing.T, jsonKey string, badValues []interface{}) {
	appAssert, uuid := CreateSampleUser(t, nil)

	for _, badValue := range badValues {
		sampleClubPermutation := *SampleClubFactory(uuid)
		sampleClubPermutation[jsonKey] = badValue

		TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/clubs/",
			Body:   &sampleClubPermutation,
		}.TestOnErrorAndDB(t, &appAssert,
			ErrorWithDBTester{
				Error:    errors.FailedToValidateClub,
				DBTester: TestNumClubsRemainsAt1,
			},
		)
	}
	appAssert.Close()
}

func TestCreateClubFailsOnInvalidDescription(t *testing.T) {
	AssertCreateBadClubDataFails(t,
		"description",
		[]interface{}{
			"Not an URL",
			"@#139081#$Ad_Axf",
			// "https://google.com", <-- TODO fix once we handle mongo urls
		},
	)
}

func TestCreateClubFailsOnInvalidRecruitmentCycle(t *testing.T) {
	AssertCreateBadClubDataFails(t,
		"recruitment_cycle",
		[]interface{}{
			"1234",
			"garbanzo",
			"@#139081#$Ad_Axf",
			"https://google.com",
		},
	)
}

func TestCreateClubFailsOnInvalidRecruitmentType(t *testing.T) {
	AssertCreateBadClubDataFails(t,
		"recruitment_type",
		[]interface{}{
			"1234",
			"garbanzo",
			"@#139081#$Ad_Axf",
			"https://google.com",
		},
	)

}

func TestCreateClubFailsOnInvalidApplicationLink(t *testing.T) {
	AssertCreateBadClubDataFails(t,
		"application_link",
		[]interface{}{
			"Not an URL",
			"@#139081#$Ad_Axf",
		},
	)

}

func TestCreateClubFailsOnInvalidLogo(t *testing.T) {
	AssertCreateBadClubDataFails(t,
		"logo",
		[]interface{}{
			"Not an URL",
			"@#139081#$Ad_Axf",
			//"https://google.com", <-- TODO uncomment once we figure out s3 url validation
		},
	)
}

func TestUpdateClubWorks(t *testing.T) {
	appAssert, userUUID, clubUUID := CreateSampleClub(t, nil)

	updatedClub := SampleClubFactory(userUUID)
	(*updatedClub)["name"] = "Updated Name"
	(*updatedClub)["preview"] = "Updated Preview"

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   fmt.Sprintf("/api/v1/clubs/%s", clubUUID),
		Body:   updatedClub,
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				AssertClubBodyRespDB(app, assert, resp, updatedClub)
			},
		},
	).Close()
}

func TestUpdateClubFailsOnInvalidBody(t *testing.T) {
	appAssert, userUUID, clubUUID := CreateSampleClub(t, nil)

	body := SampleClubFactory(userUUID)

	for _, invalidData := range []map[string]interface{}{
		{"description": "Not a URL"},
		{"recruitment_cycle": "1234"},
		{"recruitment_type": "ALLLLWAYSSSS"},
		{"application_link": "Not an URL"},
		{"logo": "@12394X_2"},
	} {
		TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/clubs/%s", clubUUID),
			Body:   &invalidData,
		}.TestOnErrorAndDB(t, &appAssert,
			ErrorWithDBTester{
				Error: errors.FailedToValidateClub,
				DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
					var dbClubs []models.Club

					err := app.Conn.Order("created_at desc").Find(&dbClubs).Error

					assert.NilError(err)

					assert.Equal(2, len(dbClubs))

					dbClub := dbClubs[0]

					var dbAdmins []models.User

					err = app.Conn.Model(&dbClub).Association("Admin").Find(&dbAdmins)

					assert.NilError(err)

					assert.Equal(1, len(dbAdmins))

					assert.Equal((*body)["user_id"].(uuid.UUID), dbAdmins[0].ID)
					assert.Equal((*body)["name"].(string), dbClub.Name)
					assert.Equal((*body)["preview"].(string), dbClub.Preview)
					assert.Equal((*body)["description"].(string), dbClub.Description)
					assert.Equal((*body)["num_members"].(int), dbClub.NumMembers)
					assert.Equal((*body)["is_recruiting"].(bool), dbClub.IsRecruiting)
					assert.Equal(models.RecruitmentCycle((*body)["recruitment_cycle"].(string)), dbClub.RecruitmentCycle)
					assert.Equal(models.RecruitmentType((*body)["recruitment_type"].(string)), dbClub.RecruitmentType)
					assert.Equal((*body)["application_link"].(string), dbClub.ApplicationLink)
					assert.Equal((*body)["logo"].(string), dbClub.Logo)
				},
			},
		)
	}
	appAssert.Close()
}

func TestUpdateClubFailsBadRequest(t *testing.T) {
	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/clubs/%s", badRequest),
			Body:   SampleUserFactory(),
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

func TestUpdateClubFailsOnClubIdNotExist(t *testing.T) {
	appAssert, userUUID := CreateSampleUser(t, nil)

	uuid := uuid.New()

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   fmt.Sprintf("/api/v1/clubs/%s", uuid),
		Body:   SampleClubFactory(userUUID),
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

func TestDeleteClubWorks(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(t, nil)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/%s", clubUUID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusNoContent,
			DBTester: TestNumClubsRemainsAt1,
		},
	).Close()
}

func TestDeleteClubNotExist(t *testing.T) {
	uuid := uuid.New()
	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/%s", uuid),
	}.TestOnErrorAndDB(t, nil,
		ErrorWithDBTester{
			Error: errors.ClubNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var club models.Club

				err := app.Conn.Where("id = ?", uuid).First(&club).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))

				AssertNumClubsRemainsAtN(app, assert, resp, 1)
			},
		},
	).Close()
}

func TestDeleteClubBadRequest(t *testing.T) {
	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"hello",
		"null",
	}

	for _, badRequest := range badRequests {
		TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/clubs/%s", badRequest),
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}

}

// contact tests

/*
C- test contact creation works with valid data
test contact creation fails with invalid data
- type
  - missing
  - not a valid type

R- test get contacts works

test delete contact works
*/

func SampleContactFactory() *map[string]interface{} {
	return &map[string]interface{}{
		"type":    "email",
		"content": "jermaine@gmail.com",
	}
}

func ManyContactsFactory() *map[string]map[string]interface{} {
	arr := make(map[string]map[string]interface{})

	arr["email"] = map[string]interface{}{
		"type":    "email",
		"content": "cheeseClub@gmail.com",
	}

	arr["youtube"] = map[string]interface{}{
		"type":    "youtube",
		"content": "youtube.com/cheeseClub",
	}

	arr["facebook"] = map[string]interface{}{
		"type":    "facebook",
		"content": "facebook.com/cheeseClub",
	}

	arr["discord"] = map[string]interface{}{
		"type":    "discord",
		"content": "discord.com/cheeseClub",
	}

	arr["instagram"] = map[string]interface{}{
		"type":    "instagram",
		"content": "instagram.com/cheeseClub",
	}
	arr["github"]= map[string]interface{}{
		"type":    "github",
		"content": "github.com/cheeseClub",
	}

	return &arr
}

func AssertContactBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respContact models.Contact

	// decode the response body into a respContact
	err := json.NewDecoder(resp.Body).Decode(&respContact)

	assert.NilError(err)

	var dbContacts []models.Contact

	// get all contacts from the database ordered by created_at and store them in dbContacts
	// err = app.Conn.Order("created_at desc").Find(&dbContacts).Error

	assert.NilError(err)

	// assert.Equal(1, len(dbContacts))

	dbContact := dbContacts[0]

	assert.Equal(dbContact.ID, respContact.ID)
	assert.Equal(dbContact.Type, respContact.Type)
	assert.Equal(dbContact.Content, respContact.Content)

	return dbContact.ID
}

func AssertSampleContactBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, clubUUID uuid.UUID) uuid.UUID {
	return AssertContactBodyRespDB(app, assert, resp, SampleContactFactory())
}

func CreateSampleContact(t *testing.T, existingAppAssert *ExistingAppAssert) (eaa ExistingAppAssert, clubUUID uuid.UUID, contactUUID uuid.UUID) {
	appAssert, _, clubUUID := CreateSampleClub(t, nil)

	for _, contact := range *ManyContactsFactory() {

		appAssert = TestRequest{
			Method: fiber.MethodPut,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/contacts", clubUUID),
			Body:   &contact,
		}.TestOnStatusAndDB(t, &appAssert,
			DBTesterWithStatus{
				Status: fiber.StatusOK,
				DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
					AssertContactBodyRespDB(app, assert, resp, &contact)
				},
			},
		)
	}


	var sampleContactUUID uuid.UUID

	newAppAssert := TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/contacts", clubUUID),
		Body:   SampleContactFactory(),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				sampleContactUUID = AssertSampleContactBodyRespDB(app, assert, resp, clubUUID)
			},
		},
	)

	if existingAppAssert == nil {
		return newAppAssert, clubUUID, sampleContactUUID
	} else {
		return *existingAppAssert, clubUUID, sampleContactUUID
	}
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

var TestNumContactsRemainsAt0 = func(app TestApp, assert *assert.A, resp *http.Response) {
	AssertNumContactsRemainsAtN(app, assert, resp, 0)
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
				DBTester: TestNumContactsRemainsAt0,
			},
		)
	}
	appAssert.Close()
}

func TestCreateContactFailsOnInvalidType(t *testing.T) {
	AssertCreateBadContactDataFails(t,
		"type",
		[]interface{}{
			"Not a valid type",
			"@#139081#$Ad_Axf",
		},
	)
}

func TestCreateContactFailsOnInvalidContent(t *testing.T) {
	AssertCreateBadContactDataFails(t,
		"content",
		[]interface{}{
			"Not a valid url",
			"@#139081#$Ad_Axf",
		},
	)
}

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
// func TestPutContactUpdatesExistingContact(t *testing.T){
// 	appAssert, clubUUID, contactUUID := CreateSampleContact(t, nil)

// 	updatedContact := SampleContactFactory()
// 	(*updatedContact)["content"] = ""

// }

// func TestGetContactByIdWorks(t *testing.T) {
// 	appAssert, clubUUID, contactUUID := CreateSampleContact(t, nil)

// 	TestRequest{
// 		Method: fiber.MethodGet,
// 		Path:   fmt.Sprintf("/api/v1/clubs/%s/contacts", clubUUID),
// 	}.TestOnStatusAndDB(t, &appAssert,
// 		DBTesterWithStatus{
// 			Status: fiber.StatusOK,
// 			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
// 				var respContacts []models.Contact

// 				err := json.NewDecoder(resp.Body).Decode(&respContacts)

// 				assert.NilError(err)

// 				assert.Equal(1, len(respContacts))

// 				respContact := respContacts[0]

// 				var dbContacts []models.Contact

// 				err = app.Conn.Order("created_at desc").Find(&dbContacts).Error

// 				assert.NilError(err)

// 				assert.Equal(1, len(dbContacts))

// 				dbContact := dbContacts[0]

// 				assert.Equal(dbContact.ID, respContact.ID)
// 				assert.Equal(dbContact.Type, respContact.Type)
// 				assert.Equal(dbContact.Content, respContact.Content)
// 			},
// 		},
// 	).Close()
// }

// func TestGetContactsByClubIDWorks(t *testing.T ){
// 	appAssert, club1, contact1 := CreateSampleContact(t, nil)
	
// 	appAssert, club

// 	appAssert = appTestRequest{
// 		Method: fiber.MethodGet,
// 		Path: fmt.Sprintf("/api/v1/clubs/%s/contacts/%s", clubUUID, contactUUID),
// 	}.TestOnStatusAndDB(t, &appAssert,
// 		DBTesterWithStatus{
// 			Status: fiber.StatusOK,
// 			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
// 				var respContact models.Contact

// 				err := json.NewDecoder(resp.Body).Decode(&respContact)

// 				assert.NilError(err)

// 				var dbContact models.Contact

// 				err = app.Conn.Order("created_at desc").First(&dbContact).Error

// 				assert.NilError(err)

// 				assert.Equal(dbContact.ID, respContact.ID)
// 				assert.Equal(dbContact.Type, respContact.Type)
// 				assert.Equal(dbContact.Content, respContact.Content)
// 			},
// 		},
// 	)
// }
