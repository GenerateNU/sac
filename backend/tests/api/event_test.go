package tests

/* TEST CASES:
test createEventSeries fails on:
	the stuff from above
	invalid recurringType (ex. Annually)
	invalid separationCount

test updateEvent fails on:

*/
import (
	stdliberrors "errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventFactory func() *map[string]interface{}

func SampleEventFactory() *map[string]interface{} {
	return &map[string]interface{}{
		"name":         "Generate",
		"preview":      "Generate is Northeastern's premier student-led product development studio.",
		"content":      "Come join us for Generate's end-of-semester showcase",
		"start_time":   "2023-12-20T18:00:00Z",
		"end_time":     "2023-12-20T21:00:00Z",
		"location":     "Carter Fields",
		"event_type":   "open",
		"is_recurring": false,
	}
}

func SampleEventSeriesFactory() *map[string]interface{} {
	return &map[string]interface{}{
		"name":         "Software Development",
		"preview":      "CS4500 at northeastern",
		"content":      "Software development with ben lerner",
		"start_time":   "2023-12-20T18:00:00Z",
		"end_time":     "2023-12-20T21:00:00Z",
		"location":     "ISEC",
		"event_type":   "membersOnly",
		"is_recurring": true,
		"series": models.CreateSeriesRequestBody{
			RecurringType:   "weekly",
			MaxOccurrences:  10,
			SeparationCount: 0,
			DayOfWeek:       3,
			WeekOfMonth:     2,
			DayOfMonth:      1,
		},
	}
}

func CompareEvents(eaa h.ExistingAppAssert, event1, event2 models.Event) {
	eaa.Assert.Equal(event1.ID, event2.ID)
	eaa.Assert.Equal(event1.Name, event2.Name)
	eaa.Assert.Equal(event1.Preview, event2.Preview)
	eaa.Assert.Equal(event1.Content, event2.Content)
	eaa.Assert.Equal(event1.StartTime.Compare(event2.StartTime), 0)
	eaa.Assert.Equal(event1.EndTime.Compare(event2.EndTime), 0)
	eaa.Assert.Equal(event1.Location, event2.Location)
	eaa.Assert.Equal(event1.IsRecurring, event2.IsRecurring)
}

func GetRespAndDBEvents(eaa h.ExistingAppAssert, resp *http.Response) ([]models.Event, []models.Event) {
	var respEventList []models.Event

	err := json.NewDecoder(resp.Body).Decode(&respEventList) 

	eaa.Assert.NilError(err)

	var dbEvents []models.Event

	err = eaa.App.Conn.Order("created_at desc").Find(&dbEvents).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(len(respEventList), len(dbEvents))

	return respEventList, dbEvents
}

func AssertEventListBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) []uuid.UUID {
	respEventList, dbEvents := GetRespAndDBEvents(eaa, resp)

	var uuidList []uuid.UUID

	for i, respEvent := range respEventList {
		dbEvent := dbEvents[i]

		CompareEvents(eaa, dbEvent, respEvent)

		uuidList = append(uuidList, dbEvent.ID)
	}

	return uuidList
}

func AssertEventBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) uuid.UUID {	
	var respEvent models.Event

	err := json.NewDecoder(resp.Body).Decode(&respEvent)

	eaa.Assert.NilError(err)

	var dbEvents []models.Event

	err = eaa.App.Conn.Order("created_at desc").Find(&dbEvents).Error

	eaa.Assert.NilError(err)

	dbEvent := dbEvents[0]

	CompareEvents(eaa, dbEvent, respEvent)

	bodyStartTime, err := time.Parse(time.RFC3339, (*body)["start_time"].(string))
	eaa.Assert.NilError(err)
	bodyEndTime, err := time.Parse(time.RFC3339, (*body)["end_time"].(string))
	eaa.Assert.NilError(err)

	eaa.Assert.Equal((*body)["name"].(string), dbEvent.Name)
	eaa.Assert.Equal((*body)["preview"].(string), dbEvent.Preview)
	eaa.Assert.Equal((*body)["content"].(string), dbEvent.Content)
	eaa.Assert.Equal(bodyStartTime.Compare(dbEvent.StartTime), 0)
	eaa.Assert.Equal(bodyEndTime.Compare(dbEvent.EndTime), 0)
	eaa.Assert.Equal((*body)["location"].(string), dbEvent.Location)
	eaa.Assert.Equal(models.EventType((*body)["event_type"].(string)), dbEvent.EventType)
	eaa.Assert.Equal((*body)["is_recurring"].(bool), dbEvent.IsRecurring)

	return dbEvent.ID
}

func AssertSampleEventBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response) []uuid.UUID {
	sampleEvent := SampleEventFactory()
	return AssertEventListBodyRespDB(eaa, resp, sampleEvent)
}

func CreateSampleEvent(existingAppAssert h.ExistingAppAssert, factoryFunction EventFactory) (eaa h.ExistingAppAssert, eventUUID []uuid.UUID) {
	var sampleEventUUIDs []uuid.UUID

	newAppAssert := existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/events/",
			Body:   factoryFunction(),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				sampleEventUUIDs = AssertSampleEventBodyRespDB(eaa, resp)
			},
		},
	)

	return newAppAssert, sampleEventUUIDs
}

func TestCreateEventWorks(t *testing.T) {
	existingAppAssert, _ := CreateSampleEvent(h.InitTest(t), SampleEventFactory)
	existingAppAssert.Close()
}

func TestCreateEventSeriesWorks(t *testing.T) {
	existingAppAssert, _ := CreateSampleEvent(h.InitTest(t), SampleEventSeriesFactory)
	existingAppAssert.Close()
}

func TestGetEventsWorks(t *testing.T) {
	existingAppAssert, _ := CreateSampleEvent(h.InitTest(t), SampleEventFactory)

	existingAppAssert.TestOnStatusAndTester(h.TestRequest{
		Method: fiber.MethodGet,
		Path:   "/api/v1/events/",
		Role:   &models.Super,
	},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				respEvents, dbEvents := GetRespAndDBEvents(eaa, resp)

				respEvent := respEvents[0]

				dbEvent := dbEvents[0]

				CompareEvents(eaa, dbEvent, respEvent)
			},
		},
	).Close()
}

func AssertNumEventsRemainsAtN(eaa h.ExistingAppAssert, resp *http.Response, n int) {
	var dbEvents []models.Event

	err := eaa.App.Conn.Order("created_at desc").Find(&dbEvents).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(n, len(dbEvents))
}

func AssertCreateBadEventDataFails(t *testing.T, jsonKey string, badValues []interface{}, expectedErr errors.Error) {
	appAssert, _, _ := CreateSampleStudent(t, nil)

	for _, badValue := range badValues {
		sampleEventPermutation := *SampleEventFactory()
		sampleEventPermutation[jsonKey] = badValue

		appAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodPost,
				Path:   "/api/v1/events/",
				Body:   &sampleEventPermutation,
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error: expectedErr,
				Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
					AssertNumEventsRemainsAtN(eaa, resp, 0)
				},
			},
		)
	}
	appAssert.Close()
}

func TestCreateEventFailsOnInvalidStartTime(t *testing.T) {
	AssertCreateBadEventDataFails(t,
		"start_time",
		[]interface{}{
			"2023-12-20, 18:00",
		},
		errors.FailedToParseRequestBody,
	)
}

func TestCreateEventFailsOnInvalidEndTime(t *testing.T) {
	AssertCreateBadEventDataFails(t,
		"end_time",
		[]interface{}{
			"2023-12-20, 22:00",
		},
		errors.FailedToParseRequestBody,
	)
}

func TestCreateEventFailsOnEndTimeBeforeStartTime(t *testing.T) {
	AssertCreateBadEventDataFails(t,
		"end_time",
		[]interface{}{
			"2023-12-20T17:00:00Z",
		},
		errors.FailedToValidateEvent,
	)
}

func TestCreateEventFailsOnInvalidEventType(t *testing.T) {
	AssertCreateBadEventDataFails(t,
		"event_type",
		[]interface{}{
			"everyone",
			"open membersOnly",
		},
		errors.FailedToValidateEvent,
	)
}

func TestUpdateEventWorks(t *testing.T) {
	appAssert, eventUUID := CreateSampleEvent(h.InitTest(t), SampleEventFactory)

	updatedEvent := SampleEventFactory()
	(*updatedEvent)["name"] = "Updated Name"
	(*updatedEvent)["preview"] = "Updated Preview"

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/events/%s", eventUUID),
			Body:   updatedEvent,
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertEventBodyRespDB(eaa, resp, updatedEvent)
			},
		},
	).Close()
}

func TestUpdateEventFailsOnInvalidBody(t *testing.T) {
	appAssert, eventUUID := CreateSampleEvent(h.InitTest(t), SampleEventFactory)

	body := SampleEventFactory()

	for _, invalidData := range []map[string]interface{}{
		{"start_time": "Not a datetime"},
		{"end_time": "another non-datetime"},
	} {
		invalidData := invalidData
		appAssert = appAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodPatch,
				Path:   fmt.Sprintf("/api/v1/events/%s", eventUUID),
				Body:   &invalidData,
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error: errors.FailedToParseRequestBody,
				Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
					var dbEvents []models.Event

					err := eaa.App.Conn.Order("created_at desc").Find(&dbEvents).Error

					eaa.Assert.NilError(err)

					dbEvent := dbEvents[0]

					bodyStartTime, err := time.Parse(time.RFC3339, (*body)["start_time"].(string))
					eaa.Assert.NilError(err)
					bodyEndTime, err := time.Parse(time.RFC3339, (*body)["end_time"].(string))
					eaa.Assert.NilError(err)

					eaa.Assert.Equal((*body)["name"].(string), dbEvent.Name)
					eaa.Assert.Equal((*body)["preview"].(string), dbEvent.Preview)
					eaa.Assert.Equal((*body)["content"].(string), dbEvent.Content)
					eaa.Assert.Equal(bodyStartTime.Compare(dbEvent.StartTime), 0)
					eaa.Assert.Equal(bodyEndTime.Compare(dbEvent.EndTime), 0)
					eaa.Assert.Equal((*body)["location"].(string), dbEvent.Location)
					eaa.Assert.Equal(models.EventType((*body)["event_type"].(string)), dbEvent.EventType)
					eaa.Assert.Equal((*body)["is_recurring"].(bool), dbEvent.IsRecurring)
				},
			},
		)
	}
	appAssert.Close()
}

func TestUpdateEventFailsBadRequest(t *testing.T) {
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
				Method: fiber.MethodPatch,
				Path:   fmt.Sprintf("/api/v1/clubs/%s", badRequest),
				Body:   SampleEventFactory(),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}

	appAssert.Close()
}

func TestUpdateEventFailsOnEventIdNotExist(t *testing.T) {
	uuid := uuid.New()

	h.InitTest(t).TestOnErrorAndTester(h.TestRequest{
		Method:             fiber.MethodPatch,
		Path:               fmt.Sprintf("/api/v1/events/%s", uuid),
		Body:               SampleEventFactory(),
		Role:               &models.Super,
		TestUserIDReplaces: h.StringToPointer("user_id"),
	},
		h.ErrorWithTester{
			Error: errors.EventNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var event models.Event

				err := eaa.App.Conn.Where("id = ?", uuid).First(&event).Error

				eaa.Assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestDeleteEventWorks(t *testing.T) {
	appAssert, eventUUID := CreateSampleEvent(h.InitTest(t), SampleEventFactory)

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/events/%s", eventUUID),
		},
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertNumEventsRemainsAtN(eaa, resp, 0)
			},
		},
	).Close()
}

func TestDeleteEventNotExist(t *testing.T) {
	uuid := uuid.New()
	h.InitTest(t).TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/events/%s", uuid),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error: errors.EventNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var event []models.Event
				err := eaa.App.Conn.Where("id = ?", uuid).First(&event).Error
				eaa.Assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))

				AssertNumEventsRemainsAtN(eaa, resp, 0)
			},
		},
	).Close()
}

func TestDeleteEventBadRequest(t *testing.T) {
	appAssert := h.InitTest(t)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"hello",
		"null",
	}

	for _, badRequest := range badRequests {
		appAssert = appAssert.TestOnError(
			h.TestRequest{
				Method: fiber.MethodDelete,
				Path:   fmt.Sprintf("/api/v1/events/%s", badRequest),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}

	appAssert.Close()
}
