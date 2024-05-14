package middleware

import (
	"fmt"

	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(c *fiber.Ctx) error {
	limit, err := parseQueryParam(c, "limit", constants.DEFAULT_LIMIT)
	if err != nil {
		return utilities.BadRequest(err)
	}

	page, err := parseQueryParam(c, "page", constants.DEFAULT_PAGE)
	if err != nil {
		return utilities.BadRequest(err)
	}

	c.Locals(constants.PAGINATION_MIDDLEWARE_KEY, NewPagination(*limit, *page))

	return c.Next()
}

func GetPagination(c *fiber.Ctx) (*Pagination, error) {
	paginationValue := c.Locals(constants.PAGINATION_MIDDLEWARE_KEY)
	if paginationValue == nil {
		return nil, fmt.Errorf("expected pagination to be defined with key %s. ensure your are using the pagination middleware", constants.PAGINATION_MIDDLEWARE_KEY)
	}

	pagination, ok := paginationValue.(*Pagination)
	if !ok {
		return nil, fmt.Errorf("expected pagination to be of type *middleware.Pagination. got %T", paginationValue)
	}

	return pagination, nil
}

type Pagination struct {
	Limit int
	Page  int
}

func NewPagination(limit int, page int) *Pagination {
	return &Pagination{
		Limit: limit,
		Page:  page,
	}
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}

func (p *Pagination) IntoScope(db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.Offset()).Limit(p.Limit)
	}
}

func parseQueryParam(c *fiber.Ctx, param string, defaultValue int) (*int, error) {
	valueStr := c.Query(param, "")
	if valueStr == "" {
		return &defaultValue, nil
	}
	valueInt, err := utilities.ValidateNonNegative(valueStr)
	if err != nil {
		return nil, fmt.Errorf("invalid value for %s. got %s", param, valueStr)
	}
	return valueInt, nil
}
