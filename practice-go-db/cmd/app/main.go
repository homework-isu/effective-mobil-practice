package main

import (
	"context"
	"fmt"
	"go-db/pkg/config"
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
	conn, err := dbm.GetConnection(context.Background())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	res := ""
	_ = conn.QueryRowContext(context.Background(), "select 'hello world from psql'").Scan(&res)
	fmt.Println(res)
}