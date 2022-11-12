package service

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/repository"
	log "github.com/sirupsen/logrus"
)

type ExampleService interface {
	GetByID(ctx context.Context) (string, error)
}

type exampleService struct {
	cfg        *config.Config
	log        *log.Logger
	repository repository.Repository
}

func NewExampleService(cfg *config.Config, log *log.Logger, repository repository.Repository) ExampleService {
	return &exampleService{
		cfg:        cfg,
		log:        log,
		repository: repository,
	}
}

func (e *exampleService) GetByID(ctx context.Context) (string, error) {
	res, err := e.repository.Sessions().GetByID(ctx)
	if err != nil {
		return "", err
	}
	return res, nil
}
