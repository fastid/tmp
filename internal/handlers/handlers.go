package handlers

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/services"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Handlers interface {
	Register(router *echo.Group)
	HealthCheck() HealthCheckHandler
	Unseal() UnsealHandler
}

type handlers struct {
	cfg         *config.Config
	log         *log.Logger
	srv         services.Services
	healthCheck HealthCheckHandler
	unseal      UnsealHandler
}

func New(cfg *config.Config, log *log.Logger, srv services.Services) Handlers {
	healthCheck := NewHealthCheckHandler(cfg, log, srv)
	unseal := NewUnsealHandler(cfg, log, srv)
	return &handlers{
		cfg:         cfg,
		log:         log,
		srv:         srv,
		healthCheck: healthCheck,
		unseal:      unseal,
	}
}

func (h *handlers) Register(router *echo.Group) {
	h.healthCheck.Register(router)
	h.unseal.Register(router)
}

func (h *handlers) HealthCheck() HealthCheckHandler {
	return h.healthCheck
}

func (h *handlers) Unseal() UnsealHandler {
	return h.unseal
}
