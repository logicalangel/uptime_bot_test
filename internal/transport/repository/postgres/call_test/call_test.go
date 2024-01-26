package call_test

import (
	"context"
	"errors"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/logicalangel/tashil_test/config"
	"github.com/logicalangel/tashil_test/internal/consts"
	"github.com/logicalangel/tashil_test/internal/model"
	"github.com/logicalangel/tashil_test/internal/transport/database/postgres"
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

func TestGet(t *testing.T) {
	db := postgres.New(config.Postgres{Dsn: "postgresql://postgres:postgres@127.0.0.1:5432/postgres"})
	db.GetConnection().AutoMigrate(&model.Call{})

	r := call.New(db)

	_, err := r.Create(context.TODO(), 1, 200, 300)
	if err != nil {
		if !errors.Is(err, consts.ErrApiNotfound) {
			t.Error(err)
			return
		}
	}

	_, err = r.Get(context.TODO(), 1)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestCreate(t *testing.T) {
	db := postgres.New(config.Postgres{Dsn: "postgresql://postgres:postgres@127.0.0.1:5432/postgres"})
	db.GetConnection().AutoMigrate(&model.Call{})

	r := call.New(db)

	call, err := r.Create(context.TODO(), 1, 200, 300)
	if err != nil {
		if !errors.Is(err, consts.ErrApiNotfound) {
			t.Error(err)
			return
		}
	}

	if call.ID < 1 {
		t.Error("id is not set")
	}
}
