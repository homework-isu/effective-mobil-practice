package repository

// must run only with test-env

import (
	"errors"
	"go-db/internal/core/domain"
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
