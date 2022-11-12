package repositories

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	log "github.com/sirupsen/logrus"
)

type Repositories interface {
	Sessions() SessionsRepository
	Users() UsersRepository
}

type repositories struct {
	cfg                *config.Config
	log                *log.Logger
	db                 db.DB
	usersRepository    UsersRepository
	sessionsRepository SessionsRepository
}

func New(cfg *config.Config, log *log.Logger, db db.DB) Repositories {
	usersRepository := NewUsersRepository(cfg, log, db)
	sessionsRepository := NewSessionsRepository(cfg, log, db)

	return &repositories{
		cfg:                cfg,
		log:                log,
		db:                 db,
		usersRepository:    usersRepository,
		sessionsRepository: sessionsRepository,
	}
}

func (r *repositories) Sessions() SessionsRepository {
	return r.sessionsRepository
}

func (r *repositories) Users() UsersRepository {
	return r.usersRepository
}
