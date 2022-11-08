package repository

import (
	"fmt"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	log "github.com/sirupsen/logrus"
)

type UsersRepository interface {
	Create() error
	GetByID() error
}

type usersRepository struct {
	cfg *config.Config
	log *log.Logger
	db  db.DB
}

func NewUsersRepository(cfg *config.Config, log *log.Logger, db db.DB) UsersRepository {
	return &usersRepository{cfg: cfg, log: log, db: db}
}

func (r *usersRepository) Create() error {
	fmt.Println("Create")
	return nil
}

func (r *usersRepository) GetByID() error {
	fmt.Println("Create")
	return nil
}
