package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/fastid/fastid/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type DB interface {
	GetConnectContext(ctx context.Context) (*sqlx.Conn, error)
	GetConnect() *sqlx.DB
	Close() error
}

type database struct {
	cfg *config.Config
	db  *sqlx.DB
}

func NewDB(cfg *config.Config, driverName string) (DB, error) {

	var dsn string

	if driverName == "sqlite3" {
		dsn = ":memory:"
	} else if driverName == "postgres" {
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s application_name=%s sslmode=%s",
			cfg.DATABASE.Host,
			cfg.DATABASE.Port,
			cfg.DATABASE.User,
			cfg.DATABASE.Password,
			cfg.DATABASE.DBName,
			cfg.DATABASE.ApplicationName,
			cfg.DATABASE.SslMode,
		)
	} else {
		return nil, errors.New(fmt.Sprintf("Unable to load the driver %s", driverName))
	}

	db, err := sqlx.Connect(driverName, dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DATABASE.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DATABASE.MaxIdleConns)
	db.SetConnMaxLifetime(time.Second * time.Duration(cfg.DATABASE.ConnMaxLifetime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(cfg.DATABASE.ConnMaxIdleTime))

	return &database{db: db, cfg: cfg}, nil
}

func (d *database) GetConnect() *sqlx.DB {
	return d.db
}

func (d *database) GetConnectContext(ctx context.Context) (*sqlx.Conn, error) {
	return d.db.Connx(ctx)
}

func (d *database) Close() error {
	return d.db.Close()
}