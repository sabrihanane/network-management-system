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

func GetNetworkByName(c *fiber.Ctx) error {
	name := c.Params("name")
	var network models.Network
	err := database.DB.Where("name = ?", name).First(&network)
	if err.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Network not found"})
	}
	return c.Status(fiber.StatusOK).JSON(network)
}
