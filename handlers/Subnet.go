package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sabrihanane/go-network-api-fiber-postgres/database"
	"github.com/sabrihanane/go-network-api-fiber-postgres/models"
)

func GetSubnets(c *fiber.Ctx) error {
	var subnets []models.Subnet
	database.DB.Find(&subnets)
	return c.JSON(subnets)
}

func GetSubnetByName(c *fiber.Ctx) error {
	name := c.Params("name")
	var subnet models.Subnet
	err := database.DB.Where("name = ?", name).First(&subnet)
	if err.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Subnet not found"})
	}
	return c.Status(fiber.StatusOK).JSON(subnet)
}

func CreateSubnet(c *fiber.Ctx) error {
	subnet := new(models.Subnet)

	if err := c.BodyParser(subnet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	// Validate the data
	if subnet.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Subnet name can not be empty!"})
	}

	var subnetModel models.Subnet
	result := database.DB.Where("name = ?", subnet.Name).First(&subnetModel)

	if result.Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Subnet with the same name already exists"})
	}

	database.DB.Create(&subnet)
	return c.JSON(subnet)
}
