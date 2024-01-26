package main

import (
	"github.com/logicalangel/tashil_test/config"
	"github.com/logicalangel/tashil_test/internal/service"
	"github.com/logicalangel/tashil_test/internal/transport/database/postgres"
	"github.com/logicalangel/tashil_test/internal/transport/repository/postgres/api"
	"github.com/logicalangel/tashil_test/internal/transport/repository/postgres/call"
	"github.com/logicalangel/tashil_test/internal/transport/server/fiber"
	"github.com/logicalangel/tashil_test/internal/transport/server/fiber/controller"
)

// @title TashilKar Test App
// @version 1.0
// @description This app handles api tracking
// @contact.name TashilKar
// @contact.email rastegar.parya3@gmail.com
// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	db := postgres.New(cfg.Postgres)

	apiRepository := api.New(db)
	callRepository := call.New(db)

	_apiService := service.NewApi(apiRepository, callRepository)

	apiController := controller.NewApi(_apiService)

	fiber.Start(cfg.Server, apiController)
}
