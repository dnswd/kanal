package api

import (
	db "github.com/dnswd/kanal/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

func (driver *Driver) index(c *fiber.Ctx) error {
	c.SendString("Hello world")
	return c.SendStatus(200)
}

func (driver *Driver) createNews(c *fiber.Ctx) error {
	param := new(db.CreateNewsParams)
	err := c.BodyParser(param)

	news, err := driver.store.CreateNews(c.Context(), *param)
	c.JSON(news)
	return err
}

func (driver *Driver) listNews(c *fiber.Ctx) error {
	news, err := driver.store.ListNews(c.Context())
	c.JSON(news)
	return err
}