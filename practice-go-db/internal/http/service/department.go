package service

import (
	"context"
	"go-db/internal/core/domain"
	"go-db/internal/core/dto"
)

type DepartmentService interface {
	AddDepartment(context.Context, *dto.AddDepartmentDTO) (*domain.Department, error)
	DeleteDepartment(context.Context, *dto.IdDerartmentDTO) error
	GetDepartmentById(context.Context, *dto.IdDerartmentDTO) (*domain.Department, error)
	GetDepartments(context.Context, *dto.LimitOffsetDTO) ([]domain.Department, error)
	RenameDepartment(context.Context, *dto.RenameDerartmentDTO) (*domain.Department, error)
}
