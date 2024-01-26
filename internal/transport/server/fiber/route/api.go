package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/logicalangel/tashil_test/internal/transport/server/fiber/controller"
)

func SetApiRoutes(server *fiber.App, api controller.Api) {
	docAPI := server.Group("/api")
	docAPI.Get("/", api.GetAll)
	docAPI.Post("/", api.Create)
	docAPI.Get("/:api_id", api.Get)
	docAPI.Post("/:api_id/start", api.Start)
	docAPI.Post("/:api_id/stop", api.Stop)
	docAPI.Delete("/:api_id", api.Delete)
}
