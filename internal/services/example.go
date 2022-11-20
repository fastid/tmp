package services

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/repositories"
	log "github.com/sirupsen/logrus"
)

type ExampleService interface {
	Register(srv Services)
	GetByID(ctx context.Context) (string, error)
}

type exampleService struct {
	cfg          *config.Config
	log          *log.Logger
	repositories repositories.Repositories
	services     Services
}

func NewExampleService(cfg *config.Config, log *log.Logger, repositories repositories.Repositories) ExampleService {
	return &exampleService{
		cfg:          cfg,
		log:          log,
		repositories: repositories,
	}
}

func (s *exampleService) Register(srv Services) {
	s.services = srv
}

func (s *exampleService) GetByID(ctx context.Context) (string, error) {
	res, err := s.repositories.Sessions().GetByID(ctx)
	if err != nil {
		return "", err
	}
	return res, nil
}
