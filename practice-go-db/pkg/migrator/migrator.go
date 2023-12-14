package migrator

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type migrator struct {
	db        *sql.DB
	sourceUrl string
	dbName    string
}

func NewMigrator(db *sql.DB, sourceUrl, dbName string) *migrator {
	return &migrator{
		sourceUrl: sourceUrl,
		db:        db,
		dbName:    dbName,
	}
}

// steps must be upper zero
func (m migrator) Up(step int) error {
	return m.moveMigration(step)
}

// steps must be upper zero
func (m migrator) Down(step int) error {
	return m.moveMigration(-step)
}

func (m migrator) ActualUpdate() error {
	return m.moveMigration()
}

func (m migrator) moveMigration(steps ...int) error {
	if len(steps) > 1 {
		return fmt.Errorf("too many arguments, must be only one or zero")
	}
	driver, err := postgres.WithInstance(m.db, &postgres.Config{})
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithDatabaseInstance(
		m.sourceUrl,
		m.dbName,
		driver,
	)
	if err != nil {
		return err
	}

	if len(steps) == 1 {
		return migration.Steps(steps[0])
	}

	return migration.Up()
}
