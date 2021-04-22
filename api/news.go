package api

import (
	"github.com/gofiber/fiber/v2"
)

func (driver *Driver) index(c *fiber.Ctx) error {
	c.SendString("Hello world")
	return c.SendStatus(200)
}