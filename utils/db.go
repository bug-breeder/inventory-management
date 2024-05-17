package utils

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() *Database {
	return &Database{}
}

// Init initializes the database connection
func (d *Database) Init(connStr string) error {
	fmt.Println("Establishing connection to database...")

	var err error
	d.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = d.DB.Ping()
	if err != nil {
		return err
	}

	var dbName string
	err = d.DB.QueryRow("SELECT current_database()").Scan(&dbName)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully connected to database: %s\n", dbName)
	return nil
}

// Close closes the database connection
func (d *Database) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}

// RunMigrations runs the database migrations
func (d *Database) RunMigrations(migrationPath string) error {
	driver, err := postgres.WithInstance(d.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		"postgres", driver)
	if err != nil {
		return err
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
