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

	router.Get("/", driver.index)
	
	driver.router = router
}