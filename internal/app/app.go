package app

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/handlers"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/repository"
	"github.com/fastid/fastid/internal/service"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run(cfg *config.Config) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	group := e.Group("/api/v1")

	// Logger
	log := logger.NewLogger(cfg)
	log.Infoln("Starting the server")

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
