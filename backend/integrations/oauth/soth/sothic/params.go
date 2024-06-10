package sothic

import "github.com/gofiber/fiber/v2"

type Params struct {
	c *fiber.Ctx
}

func (p *Params) Get(key string) string {
	return p.c.Query(key)
}
