package service

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/repository"
)

type ExampleService interface {
	GetByID(ctx context.Context) (string, error)
}

type exampleService struct {
	cfg                *config.Config
	sessionsRepository repository.SessionsRepository
	usersRepository    repository.UsersRepository
}

func NewExampleService(cfg *config.Config, sessionsRepository repository.SessionsRepository, usersRepository repository.UsersRepository) ExampleService {

	return &exampleService{
		cfg:                cfg,
		sessionsRepository: sessionsRepository,
		usersRepository:    usersRepository,
	}
}

func (e *exampleService) GetByID(ctx context.Context) (string, error) {
	res, err := e.sessionsRepository.GetByID(ctx)
	if err != nil {
		return "", err
	}
	return res, nil
}
