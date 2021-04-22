package api

import (
	"database/sql"
	"strconv"

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
	idParam:= c.Params("id")
	err := c.BodyParser(params)
	id, err := strconv.ParseInt(idParam, 10, 32)
	params.ID = int32(id)
	if err != nil { 
		return err
	}
	err = driver.store.RenameTag(c.Context(), *params)
	return err
}

func (driver *Driver) deleteTag(c *fiber.Ctx) error {
	idParam:= c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil { 
		return err
	}
	err = driver.store.DeleteTag(c.Context(), int32(id))
	return err
}

func (driver *Driver) AddTagToBlog(c *fiber.Ctx) error {
	param := new(db.AddTagToBlogParams)
	idParam:= c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil { 
		return err
	}
	tagidParam:= c.Params("tagid")
	tagid, err := strconv.ParseInt(tagidParam, 10, 32)
	if err != nil { 
		return err
	}
	param.NewsID = sql.NullInt32{ Int32: int32(id), Valid: true}
	param.TagID = sql.NullInt32{ Int32: int32(tagid), Valid: true}

	tag, err := driver.store.AddTagToBlog(c.Context(), *param)
	c.JSON(tag)
	return err
}
