package tests

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

func SampleSeriesFactory() *map[string]interface{} {
	return CustomSampleSeriesFactory(
		models.CreateSeriesRequestBody{
			RecurringType:   "daily",
			MaxOccurrences:  10,
			SeparationCount: 4,
			DayOfWeek:       3,
			WeekOfMonth:     2,
			DayOfMonth:      1,
		},
	)
}

func CustomSampleSeriesFactory(series models.CreateSeriesRequestBody) *map[string]interface{} {
	return &map[string]interface{}{
		"name":         "Software Development",
		"preview":      "CS4500 at northeastern",
		"content":      "Software development with ben lerner",
		"start_time":   "2024-03-20T18:00:00Z",
		"end_time":     "2024-03-20T21:00:00Z",
		"location":     "ISEC",
		"event_type":   "membersOnly",
		"is_recurring": true,
		"series":       series,
	}
}

func CompareEventSlices(eaa h.ExistingAppAssert, dbEvents, respEvents []models.Event) {
	for i, respEvent := range respEvents {
		dbEvent := dbEvents[i]
		CompareEvents(eaa, dbEvent, respEvent)
	}
}

func CompareEvents(eaa h.ExistingAppAssert, dbEvent, respEvent models.Event) {
	eaa.Assert.Equal(dbEvent.ID, respEvent.ID)
	eaa.Assert.Equal(dbEvent.Name, respEvent.Name)
	eaa.Assert.Equal(dbEvent.Preview, respEvent.Preview)
	eaa.Assert.Equal(dbEvent.Content, respEvent.Content)
	eaa.Assert.Equal(dbEvent.StartTime.Compare(respEvent.StartTime), 0)
	eaa.Assert.Equal(dbEvent.EndTime.Compare(respEvent.EndTime), 0)
	eaa.Assert.Equal(dbEvent.Location, respEvent.Location)
	eaa.Assert.Equal(dbEvent.IsRecurring, respEvent.IsRecurring)
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
	respEvents, dbEvents := GetRespAndDBEvents(eaa, resp)

	var uuidList []uuid.UUID

	for i, respEvent := range respEvents {
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

func CreateSampleEvent(existingAppAssert h.ExistingAppAssert, factoryFunction EventFactory) (h.ExistingAppAssert, []uuid.UUID) {
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

func AssertNumEventsRemainsAtN(eaa h.ExistingAppAssert, resp *http.Response, n int) {
	var dbEvents []models.Event

	err := eaa.App.Conn.Order("created_at desc").Find(&dbEvents).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(n, len(dbEvents))
}

func AssertNumSeriesRemainsAtN(eaa h.ExistingAppAssert, resp *http.Response, n int) {
	var dbSeries []models.Series

	err := eaa.App.Conn.Order("created_at desc").Find(&dbSeries).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(n, len(dbSeries))
}

func TestCreateEventWorks(t *testing.T) {
	existingAppAssert, _ := CreateSampleEvent(h.InitTest(t), SampleEventFactory)
	existingAppAssert.Close()
}

func TestCreateEventSeriesWorks(t *testing.T) {
	existingAppAssert, _ := CreateSampleEvent(h.InitTest(t), SampleSeriesFactory)
	existingAppAssert.Close()
}

func TestGetEventWorks(t *testing.T) {
	existingAppAssert, eventUUID := CreateSampleEvent(h.InitTest(t), SampleEventFactory)

	existingAppAssert.TestOnStatusAndTester(h.TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/events/%s", eventUUID[0]),
		Role:   &models.Super,
	},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertEventListBodyRespDB(eaa, resp, SampleEventFactory())
			},
		},
	).Close()
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

func TestGetSeriesByEventIDWorks(t *testing.T) {
	existingAppAssert, eventUUIDs := CreateSampleEvent(h.InitTest(t), SampleSeriesFactory)

	existingAppAssert.TestOnStatusAndTester(h.TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/events/%s/series", eventUUIDs[2]),
		Role:   &models.Super,
	},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				respEvents, dbEvents := GetRespAndDBEvents(eaa, resp)
				CompareEventSlices(eaa, dbEvents, respEvents)
			},
		},
	).Close()
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

func AssertCreateBadEventSeriesDataFails(t *testing.T, badSeries models.CreateSeriesRequestBody, expectedErr errors.Error) {
	appAssert, _, _ := CreateSampleStudent(t, nil)

	sampleSeriesPermutation := CustomSampleSeriesFactory(badSeries)

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/events/",
			Body:   sampleSeriesPermutation,
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error: expectedErr,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertNumEventsRemainsAtN(eaa, resp, 0)
			},
		},
	)
	appAssert.Close()
}

func TestCreateSeriesFailsOnInvalidRecurringType(t *testing.T) {
	AssertCreateBadEventSeriesDataFails(t,
		models.CreateSeriesRequestBody{
			RecurringType:   "annually",
			MaxOccurrences:  10,
			SeparationCount: 0,
			DayOfWeek:       3,
			WeekOfMonth:     2,
			DayOfMonth:      1,
		},
		errors.FailedToValidateEventSeries,
	)
}

func TestCreateSeriesFailsOnInvalidMaxOccurrences(t *testing.T) {
	AssertCreateBadEventSeriesDataFails(t,
		models.CreateSeriesRequestBody{
			RecurringType:   "weekly",
			MaxOccurrences:  -1,
			SeparationCount: 0,
			DayOfWeek:       3,
			WeekOfMonth:     2,
			DayOfMonth:      1,
		},
		errors.FailedToValidateEventSeries,
	)
}

func TestCreateSeriesFailsOnInvalidSeparationCount(t *testing.T) {
	AssertCreateBadEventSeriesDataFails(t,
		models.CreateSeriesRequestBody{
			RecurringType:   "weekly",
			MaxOccurrences:  10,
			SeparationCount: -1,
			DayOfWeek:       3,
			WeekOfMonth:     2,
			DayOfMonth:      1,
		},
		errors.FailedToValidateEventSeries,
	)
}

func TestCreateSeriesFailsOnInvalidDayOfWeek(t *testing.T) {
	AssertCreateBadEventSeriesDataFails(t,
		models.CreateSeriesRequestBody{
			RecurringType:   "weekly",
			MaxOccurrences:  10,
			SeparationCount: 0,
			DayOfWeek:       8,
			WeekOfMonth:     2,
			DayOfMonth:      1,
		},
		errors.FailedToValidateEventSeries,
	)
}

func TestCreateSeriesFailsOnInvalidWeekOfMonth(t *testing.T) {
	AssertCreateBadEventSeriesDataFails(t,
		models.CreateSeriesRequestBody{
			RecurringType:   "weekly",
			MaxOccurrences:  10,
			SeparationCount: 0,
			DayOfWeek:       5,
			WeekOfMonth:     -5,
			DayOfMonth:      1,
		},
		errors.FailedToValidateEventSeries,
	)
}

func TestCreateSeriesFailsOnInvalidDayOfMonth(t *testing.T) {
	AssertCreateBadEventSeriesDataFails(t,
		models.CreateSeriesRequestBody{
			RecurringType:   "weekly",
			MaxOccurrences:  10,
			SeparationCount: 0,
			DayOfWeek:       5,
			WeekOfMonth:     2,
			DayOfMonth:      42,
		},
		errors.FailedToValidateEventSeries,
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
				AssertEventListBodyRespDB(eaa, resp, updatedEvent)
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
				Path:   fmt.Sprintf("/api/v1/events/%s", badRequest),
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

func AssertDeleteWorks(t *testing.T, factoryFunction EventFactory, requestPath string, tester h.Tester) {
	appAssert, eventUUIDs := CreateSampleEvent(h.InitTest(t), factoryFunction)

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf(requestPath, eventUUIDs[0]),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: tester,
		},
	).Close()
}

func TestDeleteEventWorks(t *testing.T) {
	AssertDeleteWorks(t, SampleEventFactory, "/api/v1/events/%s", func(eaa h.ExistingAppAssert, resp *http.Response) {
		AssertNumEventsRemainsAtN(eaa, resp, 0)
	})
}

func TestDeleteSeriesByEventIDWorks(t *testing.T) {
	AssertDeleteWorks(t, SampleEventFactory, "/api/v1/events/%s", func(eaa h.ExistingAppAssert, resp *http.Response) {
		AssertNumEventsRemainsAtN(eaa, resp, 0)
		AssertNumSeriesRemainsAtN(eaa, resp, 0)
	})
}

func AssertDeleteNotExistFails(t *testing.T, factoryFunction EventFactory, requestPath string, tester h.Tester, badUUID uuid.UUID) {
	appAssert, _ := CreateSampleEvent(h.InitTest(t), factoryFunction)

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/events/%s", badUUID),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error:  errors.EventNotFound,
			Tester: tester,
		},
	).Close()
}

func TestDeleteEventNotExist(t *testing.T) {
	uuid := uuid.New()

	AssertDeleteNotExistFails(t, SampleEventFactory, "/api/v1/events/%s", func(eaa h.ExistingAppAssert, resp *http.Response) {
		var event []models.Event
		err := eaa.App.Conn.Where("id = ?", uuid).First(&event).Error
		eaa.Assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))

		AssertNumEventsRemainsAtN(eaa, resp, 1)
	}, uuid)
}

func TestDeleteSeriesNotExist(t *testing.T) {
	uuid := uuid.New()

	AssertDeleteNotExistFails(t, SampleSeriesFactory, "/api/v1/events/%s/series", func(eaa h.ExistingAppAssert, resp *http.Response) {
		var events []models.Event
		err := eaa.App.Conn.Where("id = ?", uuid).First(&events).Error
		eaa.Assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))

		AssertNumSeriesRemainsAtN(eaa, resp, 1)
	}, uuid)
}

func AssertDeleteBadRequestFails(t *testing.T, requestPath string) {
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
				Path:   fmt.Sprintf(requestPath, badRequest),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}

	appAssert.Close()
}

func TestDeleteEventBadRequest(t *testing.T) {
	AssertDeleteBadRequestFails(t, "/api/v1/events/%s")
}

func TestDeleteSeriesBadRequest(t *testing.T) {
	AssertDeleteBadRequestFails(t, "/api/v1/events/%s/series")
}
