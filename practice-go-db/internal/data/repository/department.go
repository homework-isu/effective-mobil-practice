package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-db/internal/core/domain"
	"go-db/internal/core/dto"
)

const (
	TX_KEY = iota
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

func (r departmentRepository) GetDepartmentById(ctx context.Context, dto *dto.IdDerartmentDTO) (*domain.Department, error) {
	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return nil, err
	}

	stmt := conn.QueryRowContext(ctx, "SELECT id, title FROM departments WHERE id = $1", dto.Id)

	return scanDepartment(stmt)
}

func (r departmentRepository) GetDepartments(ctx context.Context, dto *dto.LimitOffsetDTO) ([]domain.Department, error) {
	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return nil, err
	}

	query := "SELECT id, title FROM departments"
	if dto.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", dto.Limit)
	}

	if dto.Offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", dto.Offset)
	}

	fmt.Println(query)
	stmt, err := conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	departments := make([]domain.Department, 0, dto.Limit)
	for stmt.Next() {
		department, err := scanDepartment(stmt)
		if err != nil {
			return nil, err
		}
		departments = append(departments, *department)
	}

	return departments, nil
}


func (r departmentRepository) AddDepartment(ctx context.Context, dto *dto.AddDepartmentDTO) (*domain.Department, error) {
	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return nil, err
	}

	stmt := conn.QueryRowContext(ctx, "INSERT INTO departments (title) VALUES ($1) RETURNING id, title", dto.Title)

	return scanDepartment(stmt)
}

func (r departmentRepository) DeleteDepartment(ctx context.Context, dto *dto.IdDerartmentDTO) error {
	
	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return err
	}

	stmt, err := conn.ExecContext(ctx, "DELETE FROM departments WHERE id = $1", dto.Id)
	if err != nil {
		return err
	}

	if deleted, err := stmt.RowsAffected(); err != nil || deleted == 0{
		return fmt.Errorf("nothing to delete")
	}

	return nil
}

func (r departmentRepository) RenameDepartment(ctx context.Context, dto *dto.RenameDerartmentDTO) (*domain.Department, error) {
	tx, ok := ctx.Value(TX_KEY).(*sql.Tx)
	if ok {
		stmt := tx.QueryRowContext(ctx, "UPDATE departments SET title = $1 WHERE id = $2 RETURNING id, title", dto.NewTitle, dto.Id)

		return scanDepartment(stmt)
	}

	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return nil, err
	}

	stmt := conn.QueryRowContext(ctx, "UPDATE departments SET title = $1 WHERE id = $2 RETURNING id, title", dto.NewTitle, dto.Id)

	return scanDepartment(stmt)
}

func (r departmentRepository) RenameSomeDepartments(ctx context.Context, dto []dto.RenameDerartmentDTO) ([]domain.Department, error) {
	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: false,
	})
	if err != nil {
		return nil, err
	}

	ctxWithTx := context.WithValue(ctx, TX_KEY, tx)

	departments := make([]domain.Department, 0)
	for _, singleDto := range dto {
		department, err := r.RenameDepartment(ctxWithTx, &singleDto)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		departments = append(departments, *department)
	}

	tx.Commit()
	return departments, nil
}

type sqlScanner interface {
	Scan(...any) error
}


func scanDepartment(r sqlScanner) (*domain.Department, error) {
	result := &domain.Department{}

	if err := r.Scan(&result.Id, &result.Title); err != nil {
		return nil, err
	}

	return result, nil
}