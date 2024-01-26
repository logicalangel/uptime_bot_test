package main

import (
	"context"
	"github.com/logicalangel/tashil_test/config"
	"github.com/logicalangel/tashil_test/internal/service"
	nativeClient "github.com/logicalangel/tashil_test/internal/transport/client/native"
	"github.com/logicalangel/tashil_test/internal/transport/database/postgres"
	"github.com/logicalangel/tashil_test/internal/transport/repository/postgres/api"
	"github.com/logicalangel/tashil_test/internal/transport/repository/postgres/call"
	"time"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	db := postgres.New(cfg.Postgres)

	apiRepository := api.New(db)
	callRepository := call.New(db)
	client := nativeClient.New()

	callService := service.NewCall(client, apiRepository, callRepository)

	ticker := time.NewTicker(15 * time.Second)
	for {
		select {
		case _ = <-ticker.C:
			ctx := context.TODO()
			err = callService.CallScheduledApis(ctx)
			if err != nil {
				panic(err)
			}
		}
	}
}
