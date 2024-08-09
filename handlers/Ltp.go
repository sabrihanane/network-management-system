package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sabrihanane/go-network-api-fiber-postgres/database"
	"github.com/sabrihanane/go-network-api-fiber-postgres/models"
)

func GetLtps(c *fiber.Ctx) error {
	var ltps []models.Ltp
	database.DB.Find(&ltps)
	return c.JSON(ltps)
}

func GetLtpById(c *fiber.Ctx) error {
	id := c.Params("id")
	ltpId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}

	var ltp models.Ltp
	err1 := database.DB.First(&ltp, uint(ltpId))
	if err1.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Ltp not found"})
	}
	return c.Status(fiber.StatusOK).JSON(ltp)
}

func GetLtpByName(c *fiber.Ctx) error {
	name := c.Params("name")
	var ltp models.Ltp
	err := database.DB.Where("name = ?", name).First(&ltp)
	if err.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Ltp not found"})
	}
	return c.Status(fiber.StatusOK).JSON(ltp)
}

func CreateLtp(c *fiber.Ctx) error {
	ltp := new(models.Ltp)

	if err := c.BodyParser(ltp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	// Validate the data
	if ltp.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ltp name can not be empty!"})
	}

	var ltpModel models.Ltp
	result := database.DB.Where("name = ? AND description = ? ", ltp.Name, ltp.Description).First(&ltpModel)

	if result.Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ltp with the same name and the same description already exists"})
	}

	database.DB.Create(&ltp)
	return c.JSON(ltp)
}

func UpdateLtp(c *fiber.Ctx) error {
	ltp := new(models.Ltp)
	if err := c.BodyParser(ltp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	result := database.DB.Model(&models.Ltp{}).Where("id = ?", ltp.ID).Updates(ltp)
	if result.Error != nil {
		return c.Status(fiber.StatusNotModified).JSON(fiber.Map{"error": "Updating ltp failed"})
	}
	return c.JSON(ltp)
}

func DeleteLtpById(c *fiber.Ctx) error {
	ltp := new(models.Ltp)

	id := c.Params("id")
	if err := database.DB.First(&ltp, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Ltp not found",
		})
	}

	if err := database.DB.Delete(&ltp).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete ltp",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Ltp has been deleted"})
}

// get node by name and get link by name and ltp by name, unique name
