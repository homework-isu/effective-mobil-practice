package service

import (
	"context"
	"errors"
	"go-db/internal/core/domain"
	"go-db/internal/core/dto"
	core_errors "go-db/internal/core/errors"
	"testing"
	"testing/quick"
	"time"

	mock_repository "go-db/pkg/mocks/repository"

	"github.com/golang/mock/gomock"
)

func TestAddDepartment(t *testing.T) {
	ctl := gomock.NewController(t)
	repo := mock_repository.NewMockDepartmentRepository(ctl)
	service := NewDepartmentService(repo)

	okTest := func(title string) bool {
		dto := &dto.AddDepartmentDTO{
			Title: title,
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		repo.EXPECT().AddDepartment(ctx, dto).Return(&domain.Department{Id: 1, Title: title}, nil)

		result, err := service.AddDepartment(ctx, dto)

		if result == nil {
			t.Log(err.Error())
			t.Fatal("result is nil")
		}
		return result.Title == title
	}

	withError := func(title string) bool {
		dto := &dto.AddDepartmentDTO{
			Title: title,
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		errMustBe := core_errors.ErrorFailToAddDepartment
		repo.EXPECT().AddDepartment(ctx, dto).Return(nil, errMustBe)

		_, err := service.AddDepartment(ctx, dto)

		if err == nil {
			t.Fatal("error is nil")
		}
		return errors.Is(err, errMustBe)
	}

	cfg := &quick.Config{
		MaxCount: 5,
	}

	if err := quick.Check(okTest, cfg); err != nil {
		t.Fatal(err)
	}
	if err := quick.Check(withError, cfg); err != nil {
		t.Fatal(err)
	}
}
