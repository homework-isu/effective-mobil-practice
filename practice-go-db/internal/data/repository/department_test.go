package repository

// must run only with test-env

import (
	"errors"
	"go-db/internal/core/domain"
	"go-db/internal/data/unsafe"
	"go-db/internal/core/dto"
	core_errors "go-db/internal/core/errors"
	"go-db/pkg/config"
	"go-db/pkg/psql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"context"
)

var (
	env_file_path   = "../../../test_env/.env"
	config_instance = false
)

func getInstanceConfig() (*config.Config, error) {
	if !config_instance {
		if err := config.LoadEnv(env_file_path); err != nil {
			return nil, err
		}
		config_instance = true
	}
	return config.NewConfig(), nil
}

func TestAddDepartment(t *testing.T) {
	tests := []struct {
		name    string
		inp     *dto.AddDepartmentDTO
		out     *domain.Department
		timeOut time.Duration
		err     error
	}{
		{
			name: "ok",
			inp: &dto.AddDepartmentDTO{
				Title: "TestDepartment",
			},
			out: &domain.Department{
				Title: "TestDepartment",
			},
			timeOut: time.Second * 5,
			err:     nil,
		},
		{
			name: "repiated-title",
			inp: &dto.AddDepartmentDTO{
				Title: "TestDepartment",
			},
			out:     nil,
			timeOut: time.Second * 5,
			err:     core_errors.ErrorFailToAddDepartment,
		},
	}

	cfg, err := getInstanceConfig()
	if err != nil {
		t.Fatal(err)
	}
	psqlClient, err := psql.NewDbManager(cfg.GetConnectionUrl(), cfg.GetMaxConnections())
	if err != nil {
		t.Fatal(err)
	}

	repo := NewDepartmentRepository(psqlClient)
	tableCleaner := unsafe.NewUnsafeDepartmentRepository(psqlClient)

	if err := tableCleaner.CleanTable(context.Background()); err != nil {
		t.Log("fail to clean departments table")
		t.Fatal(err)
	}

	defer tableCleaner.CleanTable(context.Background())

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), test.timeOut)
			defer cancel()

			result, err := repo.AddDepartment(ctx, test.inp)
			
			if test.err != nil {
				if err == nil {
					t.Fatal("error is nil")
				}
				err = errors.Unwrap(err)
				assert.Equal(t, test.err, errors.Unwrap(err))
			} else {
				if result == nil {
					t.Log(test.name)
					t.Log(err.Error())
					t.Fatal("result is nil")
				}
				assert.Equal(t, test.out.Title, result.Title)
			}
		})
	}
}


func TestGetDepartments(t *testing.T) {
	tests := []struct {
		name    string
		toInsert []dto.AddDepartmentDTO
		inp     *dto.LimitOffsetDTO
		out     []domain.Department
		timeOut time.Duration
		err     error
	}{
		{
			name: "ok",
			toInsert: []dto.AddDepartmentDTO{
				{Title: "IT"},
				{Title: "Marketing"},
				{Title: "IoT"},
			},
			inp: &dto.LimitOffsetDTO{},
			out: []domain.Department{
				{Title: "IT"},
				{Title: "Marketing"},
				{Title: "IoT"},
			},
			timeOut: time.Second * 5,
			err:     nil,
		},
		{
			name: "ok with limit",
			toInsert: []dto.AddDepartmentDTO{
				{Title: "IT"},
				{Title: "Marketing"},
				{Title: "IoT"},
			},
			inp: &dto.LimitOffsetDTO{
				Limit: 2,
			},
			out: []domain.Department{
				{Title: "IT"},
				{Title: "Marketing"},
			},
			timeOut: time.Second * 5,
			err:     nil,
		},
		{
			name: "ok with offset",
			toInsert: []dto.AddDepartmentDTO{
				{Title: "IT"},
				{Title: "Marketing"},
				{Title: "IoT"},
			},
			inp: &dto.LimitOffsetDTO{
				Offset: 1,
			},
			out: []domain.Department{
				{Title: "Marketing"},
				{Title: "IoT"},
			},
			timeOut: time.Second * 5,
			err:     nil,
		},
		{
			name: "ok with limit and offset",
			toInsert: []dto.AddDepartmentDTO{
				{Title: "IT"},
				{Title: "Marketing"},
				{Title: "IoT"},
				{Title: "HR"},
			},
			inp: &dto.LimitOffsetDTO{
				Limit: 2,
				Offset: 1,
			},
			out: []domain.Department{
				{Title: "Marketing"},
				{Title: "IoT"},
			},
			timeOut: time.Second * 5,
			err:     nil,
		},
	}

	cfg, err := getInstanceConfig()
	if err != nil {
		t.Fatal(err)
	}
	psqlClient, err := psql.NewDbManager(cfg.GetConnectionUrl(), cfg.GetMaxConnections())
	if err != nil {
		t.Fatal(err)
	}

	repo := NewDepartmentRepository(psqlClient)
	tableCleaner := unsafe.NewUnsafeDepartmentRepository(psqlClient)

	if err := tableCleaner.CleanTable(context.Background()); err != nil {
		t.Log("fail to clean departments table")
		t.Fatal(err)
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), test.timeOut)
			defer cancel()
			defer tableCleaner.CleanTable(context.Background())

			for _, dto := range test.toInsert {
				_, err := repo.AddDepartment(ctx, &dto)
				if err != nil {
					t.Fatal(err)
				}
			}
			
			result, err := repo.GetDepartments(ctx, test.inp)
			
			if test.err != nil {
				if err == nil {
					t.Fatal("error is nil")
				}
				err = errors.Unwrap(err)
				assert.Equal(t, test.err, errors.Unwrap(err))
			} else {
				if result == nil {
					t.Log(err.Error())
					t.Fatal("result is nil")
				}
				t.Log(result)
				assert.Equal(t, len(test.out), len(result))
				for i := range result {
					assert.Equal(t, test.out[i].Title, result[i].Title)
				}
			}
		})
	}
}