package service

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/repository"
	log "github.com/sirupsen/logrus"
)

type Service interface {
	Example() ExampleService
}

type service struct {
	cfg            *config.Config
	log            *log.Logger
	exampleService ExampleService
}

func NewService(cfg *config.Config, log *log.Logger, repository repository.Repository) Service {
	exampleService := NewExampleService(cfg, log, repository)
	return &service{cfg: cfg, log: log, exampleService: exampleService}
}

func (s *service) Example() ExampleService {
	return s.exampleService
}

//func (r *repository) Sessions() SessionsRepository {
//	return r.sessionsRepository
//}
//
//func (r *repository) Users() UsersRepository {
//	return r.usersRepository
//}
