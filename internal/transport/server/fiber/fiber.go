package fiber

import (
	"github.com/gofiber/swagger"
	_ "github.com/logicalangel/tashil_test/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/logicalangel/tashil_test/config"
	"github.com/logicalangel/tashil_test/internal/transport/server/fiber/controller"
	"github.com/logicalangel/tashil_test/internal/transport/server/fiber/route"
)

func Start(
	cfg config.Server,
	api controller.Api,
) {
	server := fiber.New(fiber.Config{})

	server.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("alive")
	})

	server.Use(cors.New())

	route.SetApiRoutes(server, api)

	server.Get("/swagger/*", swagger.HandlerDefault)

	server.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	err := server.Listen(cfg.Address)
	if err != nil {
		panic(err)
	}
}
