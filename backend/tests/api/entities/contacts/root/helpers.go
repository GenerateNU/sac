package root

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GenerateNU/sac/backend/entities/models"
	clubsRoot "github.com/GenerateNU/sac/backend/tests/api/entities/clubs/root"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SampleContactFactory() *map[string]interface{} {
	return &map[string]interface{}{
		"type":    "email",
		"content": "jermaine@gmail.com",
	}
}

func ManyContactsFactory() map[string](*map[string]interface{}) {
	arr := make(map[string]*map[string]interface{})

	arr["facebook"] = &map[string]interface{}{
		"type":    "facebook",
		"content": "https://facebook.com/cheeseClub",
	}

	arr["instagram"] = &map[string]interface{}{
		"type":    "instagram",
		"content": "https://instagram.com/cheeseClub",
	}

	arr["x"] = &map[string]interface{}{
		"type":    "x",
		"content": "https://x.com/cheeseClub",
	}

	arr["linkedin"] = &map[string]interface{}{
		"type":    "linkedin",
		"content": "https://linkedin.com/cheeseClub",
	}

	arr["youtube"] = &map[string]interface{}{
		"type":    "youtube",
		"content": "https://youtube.com/cheeseClub",
	}

	arr["github"] = &map[string]interface{}{
		"type":    "github",
		"content": "https://github.com/GeneratNU/sac",
	}

	arr["slack"] = &map[string]interface{}{
		"type":    "slack",
		"content": "https://join.slack.com/cheeseClub",
	}

	arr["discord"] = &map[string]interface{}{
		"type":    "discord",
		"content": "https://discord.gg/cheeseClub",
	}

	arr["email"] = &map[string]interface{}{
		"type":    "email",
		"content": "cheeseClub@gmail.com",
	}

	arr["customSite"] = &map[string]interface{}{
		"type":    "customSite",
		"content": "https://cheeseClub.com",
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
	appAssert, _, clubUUID := clubsRoot.CreateSampleClub(existingAppAssert)

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
	existingAppAssert, _, clubUUID = clubsRoot.CreateSampleClub(existingAppAssert)

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

func AssertNumContactsRemainsAtN(eaa h.ExistingAppAssert, resp *http.Response, n int) {
	var dbContacts []models.Contact

	err := eaa.App.Conn.Order("created_at desc").Find(&dbContacts).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(n, len(dbContacts))
}
