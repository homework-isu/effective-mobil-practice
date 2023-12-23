package main

import (
	"go-db/pkg/config"
	"go-db/pkg/migrator"
	"go-db/pkg/psql"
	"go-db/pkg/flag_parser"
)



func main() {
	if err := config.LoadEnv("./test_env/.env"); err != nil {
		panic(err)
	}
	cfg := config.NewConfig()
	dbm, err := psql.NewDbManager(cfg.GetConnectionUrl(), cfg.GetMaxConnections())
	if err != nil {
		panic(err)
	}

	m := migrator.NewMigrator(dbm.GetDB(), "file://migrations", "postgres")
	parser := flagparser.Parser{}

	flag, steps, err := parser.GetAction()
	if err != nil {
		panic(err)
	}

	switch flag {
		case flagparser.ActionUp:
			if err := m.Up(steps); err != nil {
				panic(err)
			}
		case flagparser.ActionDown:
			if err := m.Down(steps); err != nil {
				panic(err)
			}
		case flagparser.ActionActual:
			if err := m.ActualUpdate(); err != nil {
				panic(err)
			}
	}
}
