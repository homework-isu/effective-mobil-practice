package unsafe

import (
	"context"
	"database/sql"
)

type connectionGetter interface {
	GetConnection(ctx context.Context) (*sql.Conn, error)
}


type unsafeDepartmentRepository struct {
	dbManager connectionGetter
}

// Use it only for testing purposes
func NewUnsafeDepartmentRepository(dbManager connectionGetter) *unsafeDepartmentRepository {
	return &unsafeDepartmentRepository{
		dbManager: dbManager,
	}
}

// Use this function for deleting all data from table "departments"
func (r unsafeDepartmentRepository) CleanTable(ctx context.Context) error {
	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return err
	}

	return conn.QueryRowContext(ctx, "DELETE FROM departments WHERE id > 0").Err()
}
