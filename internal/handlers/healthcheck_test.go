package handlers

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/repositories"
	"github.com/fastid/fastid/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	cfg, _ := config.New("../../configs/fastid.yml")

	// Logger
	log := logger.New(cfg)
	log.Infoln("Starting the server")

	database, err := db.NewDB(cfg, "sqlite3")
	require.NoError(t, err)

	// Storage
	repos := repositories.New(cfg, log, database)

	// Service
	srv := services.New(cfg, log, repos)

	// Handlers
	handler := New(cfg, log, srv)

	e := echo.New()
	group := e.Group("/api/v1")
	handler.Register(group)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/healthcheck/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	healthCheck := handler.HealthCheck().Get()
	err = healthCheck(c)
	require.NoError(t, err)
	require.Equal(t, rec.Body.String(), "{}\n")
}
