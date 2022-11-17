package migrations

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/fastid/fastid/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "modernc.org/sqlite"
)

//go:embed sql/*.sql
var schemaFs embed.FS

type Migration interface {
	Upgrade() error
	Downgrade() error
}

type migration struct {
	cfg    *config.Config
	db     *sql.DB
	driver *database.Driver
}

func New(cfg *config.Config, driverName string) (Migration, error) {
	var dsn string

	if driverName == "sqlite" {
		dsn = ":memory:"
	} else if driverName == "postgres" {
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s application_name=%s sslmode=%s search_path=%s",
			cfg.DATABASE.Host,
			cfg.DATABASE.Port,
			cfg.DATABASE.User,
			cfg.DATABASE.Password,
			cfg.DATABASE.DBName,
			cfg.DATABASE.ApplicationName,
			cfg.DATABASE.SslMode,
			cfg.DATABASE.Scheme,
		)
	} else {
		return nil, fmt.Errorf("Unable to load the driver %s", driverName)
	}

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	var driver database.Driver
	if driverName == "sqlite" {
		driver, err = sqlite.WithInstance(db, &sqlite.Config{})
		if err != nil {
			return nil, err
		}
	} else if driverName == "postgres" {
		driver, err = postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return nil, err
		}
	}

	return &migration{cfg: cfg, db: db, driver: &driver}, nil
}

func (m *migration) Upgrade() error {

	dirs, err := iofs.New(schemaFs, "sql")
	if err != nil {
		return err
	}

	mig, err := migrate.NewWithInstance("iofs", dirs, m.cfg.DBName, *m.driver)
	if err != nil {
		return err
	}

	if err := mig.Up(); err != nil {
		return err
	}
	return nil
}

func (m *migration) Downgrade() error {
	dirs, err := iofs.New(schemaFs, "sql")
	if err != nil {
		return err
	}

	mig, err := migrate.NewWithInstance("iofs", dirs, m.cfg.DBName, *m.driver)
	if err != nil {
		return err
	}

	if err := mig.Down(); err != nil {
		return err
	}
	return nil
}
