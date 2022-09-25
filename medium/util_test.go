package medium

import (
	"testing"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/samber/do"

	"github.com/syuparn/fridgesim/ent"
	"github.com/syuparn/fridgesim/ent/enttest"
)

func mockIDInjector(id string) func(*do.Injector) (func() string, error) {
	newIDGenerator := func() string {
		return id
	}
	return func(*do.Injector) (func() string, error) {
		return newIDGenerator, nil
	}
}

func mockDBInjector(t *testing.T) (func(*do.Injector) (*ent.Client, error), func()) {
	postgres := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().Port(15432))
	err := postgres.Start()
	if err != nil {
		postgres.Stop()
		t.Fatal(err)
	}

	client := enttest.Open(t, "postgres", "host=127.0.0.1 port=15432 user=postgres password=postgres dbname=postgres sslmode=disable")
	teardown := func() {
		client.Close()
		postgres.Stop()
	}

	newDBClient := func(i *do.Injector) (*ent.Client, error) {
		t.Log("replace DB client with mock")
		return client, nil
	}

	return newDBClient, teardown
}

func setupConfig(t *testing.T) {
	t.Setenv("FRIDGESIM_PORT", "18080")
	t.Setenv("FRIDGESIM_DBHOST", "localhost")
	t.Setenv("FRIDGESIM_DBPORT", "15432")
	t.Setenv("FRIDGESIM_DBUSER", "postgres")
	t.Setenv("FRIDGESIM_DBPASSWORD", "pass")
}
