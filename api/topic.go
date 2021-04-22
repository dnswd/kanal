package api

import (
	"github.com/gofiber/fiber/v2"
)

type CreateTopicParam struct {
	Name string `json:"name"`
}

func (driver *Driver) createTopic(c *fiber.Ctx) error {
	param := new(CreateTopicParam)
	err := c.BodyParser(param)

	err = driver.store.CreateTopic(c.Context(), param.Name)
	return err
}

func (driver *Driver) listTopic(c *fiber.Ctx) error {
	topics, err := driver.store.ListTopic(c.Context())
	c.JSON(topics)
	return err
}