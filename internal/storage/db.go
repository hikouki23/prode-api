package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"go.nhat.io/otelsql"
	"hikouki.com/prode/internal/config"
)

type SQLDB struct {
	db         client
	statements map[string]*sql.Stmt
}

type client interface {
	Prepare(query string) (*sql.Stmt, error)
	Close() error
	BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

func New(cfg config.Config) (*SQLDB, error) {
	register, err := otelsql.Register("postgres", otelsql.TraceQueryWithArgs())
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(register, cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("database could not be initialized: %v", err)
	}

	if cfg.MaxOpenConns != 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	}

	if cfg.MaxIdleConns != 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}

	if cfg.ConnMaxIdleTime != 0 {
		db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	}

	if cfg.ConnMaxLifetime != 0 {
		db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}

	var migrations *migrate.Migrate

	if cfg.MigrationsEnabled {
		if cfg.MigrationsLocation == "" {
			cfg.MigrationsLocation = "file://db/migrations"
		}

		migrations, err = migrate.New(cfg.MigrationsLocation, cfg.ConnectionString)
		if err != nil {
			return nil, fmt.Errorf("database migration could not be initialized: %v", err)
		}

		err = migrations.Up()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return nil, fmt.Errorf("database migration could not be applied: %v", err)
		}
	}

	store := &SQLDB{
		db:         db,
		statements: map[string]*sql.Stmt{},
	}

	err = store.prepareStatements()
	if err != nil {
		return nil, err
	}

	return store, nil
}
