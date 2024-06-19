package utilities

import (
	"github.com/gofiber/fiber/v2"
)

type Platform string

const (
	PlatformMobile Platform = "mobile"
	PlatformWeb    Platform = "web"
)

func GetPlatform(c *fiber.Ctx) Platform {
	switch c.Get("Platform", "") {
	case "mobile":
		return PlatformMobile
	case "web":
		return PlatformWeb
	default:
		return PlatformMobile
	}
}
