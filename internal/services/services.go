package services

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/repositories"
	log "github.com/sirupsen/logrus"
)

type Services interface {
	Example() ExampleService
}

type services struct {
	cfg            *config.Config
	log            *log.Logger
	exampleService ExampleService
}

func New(cfg *config.Config, log *log.Logger, repositories repositories.Repositories) Services {
	exampleService := NewExampleService(cfg, log, repositories)
	return &services{cfg: cfg, log: log, exampleService: exampleService}
}

func (s *services) Example() ExampleService {
	return s.exampleService
}
