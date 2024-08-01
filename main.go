package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sabrihanane/go-network-api-fiber-postgres/database"
	"github.com/sabrihanane/go-network-api-fiber-postgres/routes"
)

func main() {
	app := fiber.New()

	database.Connect()
	database.AutoMigrate()

	routes.Setup(app)

	app.Listen(":3000")
}
