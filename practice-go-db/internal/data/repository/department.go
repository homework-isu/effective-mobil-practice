package repository

import (
	"context"
	"database/sql"
	"go-db/internal/core/domain"
	"go-db/internal/core/dto"
)

type connectionGetter interface {
	GetConnection(ctx context.Context) (*sql.Conn, error)
}

type departmentRepository struct {
	dbManager connectionGetter
}

func NewDepartmentRepository(dbManager connectionGetter) *departmentRepository {
	return &departmentRepository{
		dbManager: dbManager,
	}
}

func (r departmentRepository) AddDepartment(ctx context.Context, dto *dto.AddDepartmentDTO) (*domain.Department, error) {
	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return nil, err
	}

	stmt, err := conn.QueryContext(ctx, "INSERT INTO departments (title) VALUES ($1) RETURNING id, title", dto.Title)
	if err != nil {
		return nil, err
	}

	return scanDepartment(stmt)
}


type sqlScanner interface {
	Scan(...any) error
	Next() bool
}


func scanDepartment(r sqlScanner) (*domain.Department, error) {
	result := &domain.Department{}

	if r.Next() {
		if err := r.Scan(&result.Id); err != nil {
			return nil, err
		}
		if err := r.Scan(&result.Title); err != nil {
			return nil, err
		}
	} else {
		return nil, sql.ErrNoRows
	}

	return result, nil
}