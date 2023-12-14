package main

import (
	"context"
	"fmt"
	"go-db/internal/core/dto"
	"go-db/internal/data/repository"
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

	dRepository := repository.NewDepartmentRepository(dbm)
	res, err := dRepository.AddDepartment(context.Background(), &dto.AddDepartmentDTO{
		Title: "IT",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(*res)
}
