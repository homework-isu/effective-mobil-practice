package service

import (
	"context"
	"fmt"
	"go-db/internal/core/ports/repository"
	"go-db/internal/core/domain"
	"go-db/internal/core/dto"
	core_errors "go-db/internal/core/errors"
)

type departmentService struct {
	repo repository.DepartmentRepository
}

func NewDepartmentService(repo repository.DepartmentRepository) *departmentService {
	return &departmentService{
		repo: repo,
	}
}

func (s *departmentService) AddDepartment(ctx context.Context, dto *dto.AddDepartmentDTO) (*domain.Department, error) {
	op := "Department Service: AddDepartment"
	department, err := s.repo.AddDepartment(ctx, dto)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, core_errors.ErrorFailToAddDepartment)
	}
	return department, nil
}

func (s *departmentService) DeleteDepartment(ctx context.Context, dto *dto.IdDerartmentDTO) error {
	op := "Department Service: DeleteDepartment"
	err := s.repo.DeleteDepartment(ctx, dto)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *departmentService) RenameDepartment(ctx context.Context, dto *dto.RenameDerartmentDTO) (*domain.Department, error) {
	op := "Department Service: RenameDepartment"
	department, err := s.repo.RenameDepartment(ctx, dto)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, core_errors.ErrorNoSuchDepartment)
	}
	return department, nil
}

func (s *departmentService) GetDepartmentById(ctx context.Context, dto *dto.IdDerartmentDTO) (*domain.Department, error) {
	op := "Department Service: GetDepartmentById"
	department, err := s.repo.GetDepartmentById(ctx, dto)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, core_errors.ErrorNoSuchDepartment)
	}
	return department, nil
}

func (s *departmentService) GetDepartments(ctx context.Context, dto *dto.LimitOffsetDTO) ([]domain.Department, error) {
	op := "Department Service: GetDepartments"
	department, err := s.repo.GetDepartments(ctx, dto)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return department, nil
}
