package handlers

import (
	"github.com/fastid/fastid/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthCheckHandler interface {
	Register(router *echo.Group)
	get() echo.HandlerFunc
}

type healthCheckHandler struct {
	exampleService service.ExampleService
}

func NewHealthCheckHandler(exampleService service.ExampleService) *healthCheckHandler {
	return &healthCheckHandler{exampleService: exampleService}
}

func (h *healthCheckHandler) Register(router *echo.Group) {
	router.Add("GET", "/healthcheck/", h.get())
}

func (h *healthCheckHandler) get() echo.HandlerFunc {
	return func(e echo.Context) error {
		res, err := h.exampleService.GetByID(e.Request().Context())
		if err != nil {
			return err
		}
		return e.String(http.StatusOK, res)

		//e.Request().Context()
		//return e.JSON(http.StatusOK, make(map[string]string))
	}
}
