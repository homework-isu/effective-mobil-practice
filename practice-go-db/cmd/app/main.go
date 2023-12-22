package main

import (

	// "time"

	"fmt"
	"go-db/internal/core/service"
	"go-db/internal/data/repository"
	"go-db/internal/http/api"
	"go-db/internal/http/hendler"
	"time"

	"go-db/pkg/config"
	"go-db/pkg/psql"
)

func main() {
	if err := config.LoadEnv(".env"); err != nil {
		panic(err)
	}

	cfg := config.NewConfig()

	pgClient, err := psql.NewDbManager(cfg.GetConnectionUrl(), cfg.GetMaxConnections())
	if err != nil {
		panic(err)
	}

	departmentRepository := repository.NewDepartmentRepository(pgClient)
	departmentServcie := service.NewDepartmentService(departmentRepository)
	departmentHendler := hendler.NewDepartmentHendler(departmentServcie, time.Duration(10 * time.Second))
	api := api.NewApi(cfg.GetHttpPort(), departmentHendler)

	fmt.Println("server works on :", cfg.GetHttpPort())
	if err := api.Run(); err != nil {
		panic(err)
	}
}
