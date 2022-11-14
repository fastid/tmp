package app

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/handlers"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/repositories"
	"github.com/fastid/fastid/internal/services"
	"github.com/fastid/fastid/internal/swagger"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func urlSkipper(c echo.Context) bool {
	if strings.HasPrefix(c.Path(), "/metrics") {
		return true
	} else if strings.HasPrefix(c.Path(), "/api/v1/healthcheck/") {
		return true
	}
	return false
}

func Run(cfg *config.Config) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	group := e.Group("/api/v1")

	// Logger
	log := logger.New(cfg)
	log.Infoln("Starting the server")

	if cfg.DATABASE.DriverName == "sqlite3" {
		log.Warningln("Warning! Your server is running with sqlite, the data will be deleted after reboot")
	}

	// Prometheus
	prom := prometheus.NewPrometheus("fastid", urlSkipper)
	prom.Use(e)

	// Middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogRequestID: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"uri":          values.URI,
				"status":       values.Status,
				"x-request-id": values.RequestID,
			}).Info("request")

			return nil
		}},
	))

	// DB
	database, err := db.NewDB(cfg, cfg.DATABASE.DriverName)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Repository
	repo := repositories.New(cfg, log, database)

	// Service
	srv := services.New(cfg, log, repo)

	// Swagger
	sw := swagger.New(cfg, log)
	sw.Register(e)

	// Handlers
	handler := handlers.New(cfg, log, srv)
	handler.Register(group)

	go func() {
		if err := e.Start(cfg.HTTP.Listen); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctxShutdown); err != nil {
		e.Logger.Fatal(err)
	}
}
