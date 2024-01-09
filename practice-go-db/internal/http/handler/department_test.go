package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-db/internal/core/domain"
	"go-db/internal/core/dto"
	core_errors "go-db/internal/core/errors"

	mock_service "go-db/pkg/mocks/service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"time"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func TestAddDepartment(t *testing.T) {

	tests := []struct {
		name    string
		inp     string
		out     string
		prepair func(*mock_service.MockDepartmentService, string)
		success bool
		code    int
	}{
		{
			name: "ok",
			inp:  "IT",
			out:  "IT",
			prepair: func(service *mock_service.MockDepartmentService, title string) {
				service.EXPECT().AddDepartment(gomock.Any(), gomock.Any()).Return(&domain.Department{Id: 1, Title: title}, nil)
			},
			success: true,
			code: http.StatusOK,
		},
		{
			name: "with error",
			inp:  "IT",
			out:  core_errors.ErrorFailToAddDepartment.Error(),
			prepair: func(service *mock_service.MockDepartmentService, title string) {
				service.EXPECT().AddDepartment(gomock.Any(), gomock.Any()).Return(nil, core_errors.ErrorFailToAddDepartment)
			},
			success: false,
			code: http.StatusBadRequest,
		},
	}

	r := gin.Default()

	ctl := gomock.NewController(t)
	service := mock_service.NewMockDepartmentService(ctl)

	handler := NewDepartmentHandler(service, time.Second*5)
	r.POST("/department", handler.AddDepartment)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.prepair != nil {
				test.prepair(service, test.inp)
			}

			dto := &dto.AddDepartmentDTO{Title: test.inp}
			dtoBytes, err := json.Marshal(dto)
			if err != nil {
				t.Fatal(err)
			}

			req:= httptest.NewRequest("POST", "/department", bytes.NewBuffer(dtoBytes))

			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, test.code, w.Code)
			if test.success {
				assert.Equal(t, fmt.Sprintf(`{"id":%d,"title":"%s"}`, 1, test.out), w.Body.String())
			} else {
				assert.Equal(t, fmt.Sprintf(`{"error":"%s"}`, test.out), w.Body.String())
			}
		})
	}
}
