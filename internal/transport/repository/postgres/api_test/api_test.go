package api_test

import (
	"context"
	"errors"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/logicalangel/tashil_test/config"
	"github.com/logicalangel/tashil_test/internal/consts"
	"github.com/logicalangel/tashil_test/internal/model"
	"github.com/logicalangel/tashil_test/internal/transport/database/postgres"
	"github.com/logicalangel/tashil_test/internal/transport/repository/postgres/api"
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

func TestGetAll(t *testing.T) {
	db := postgres.New(config.Postgres{Dsn: "postgresql://postgres:postgres@127.0.0.1:5432/postgres"})
	db.GetConnection().AutoMigrate(&model.Api{})

	r := api.New(db)

	_, err := r.GetAll(context.TODO())
	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	db := postgres.New(config.Postgres{Dsn: "postgresql://postgres:postgres@127.0.0.1:5432/postgres"})
	db.GetConnection().AutoMigrate(&model.Api{})

	r := api.New(db)

	_, err := r.Get(context.TODO(), 1)
	if err != nil {
		if !errors.Is(err, consts.ErrApiNotfound) {
			t.Error(err)
			return
		}
	}

	createdApi, err := r.Create(context.TODO(), "google.com", "get", 1, map[string]interface{}{}, "")
	if err != nil {
		t.Error(err)
		return
	}

	gotApi, err := r.Get(context.TODO(), createdApi.ID)
	if err != nil {
		t.Error(err)
		return
	}

	if gotApi.ID < 1 {
		t.Error("got api is not filled")
		return
	}

	if createdApi.ID < 1 {
		t.Error("created api is not filled")
		return
	}

	if gotApi.ID != createdApi.ID {
		t.Error("created and got api is not equal")
		return
	}
}

func TestDelete(t *testing.T) {
	db := postgres.New(config.Postgres{Dsn: "postgresql://postgres:postgres@127.0.0.1:5432/postgres"})
	db.GetConnection().AutoMigrate(&model.Api{})

	r := api.New(db)

	createdApi, err := r.Create(context.TODO(), "google.com", "get", 1, map[string]interface{}{}, "")
	if err != nil {
		t.Error(err)
		return
	}

	err = r.Delete(context.TODO(), createdApi.ID)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestStart(t *testing.T) {
	db := postgres.New(config.Postgres{Dsn: "postgresql://postgres:postgres@127.0.0.1:5432/postgres"})
	db.GetConnection().AutoMigrate(&model.Api{})

	r := api.New(db)

	createdApi, err := r.Create(context.TODO(), "google.com", "get", 1, map[string]interface{}{}, "")
	if err != nil {
		t.Error(err)
		return
	}

	gotApi, err := r.Start(context.TODO(), createdApi.ID)
	if err != nil {
		t.Error(err)
		return
	}
	if gotApi.ID < 1 {
		t.Error("api marshal is not correct")
		return
	}
}

func TestStop(t *testing.T) {
	db := postgres.New(config.Postgres{Dsn: "postgresql://postgres:postgres@127.0.0.1:5432/postgres"})
	db.GetConnection().AutoMigrate(&model.Api{})

	r := api.New(db)

	createdApi, err := r.Create(context.TODO(), "google.com", "get", 1, map[string]interface{}{}, "")
	if err != nil {
		t.Error(err)
		return
	}

	gotApi, err := r.Stop(context.TODO(), createdApi.ID)
	if err != nil {
		t.Error(err)
		return
	}
	if gotApi.ID < 1 {
		t.Error("api marshal is not correct")
		return
	}
}
