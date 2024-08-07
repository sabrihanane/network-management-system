package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sabrihanane/go-network-api-fiber-postgres/handlers"
)

func Setup(app *fiber.App) {
	ltpGroup := app.Group("/ltp")
	ltpGroup.Get("/get_ltps", handlers.GetLtps)
	ltpGroup.Get("/get_ltp_by_id/:id", handlers.GetLtpById)
	ltpGroup.Get("/get_ltp_by_name/:name", handlers.GetLtpByName)
	ltpGroup.Post("/create_ltp", handlers.CreateLtp)
	ltpGroup.Put("/update_ltp", handlers.UpdateLtp)
	ltpGroup.Delete("/delete_ltp_by_id/:id", handlers.DeleteLtpById)

	nodeGroup := app.Group("/node")
	nodeGroup.Get("/get_nodes", handlers.GetNodes)
	nodeGroup.Get("/get_node_by_id/:id", handlers.GetNodeById)
	nodeGroup.Get("/get_node_by_name/:name", handlers.GetNodeByName)
	nodeGroup.Post("/create_node", handlers.CreateNode)
	nodeGroup.Put("/update_node", handlers.UpdateNode)

	linkGroup := app.Group("/link")
	linkGroup.Get("/get_links", handlers.GetLinks)
	linkGroup.Get("/get_link_by_id/:id", handlers.GetLinkById)
	linkGroup.Get("/get_link_by_name/:name", handlers.GetLinkByName)
	linkGroup.Post("/create_link", handlers.CreateLink)
	linkGroup.Put("/update_link", handlers.UpdateLink)

	subnetGroup := app.Group("/subnet")
	subnetGroup.Get("/get_subnets", handlers.GetSubnets)
	subnetGroup.Get("/get_subnet_by_name/:name", handlers.GetSubnetByName)
	subnetGroup.Post("/create_subnet", handlers.CreateSubnet)

}
