package app

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/handlers"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/repository"
	"github.com/fastid/fastid/internal/service"
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
	}
	//} else if strings.HasPrefix(c.Path(), "/api/v1/healthcheck/") {
	//	return true
	//}
	return false
}

func Run(cfg *config.Config) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	group := e.Group("/api/v1")

	// Logger
	log := logger.NewLogger(cfg)
	log.Infoln("Starting the server")

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
		},
	}))

	// DB
	database, err := db.NewDB(cfg, "postgres")
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Storage
	usersRepository := repository.NewUsersRepository(cfg, log, database)
	sessionsRepository := repository.NewSessionsRepository(cfg, log, database)

	// Service
	exampleService := service.NewExampleService(cfg, sessionsRepository, usersRepository)

	// Handlers
	healthCheckHandler := handlers.NewHealthCheckHandler(exampleService)
	healthCheckHandler.Register(group)

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
