package base

import (
	"log"
	"net/http"
	"os"

	search_types "github.com/GenerateNU/sac/backend/search/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

type SearchController struct {
	searchService SearchServiceInterface
}

func NewSearchController(searchService SearchServiceInterface) *SearchController {
	return &SearchController{searchService: searchService}
}

// SearchClubs godoc
//
// @Summary		Searches through clubs
// @Description	Searches through clubs
// @ID			search-clubs
// @Tags      	search
// @Accept		json
// @Produce		json
// @Param		searchQuery	query	search_types.ClubSearchRequest	true	"Search Body"
// @Success		200	  {object}	    []search_types.SearchResult[models.Club]
// @Failure     404   {object}		error
// @Failure     500   {object}      error
// @Router		/search/clubs  [get]
func (s *SearchController) SearchClubs(c *fiber.Ctx) error {
	log.SetOutput(os.Stdout)

	var searchQuery search_types.ClubSearchRequest

	if err := c.BodyParser(&searchQuery); err != nil {
		return utilities.InvalidJSON()
	}

	result, err := s.searchService.SearchClubs(searchQuery)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(result)
}

// SearchEvents godoc
//
// @Summary		Searches through events
// @Description	Searches through events
// @ID			search-event
// @Tags      	search
// @Accept		json
// @Produce		json
// @Param		searchQuery	query	search_types.EventSearchRequest	true	"Search Body"
// @Success		200	  {object}	    []search_types.SearchResult[models.Event]
// @Failure     404   {object}		error
// @Failure     500   {object}      error
// @Router		/search/events  [get]
func (s *SearchController) SearchEvents(c *fiber.Ctx) error {
	var searchQuery search_types.EventSearchRequest

	if err := c.BodyParser(&searchQuery); err != nil {
		return utilities.InvalidJSON()
	}

	result, err := s.searchService.SearchEvents(searchQuery)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(result)
}
