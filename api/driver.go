package api

import (
	db "github.com/dnswd/kanal/db/sqlc"
	"github.com/dnswd/kanal/util"
	"github.com/gofiber/fiber/v2"
)


type Driver struct {
	config util.Config
	router *fiber.App
	store db.Store
}

func InitDriver(config util.Config, store db.Store) *Driver {
	driver := &Driver{
		config: config,
		store: store,
	}

	driver.setupRouter()
	return driver 
}

func (driver *Driver) Listen() error {
	return driver.router.Listen(driver.config.Host)
}

func (driver *Driver) setupRouter() {
	router := fiber.New()

	router.Get("/news", driver.listNews)
	router.Get("/news/topic", driver.listTopics)
	router.Post("/news/topic", driver.createTopic)
	router.Get("/news/:topic", driver.listNewsByTopic)
	router.Put("/news/:topic", driver.renameTopic)
	
	router.Post("/article", driver.createArticle)
	router.Put("/article", driver.replaceArticle)
	router.Get("/article/:id", driver.retrieveArticle)
	router.Delete("/article/:id", driver.hardDeleteArticle)
	router.Get("/article/:id/tags", driver.getNewsTags)
	router.Post("/article/:id/tags/:tagid", driver.AddTagToBlog)
	router.Post("/article/:id/publish", driver.publishArticle)
	router.Post("/article/:id/unpublish", driver.unpublishArticle)

	router.Get("/tags", driver.listTags)
	router.Post("/tags", driver.createTag)
	router.Put("/tags/:id", driver.renameTag)
	router.Delete("/tags/:id", driver.deleteTag)
	
	driver.router = router
}