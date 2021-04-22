package api

import (
	db "github.com/dnswd/kanal/db/sqlc"
	"github.com/gofiber/fiber/v2"
)


type CreateTagParam struct {
	Name string `json:"name"`
}


func (driver *Driver) listTags(c *fiber.Ctx) error {
	tags, err := driver.store.ListTag(c.Context())
	c.JSON(tags)
	return err
}

func (driver *Driver) createTag(c *fiber.Ctx) error {
	params := new(CreateTagParam)
	err := c.BodyParser(params)
	if err != nil { 
		return err
	}
	tag, err := driver.store.AddTag(c.Context(), params.Name)
	c.JSON(tag)
	return err
}

func (driver *Driver) renameTag(c *fiber.Ctx) error {
	params := new(db.RenameTagParams)
	err := c.BodyParser(params)
	if err != nil { 
		return err
	}
	err = driver.store.RenameTag(c.Context(), *params)
	return err
}

func (driver *Driver) deleteTag(c *fiber.Ctx) error {
	params := new(db.RenameTagParams)
	err := c.BodyParser(params)
	if err != nil { 
		return err
	}
	err = driver.store.RenameTag(c.Context(), *params)
	return err
}
