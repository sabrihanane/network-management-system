package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sabrihanane/go-network-api-fiber-postgres/database"
	"github.com/sabrihanane/go-network-api-fiber-postgres/models"
)

func GetLinks(c *fiber.Ctx) error {
	var links []models.Link
	database.DB.Find(&links)
	return c.JSON(links)
}

func GetLinkById(c *fiber.Ctx) error {
	id := c.Params("id")
	linkId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}

	var link models.Link
	err1 := database.DB.First(&link, uint(linkId))
	if err1.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Link not found"})
	}
	return c.Status(fiber.StatusOK).JSON(link)
}

func GetLinkByName(c *fiber.Ctx) error {
	name := c.Params("name")
	var link models.Link
	err := database.DB.Where("name = ?", name).First(&link)
	if err.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Link not found"})
	}
	return c.Status(fiber.StatusOK).JSON(link)
}

func CreateLink(c *fiber.Ctx) error {
	link := new(models.Link)

	if err := c.BodyParser(link); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	// Validate the data
	if link.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Link name can not be empty!"})
	}

	var linkModel models.Link
	result := database.DB.Where("name = ?", link.Name).First(&linkModel)

	if result.Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Link with the same name already exists"})
	}

	database.DB.Create(&link)
	return c.JSON(link)
}

func UpdateLink(c *fiber.Ctx) error {
	link := new(models.Link)
	if err := c.BodyParser(link); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	result := database.DB.Model(&models.Link{}).Where("id = ?", link.ID).Updates(link)
	if result.Error != nil {
		return c.Status(fiber.StatusNotModified).JSON(fiber.Map{"error": "Updating link failed"})
	}
	return c.JSON(link)
}
