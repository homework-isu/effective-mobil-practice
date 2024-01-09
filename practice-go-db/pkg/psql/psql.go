package psql

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/jackc/pgx/v5/pgxpool"
)

type dbManager struct {
	db *sql.DB
}

func NewDbManager(url string, maxConnsInPool int) (*dbManager, error) {
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxConnsInPool)
	return &dbManager{
		db: db,
	}, nil
}

func (dbm dbManager) GetConnection(ctx context.Context) (*sql.Conn, error) {
	return dbm.db.Conn(ctx)
}

func (dbm dbManager) GetDB() *sql.DB {
	return dbm.db
}

func GetConnectionsPull(ctx context.Context, url string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}
	_, err = pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
