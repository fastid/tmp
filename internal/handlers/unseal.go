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
	router.Add("GET", "/unseal/", h.Get())
}

func (h *unsealHandler) Post() echo.HandlerFunc {
	return func(e echo.Context) error {
		h.srv.Crypt().SecretKey("51b9fcd4a19cd0d6f76ecdcbb427895047bfcdb084f570c01d4af1648e20f162")
		return e.JSON(http.StatusOK, make(map[string]string))
	}
}

func (h *unsealHandler) Get() echo.HandlerFunc {
	return func(e echo.Context) error {
		return e.JSON(http.StatusOK, make(map[string]string))
	}
}
