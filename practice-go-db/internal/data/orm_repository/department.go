package ormrepository

import (
	"context"
	"database/sql"
	"go-db/internal/core/domain"
	"go-db/internal/core/dto"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OrmDepartmentRepository struct {
	db *gorm.DB
}

func NewOrmDepartmentRepository(db *sql.DB) (*OrmDepartmentRepository, error) {
	gormDb, err := gorm.Open(
		postgres.New(postgres.Config{
			Conn: db,
		}),
	)
	if err != nil {
		return nil, err
	}

	return &OrmDepartmentRepository{
		db: gormDb,
	}, nil
}

func (r OrmDepartmentRepository) AddDepartment(ctx context.Context, dto *dto.AddDepartmentDTO) (*domain.Department, error) {
	department := &domain.Department{
		Title: dto.Title,
	}

	result := r.db.WithContext(ctx).Create(&department)

	if result.Error != nil {
		return nil, result.Error
	}

	return department, nil
}

func (r OrmDepartmentRepository) DeleteDepartment(ctx context.Context, dto *dto.IdDerartmentDTO) error {

	result := r.db.Model(&domain.Department{}).Delete(dto)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r OrmDepartmentRepository) RenameDepartment(ctx context.Context, dto *dto.RenameDerartmentDTO) (*domain.Department, error) {
	department := &domain.Department{
		Id: dto.Id,
		Title: dto.NewTitle,
	}
	result := r.db.WithContext(ctx).Save(&department)

	if result.Error != nil {
		return nil, result.Error
	}

	return department, nil
}

func (r OrmDepartmentRepository) GetDepartmentById(ctx context.Context, dto *dto.IdDerartmentDTO) (*domain.Department, error) {
	department := &domain.Department{
		Id: dto.Id,
	}

	result := r.db.WithContext(ctx).First(&department)

	if result.Error != nil {
		return nil, result.Error
	}

	return department, nil
}

func (r OrmDepartmentRepository) GetDepartments(ctx context.Context, dto *dto.LimitOffsetDTO) ([]domain.Department, error) {
	departments := make([]domain.Department, 0, dto.Limit)

	result := r.db.WithContext(ctx)

	if dto.Limit > 0 {
		result = result.Limit(int(dto.Limit))
	}

	if dto.Offset > 0 {
		result = result.Offset(int(dto.Offset))
	}
	
	result = result.Find(&departments)

	if result.Error != nil {
		return nil, result.Error
	}

	return departments, nil
}