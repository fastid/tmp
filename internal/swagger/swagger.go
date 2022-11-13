package swagger

import (
	"embed"
	"github.com/fastid/fastid/internal/config"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"net/http"
)

//go:embed assert/*
var embededFiles embed.FS

type Swagger interface {
	Register(e *echo.Echo)
	getFileSystem() http.FileSystem
}

type swagger struct {
	cfg *config.Config
	log *log.Logger
}

func New(cfg *config.Config, log *log.Logger) Swagger {
	return &swagger{cfg: cfg, log: log}
}

func (s *swagger) getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(embededFiles, "assert")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}

func (s *swagger) Register(e *echo.Echo) {
	assetHandler := http.FileServer(s.getFileSystem())
	e.GET("/swagger/*", echo.WrapHandler(http.StripPrefix("/swagger/", assetHandler)))
}
