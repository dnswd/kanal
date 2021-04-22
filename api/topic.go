package api

import (
	"strconv"

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
	idParam:= c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 32)

	err = driver.store.DeleteTag(c.Context(), int32(id))
	return err
}