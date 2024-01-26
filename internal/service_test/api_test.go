package service_test

import (
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
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
