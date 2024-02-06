package controllers

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type EventController struct {
	eventService services.EventServiceInterface
}

func NewEventController(eventService services.EventServiceInterface) *EventController {
	return &EventController{eventService: eventService}
}

func (l *EventController) GetAllEvents(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	events, err := l.eventService.GetEvents(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func (l *EventController) CreateEvent(c *fiber.Ctx) error {
	recurringPattern, err := getRecurringPattern(c)
	
	if err != nil {
		return CreateEventSeries(l, c, recurringPattern)
	}

	var eventBody models.CreateEventRequestBody
	if err := c.BodyParser(&eventBody); err != nil {
		return errors.FailedToCreateEvent.FiberError(c)
	}

	event, err := l.eventService.CreateEvent(eventBody)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(event)
}

func CreateEventSeries(l *EventController, c *fiber.Ctx, recurringPattern models.CreateRecurringPatternRequestBody) error {
	var eventBody []models.CreateEventRequestBody
	if err := c.BodyParser(&eventBody); err != nil {
		return errors.FailedToCreateEvent.FiberError(c)
	}

	event, err := l.eventService.CreateEventSeries(eventBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(event)
}

func getRecurringPattern(c *fiber.Ctx) (models.CreateRecurringPatternRequestBody, error) {
	recurringType := c.Query("recurring_type")
	separationCount, separationCountErr := strconv.Atoi(c.Query("separation_count"))
	maxOccurrences, maxOccurrencesErr := strconv.Atoi(c.Query("max_occurrences"))
	dayOfWeek, dayOfWeekErr := strconv.Atoi(c.Query("day_of_week"))
	weekOfMonth, weekOfMonthErr := strconv.Atoi(c.Query("week_of_month"))
	dayOfMonth, dayOfMonthErr := strconv.Atoi(c.Query("day_of_month"))

	recurringPattern := models.CreateRecurringPatternRequestBody{
		RecurringType:   models.RecurringType(recurringType),
		SeparationCount: separationCount,
		MaxOccurrences:  maxOccurrences,
		DayOfWeek:       dayOfWeek,
		WeekOfMonth:     weekOfMonth,
		DayOfMonth:      dayOfMonth,
	}

	if (recurringType == "" || separationCountErr != nil || maxOccurrencesErr != nil || dayOfWeekErr != nil || weekOfMonthErr != nil || dayOfMonthErr != nil) {
		return recurringPattern, errors.FailedToValidateEventSeries.FiberError(c)
	}

	return recurringPattern, nil
}

func (l *EventController) GetEvent(c *fiber.Ctx) error {
	event, err := l.eventService.GetEvent(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(event)
}

func (l *EventController) UpdateEvent(c *fiber.Ctx) error {
	var eventBody models.UpdateEventRequestBody

	if err := c.BodyParser(&eventBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	updatedEvent, err := l.eventService.UpdateEvent(c.Params("id"), eventBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(updatedEvent)
}

func (l *EventController) DeleteEvent(c *fiber.Ctx) error {
	err := l.eventService.DeleteEvent(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
