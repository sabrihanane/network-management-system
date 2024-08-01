package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sabrihanane/go-network-api-fiber-postgres/handlers"
)

func Setup(app *fiber.App) {
	ltpGroup := app.Group("/ltp")
	ltpGroup.Get("/get_ltps", handlers.GetLtps)
	ltpGroup.Get("/get_ltp_by_id/:id", handlers.GetLtpById)
	ltpGroup.Post("/create_ltp", handlers.CreateLtp)
	ltpGroup.Put("/update_ltp", handlers.UpdateLtp)

	nodeGroup := app.Group("/node")
	nodeGroup.Get("/get_nodes", handlers.GetNodes)
	nodeGroup.Get("/get_node_by_id/:id", handlers.GetNodeById)
	nodeGroup.Post("/create_node", handlers.CreateNode)
	nodeGroup.Put("/update_node", handlers.UpdateNode)

	linkGroup := app.Group("/link")
	linkGroup.Get("/get_links", handlers.GetLinks)
	linkGroup.Get("/get_link_by_id/:id", handlers.GetLinkById)
	linkGroup.Post("/create_link", handlers.CreateLink)
	linkGroup.Put("/update_link", handlers.UpdateLink)

	subnetGroup := app.Group("/subnet")
	subnetGroup.Get("/get_subnets", handlers.GetSubnets)
	subnetGroup.Post("/create_subnet", handlers.CreateSubnet)

}
