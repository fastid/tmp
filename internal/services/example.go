package services

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/repositories"
	log "github.com/sirupsen/logrus"
)

type ExampleService interface {
	GetByID(ctx context.Context) (string, error)
}

type exampleService struct {
	cfg          *config.Config
	log          *log.Logger
	repositories repositories.Repositories
}

func NewExampleService(cfg *config.Config, log *log.Logger, repositories repositories.Repositories) ExampleService {
	return &exampleService{
		cfg:          cfg,
		log:          log,
		repositories: repositories,
	}
}

func (e *exampleService) GetByID(ctx context.Context) (string, error) {
	res, err := e.repositories.Sessions().GetByID(ctx)
	if err != nil {
		return "", err
	}
	return res, nil
}
