package repository

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	log "github.com/sirupsen/logrus"
)

type Repository interface {
	Sessions() SessionsRepository
	Users() UsersRepository
}

type repository struct {
	cfg                *config.Config
	log                *log.Logger
	db                 db.DB
	usersRepository    UsersRepository
	sessionsRepository SessionsRepository
}

func NewRepository(cfg *config.Config, log *log.Logger, db db.DB) Repository {
	usersRepository := NewUsersRepository(cfg, log, db)
	sessionsRepository := NewSessionsRepository(cfg, log, db)

	return &repository{
		cfg:                cfg,
		log:                log,
		db:                 db,
		usersRepository:    usersRepository,
		sessionsRepository: sessionsRepository,
	}
}

func (r *repository) Sessions() SessionsRepository {
	return r.sessionsRepository
}

func (r *repository) Users() UsersRepository {
	return r.usersRepository
}
