package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sabrihanane/go-network-api-fiber-postgres/database"
	"github.com/sabrihanane/go-network-api-fiber-postgres/models"
)

func GetNetworks(c *fiber.Ctx) error {
	var networks []models.Network
	database.DB.Find(&networks)
	return c.JSON(networks)
}
