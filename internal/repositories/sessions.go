package repositories

import (
	"context"
	"fmt"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	log "github.com/sirupsen/logrus"
)

type SessionsRepository interface {
	Create() error
	GetByID(ctx context.Context) (string, error)
}

type sessionsRepository struct {
	cfg *config.Config
	log *log.Logger
	db  db.DB
}

func NewSessionsRepository(cfg *config.Config, log *log.Logger, db db.DB) SessionsRepository {
	return &sessionsRepository{cfg: cfg, log: log, db: db}
}

func (r *sessionsRepository) Create() error {
	fmt.Println("Create")
	return nil
}

func (r *sessionsRepository) GetByID(ctx context.Context) (string, error) {
	sql := r.db.GetConnect()

	type Version struct {
		Name string `db:"name"`
	}

	var result Version
	err := sql.GetContext(ctx, &result, "SELECT 'hello word' AS name")
	if err != nil {
		return "", err
	}

	return result.Name, nil
}
