package middleware

import (
	go_json "github.com/goccy/go-json"
	"gorm.io/gorm"

	"errors"
	"fmt"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ExtractID func(*fiber.Ctx) (*uuid.UUID, error)

func AttachExtractor(middleware func(c *fiber.Ctx, extractor ExtractID) error, extractor ExtractID) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return middleware(c, extractor)
	}
}

func ExtractFromParams(param string) ExtractID {
	return func(c *fiber.Ctx) (*uuid.UUID, error) {
		return utilities.ValidateID(c.Params(param))
	}
}

func ExtractFromBody(key string) ExtractID {
	return func(c *fiber.Ctx) (*uuid.UUID, error) {
		var bodyMap map[string]string

		bodyBytes := c.BodyRaw()

		if len(bodyBytes) == 0 {
			return nil, errors.New("empty request body")
		}

		err := go_json.Unmarshal(bodyBytes, &bodyMap)
		if err != nil {
			return nil, err
		}

		id, ok := bodyMap[key]
		if !ok {
			return nil, fmt.Errorf("missing key %s in request body", key)
		}

		return utilities.ValidateID(id)
	}
}

type QueryFn func(*gorm.DB, uuid.UUID) (*uuid.UUID, error)

func ExtractFromQuery(db *gorm.DB, id uuid.UUID, query QueryFn) ExtractID {
	return func(c *fiber.Ctx) (*uuid.UUID, error) {
		id, err := query(db, id)
		if err != nil {
			return nil, err
		}

		return id, nil
	}
}

func ExtractFromParamIntoQuery(param string, db *gorm.DB, query QueryFn) ExtractID {
	return func(c *fiber.Ctx) (*uuid.UUID, error) {
		id, err := ExtractFromParams(param)(c)
		if err != nil {
			return nil, err
		}

		ExtractFromQuery(db, *id, query)

		return id, nil
	}
}
