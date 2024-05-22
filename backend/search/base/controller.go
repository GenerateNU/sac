package base

import (
	"net/http"

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
// @ID			search-club
// @Tags      	search
// @Accept		json
// @Produce		json
// @Param		searchQuery	query	models.SearchQueryParams	true	"Search Body"
// @Success		200	  {object}	    []models.ClubSearchResult
// @Failure     404   {object}		error
// @Failure     500   {object}      error
// @Router		/search/clubs  [get]
func (s *SearchController) SearchClubs(c *fiber.Ctx) error {
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
// @ID			search-club
// @Tags      	search
// @Accept		json
// @Produce		json
// @Param		searchQuery	query	models.SearchQueryParams	true	"Search Body"
// @Success		200	  {object}	    []models.EventsSearchResult
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
