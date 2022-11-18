package handlers

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/services"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UnsealHandler interface {
	Register(router *echo.Group)
	Post() echo.HandlerFunc
}

type unsealHandler struct {
	cfg *config.Config
	log *log.Logger
	srv services.Services
}

func NewUnsealHandler(cfg *config.Config, log *log.Logger, srv services.Services) UnsealHandler {
	return &unsealHandler{cfg: cfg, log: log, srv: srv}
}

func (h *unsealHandler) Register(router *echo.Group) {
	router.Add("POST", "/unseal/", h.Post())
}

func (h *unsealHandler) Post() echo.HandlerFunc {
	return func(e echo.Context) error {
		return e.JSON(http.StatusOK, make(map[string]string))
	}
}
