package main

import (
	"github.com/logicalangel/tashil_test/config"
	"github.com/logicalangel/tashil_test/internal/model"
	"github.com/logicalangel/tashil_test/internal/transport/database/postgres"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	db := postgres.New(cfg.Postgres)

	err = db.GetConnection().AutoMigrate(&model.Api{})
	if err != nil {
		panic(err)
	}

	err = db.GetConnection().AutoMigrate(&model.Call{})
	if err != nil {
		panic(err)
	}
}
