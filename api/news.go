package api

import (
	"strconv"

	db "github.com/dnswd/kanal/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

func (driver *Driver) listNews(c *fiber.Ctx) (err error) {
	status := c.Query("status")
	switch status {
		case "draft":
		case "deleted":
		case "published":
			news, err := driver.store.ListNewsByStatus(c.Context(), db.Status(status))
			c.JSON(news)
			return err
		default:
			news, err := driver.store.ListNews(c.Context())
			c.JSON(news)
			return err
	}
	return
}

func (driver *Driver) listNewsByTopic(c *fiber.Ctx) (err error) {
	idParam:= c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	status := c.Query("status")
	switch status {
	case "draft":
	case "deleted":
	case "published":
		params := new(db.ListNewsByTopicAndStatusParams)
		params.ID = int32(id)
		params.Status = db.Status(status)
		news, err := driver.store.ListNewsByTopicAndStatus(c.Context(), *params)
		c.JSON(news)
		return err
	default:
		news, err := driver.store.ListNewsByTopic(c.Context(), int32(id))
		c.JSON(news)
		return err
	}
	return
}

func (driver *Driver) createArticle(c *fiber.Ctx) (err error) {
	param := new(db.CreateNewsParams)
	err = c.BodyParser(param)
	if err != nil { 
		return err
	}
	news, err := driver.store.CreateNews(c.Context(), *param)
	c.JSON(news)
	return err
}

func (driver *Driver) retrieveArticle(c *fiber.Ctx) (err error) {
	idParam:= c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil { 
		return err
	}
	news, err := driver.store.GetNews(c.Context(), int32(id)) //int32(id))
	c.JSON(news)
	return err
}

func (driver *Driver) replaceArticle(c *fiber.Ctx) error {
	param := new(db.UpdateNewsParams)
	err := c.BodyParser(param)
	if err != nil { 
		return err
	}
	err = driver.store.UpdateNews(c.Context(), *param)
	c.Status(200)
	return err
}

func (driver *Driver) hardDeleteArticle(c *fiber.Ctx) error {
	idParam:= c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil { 
		return err
	}
	err = driver.store.HardDeleteNews(c.Context(), int32(id))
	return err
}

func (driver *Driver) publishArticle(c *fiber.Ctx) error {
	param := new(db.PublishNewsParams)
	err := c.BodyParser(param)
	if err != nil { 
		return err
	}
	err = driver.store.PublishNews(c.Context(), *param)
	return err
}

func (driver *Driver) unpublishArticle(c *fiber.Ctx) error {
	idParam:= c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil { 
		return err
	}
	err = driver.store.UnpublishNews(c.Context(), int32(id))
	return err
}

func (driver *Driver) getNewsTags(c *fiber.Ctx) error {
	idParam:= c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil { 
		return err
	}
	tags, err := driver.store.GetNewsTags(c.Context(), int32(id))
	c.JSON(tags)
	return err
}
