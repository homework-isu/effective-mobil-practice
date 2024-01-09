package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-db/internal/core/domain"
	"go-db/internal/core/dto"

	core_errors "go-db/internal/core/errors"
	"go-db/internal/core/service"
	"go-db/internal/data/repository"
	"go-db/internal/data/unsafe"
	"go-db/internal/http/handler"

	"go-db/pkg/config"
	"go-db/pkg/psql"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
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
		inp     string
		out     string
		success bool
		code    int
	}{
		{
			name:    "ok",
			inp:     "IT",
			out:     "IT",
			success: true,
			code:    http.StatusOK,
		},
		{
			name:    "repiatable",
			inp:     "IT",
			out:     core_errors.ErrorFailToAddDepartment.Error(),
			success: false,
			code:    http.StatusBadRequest,
		},
	}

	cfg, err := getInstanceConfig()
	if err != nil {
		t.Fatal(err)
	}

	psgClient, err := psql.NewDbManager(cfg.GetConnectionUrl(), cfg.GetMaxConnections())
	if err != nil {
		t.Fatal(err)
	}

	departmentRepository := repository.NewDepartmentRepository(psgClient)
	departmentServcie := service.NewDepartmentService(departmentRepository)
	departmentHandler := handler.NewDepartmentHandler(departmentServcie, 10000)
	api := NewApi(cfg.GetHttpPort(), departmentHandler)

	tableCleaner := unsafe.NewUnsafeDepartmentRepository(psgClient)

	if err := tableCleaner.CleanTable(context.Background()); err != nil {
		t.Log("fail to clean departments table")
		t.Fatal(err)
	}

	defer tableCleaner.CleanTable(context.Background())

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dto := &dto.AddDepartmentDTO{Title: test.inp}
			dtoBytes, err := json.Marshal(dto)
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest("POST", "/department", bytes.NewBuffer(dtoBytes))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			api.server.ServeHTTP(w, req)

			assert.Equal(t, test.code, w.Code)

			if test.success {
				department := &domain.Department{}
				if err := json.NewDecoder(w.Body).Decode(&department); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, test.out, department.Title)
			} else {
				assert.Equal(t, fmt.Sprintf(`{"error":"%s"}`, test.out), w.Body.String())
			}
		})
	}

}
