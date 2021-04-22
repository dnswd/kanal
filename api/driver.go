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

	// Topic
	router.Get("/news", driver.listNews)
	router.Post("/news", nil)
	router.Get("/news/:topic", nil)
	router.Put("/news/:topic", nil)
	
	// News
	router.Post("/article", nil)
	router.Get("/article/:id", nil)
	router.Put("/article/:id", nil)
	router.Delete("/article/:id", nil)
	router.Get("/article/:id/tags", nil)
	router.Put("/article/:id/tags", nil)
	router.Get("/article/:id/author", nil)
	router.Post("/article/:id/publish", nil)
	router.Post("/article/:id/unpublish", nil)

	// Tags
	router.Get("/tags", nil)
	router.Post("/tags", nil)
	router.Put("/tags/:id", nil)
	router.Delete("/tags/:id", nil)
	
	driver.router = router
}