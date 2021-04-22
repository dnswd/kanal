package api

import (
	db "github.com/dnswd/kanal/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

type CreateTopicParam struct {
	Name string `json:"name"`
}

func (driver *Driver) listTopics(c *fiber.Ctx) error {
	topics, err := driver.store.ListTopic(c.Context())
	c.JSON(topics)
	return err
}

func (driver *Driver) createTopic(c *fiber.Ctx) error {
	param := new(CreateTopicParam)
	err := c.BodyParser(param)

	err = driver.store.CreateTopic(c.Context(), param.Name)
	return err
}

func (driver *Driver) renameTopic(c *fiber.Ctx) error {
	param := new(db.RenameTopicParams)
	err := c.BodyParser(param)

	err = driver.store.RenameTopic(c.Context(), *param)
	return err
}