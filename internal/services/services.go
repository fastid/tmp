package services

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/repositories"
	log "github.com/sirupsen/logrus"
)

type Services interface {
	Example() ExampleService
	Crypt() CryptService
}

type services struct {
	cfg            *config.Config
	log            *log.Logger
	exampleService ExampleService
	cryptService   CryptService
}

func New(cfg *config.Config, log *log.Logger, repositories repositories.Repositories) Services {
	exampleService := NewExampleService(cfg, log, repositories)
	cryptService := NewCryptService(cfg, log)

	srv := services{cfg: cfg, log: log, exampleService: exampleService, cryptService: cryptService}

	exampleService.Register(&srv)
	cryptService.Register(&srv)
	return &srv
}

func (s *services) Example() ExampleService {
	return s.exampleService
}

func (s *services) Crypt() CryptService {
	return s.cryptService
}
