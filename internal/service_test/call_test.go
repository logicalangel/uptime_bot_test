package service_test

import (
	"context"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/logicalangel/tashil_test/config"
	"github.com/logicalangel/tashil_test/internal/model"
	"github.com/logicalangel/tashil_test/internal/service"
	"github.com/logicalangel/tashil_test/internal/transport/client/mock"
	"github.com/logicalangel/tashil_test/internal/transport/database/postgres"
	"github.com/logicalangel/tashil_test/internal/transport/repository/postgres/api"
	"github.com/logicalangel/tashil_test/internal/transport/repository/postgres/call"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	postgres := embeddedpostgres.NewDatabase()
	err := postgres.Start()
	if err != nil {
		panic(err)
	}

	exitVal := m.Run()

	err = postgres.Stop()
	if err != nil {
		panic(err)
	}

	os.Exit(exitVal)
}

func TestCallScheduledApis(t *testing.T) {
	db := postgres.New(config.Postgres{Dsn: "postgresql://postgres:postgres@127.0.0.1:5432/postgres"})
	db.GetConnection().AutoMigrate(&model.Call{})

	c := mock.New()

	db.GetConnection().AutoMigrate(&model.Api{})
	db.GetConnection().AutoMigrate(&model.Call{})

	apiR := api.New(db)
	callR := call.New(db)

	s := service.NewCall(c, apiR, callR)

	err := s.CallScheduledApis(context.TODO())
	if err != nil {
		t.Error(err)
		return
	}
}
