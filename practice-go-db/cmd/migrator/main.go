package main

import (
	"go-db/pkg/config"
	"go-db/pkg/migrator"
	"go-db/pkg/psql"
)

func main() {
	if err := config.LoadEnv(".env"); err != nil {
		panic(err)
	}
	cfg := config.NewConfig()
	dbm, err := psql.NewDbManager(cfg.GetConnectionUrl(), cfg.GetMaxConnections())
	if err != nil {
		panic(err)
	}

	m := migrator.NewMigrator(dbm.GetDB(), "file://migrations", "postgres")

	err = m.Up()
	if err != nil {
		panic(err)
	}
}
