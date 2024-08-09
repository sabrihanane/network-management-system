package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sabrihanane/go-network-api-fiber-postgres/database"
	"github.com/sabrihanane/go-network-api-fiber-postgres/models"
)

func GetNodes(c *fiber.Ctx) error {
	var nodes []models.Node
	database.DB.Find(&nodes)
	return c.JSON(nodes)
}

func GetNodeById(c *fiber.Ctx) error {
	id := c.Params("id")
	nodeId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}

	var node models.Node
	err1 := database.DB.First(&node, uint(nodeId))
	if err1.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Node not found"})
	}
	return c.Status(fiber.StatusOK).JSON(node)

}

func GetNodeByName(c *fiber.Ctx) error {
	name := c.Params("name")
	var node models.Node
	err := database.DB.Where("name = ?", name).First(&node)
	if err.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Node not found"})
	}
	return c.Status(fiber.StatusOK).JSON(node)
}

func CreateNode(c *fiber.Ctx) error {
	node := new(models.Node)

	if err := c.BodyParser(node); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	// Validate the data
	if node.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Node name can not be empty!"})
	}

	var nodeModel models.Node
	result := database.DB.Where("name = ?", node.Name).First(&nodeModel)

	if result.Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Node with the same name already exists"})
	}

	database.DB.Create(&node)
	return c.JSON(&node)
}

func UpdateNode(c *fiber.Ctx) error {
	node := new(models.Node)
	if err := c.BodyParser(node); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	result := database.DB.Model(&models.Node{}).Where("id = ?", node.ID).Updates(node)
	if result.Error != nil {
		return c.Status(fiber.StatusNotModified).JSON(fiber.Map{"error": "Updating node failed"})
	}
	return c.JSON(node)
}
