package recruitment

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/gofiber/fiber/v2"
)

type ClubRecruitmentController struct {
	clubRecruitmentService ClubRecruitmentServicer
}

func NewClubRecruitmentController(clubRecruitmentService ClubRecruitmentServicer) *ClubRecruitmentController {
	return &ClubRecruitmentController{clubRecruitmentService: clubRecruitmentService}
}

func (cr *ClubRecruitmentController) CreateClubRecruitment(c *fiber.Ctx) error {
	var body CreateClubRecruitmentRequestBody
	if err := c.BodyParser(&body); err != nil {
		return utilities.InvalidJSON()
	}

	recruitment, err := cr.clubRecruitmentService.CreateClubRecruitment(c.Context(), c.Params("clubID"), body)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(recruitment)
}

func (cr *ClubRecruitmentController) GetClubRecruitment(c *fiber.Ctx) error {
	recruitment, err := cr.clubRecruitmentService.GetClubRecruitment(c.Context(), c.Params("clubID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(recruitment)
}

func (cr *ClubRecruitmentController) UpdateClubRecruitment(c *fiber.Ctx) error {
	var body UpdateClubRecruitmentRequestBody
	if err := c.BodyParser(&body); err != nil {
		return utilities.InvalidJSON()
	}

	recruitment, err := cr.clubRecruitmentService.UpdateClubRecruitment(c.Context(), c.Params("clubID"), body)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(recruitment)
}

func (cr *ClubRecruitmentController) DeleteClubRecruitment(c *fiber.Ctx) error {
	if err := cr.clubRecruitmentService.DeleteClubRecruitment(c.Context(), c.Params("clubID")); err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

func (cr *ClubRecruitmentController) CreateClubRecruitmentApplication(c *fiber.Ctx) error {
	var body CreateApplicationRequestBody
	if err := c.BodyParser(&body); err != nil {
		return utilities.InvalidJSON()
	}

	application, err := cr.clubRecruitmentService.CreateClubRecruitmentApplication(c.Context(), c.Params("clubID"), body)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(application)
}

func (cr *ClubRecruitmentController) GetClubRecruitmentApplications(c *fiber.Ctx) error {
	pageInfo, ok := fiberpaginate.FromContext(c)
	if !ok {
		return utilities.ErrExpectedPageInfo
	}

	applications, err := cr.clubRecruitmentService.GetClubRecruitmentApplications(c.Context(), c.Params("clubID"), *pageInfo)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(applications)
}

func (cr *ClubRecruitmentController) GetClubRecruitmentApplication(c *fiber.Ctx) error {
	application, err := cr.clubRecruitmentService.GetClubRecruitmentApplication(c.Context(), c.Params("clubID"), c.Params("applicationID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(application)
}

func (cr *ClubRecruitmentController) UpdateClubRecruitmentApplication(c *fiber.Ctx) error {
	var body UpdateApplicationRequestBody
	if err := c.BodyParser(&body); err != nil {
		return utilities.InvalidJSON()
	}

	application, err := cr.clubRecruitmentService.UpdateClubRecruitmentApplication(c.Context(), c.Params("clubID"), c.Params("applicationID"), body)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(application)
}

func (cr *ClubRecruitmentController) DeleteClubRecruitmentApplication(c *fiber.Ctx) error {
	if err := cr.clubRecruitmentService.DeleteClubRecruitmentApplication(c.Context(), c.Params("clubID"), c.Params("applicationID")); err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
