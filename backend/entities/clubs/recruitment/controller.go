package recruitment

import (
	"net/http"

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
		return err
	}

	club, err := cr.clubRecruitmentService.CreateClubRecruitment(c.UserContext(), c.Params("clubID"), body)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(club)
}

func (cr *ClubRecruitmentController) UpdateClubRecruitment(c *fiber.Ctx) error {
	var body UpdateClubRecruitmentRequestBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	club, err := cr.clubRecruitmentService.UpdateClubRecruitment(c.UserContext(), c.Params("clubID"), body)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(club)
}

func (cr *ClubRecruitmentController) DeleteClubRecruitment(c *fiber.Ctx) error {
	if err := cr.clubRecruitmentService.DeleteClubRecruitment(c.UserContext(), c.Params("clubID")); err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

func (cr *ClubRecruitmentController) GetClubRecruitmentApplications(c *fiber.Ctx) error {
	return nil
}

func (cr *ClubRecruitmentController) CreateClubRecruitmentApplication(c *fiber.Ctx) error {
	return nil
}

func (cr *ClubRecruitmentController) GetClubRecruitmentApplication(c *fiber.Ctx) error {
	return nil
}

func (cr *ClubRecruitmentController) UpdateClubRecruitmentApplication(c *fiber.Ctx) error {
	return nil
}

func (cr *ClubRecruitmentController) DeleteClubRecruitmentApplication(c *fiber.Ctx) error {
	return nil
}
