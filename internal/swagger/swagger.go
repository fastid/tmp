package swagger

import (
	"embed"
	"github.com/labstack/echo/v4"
	"io/fs"
	"net/http"
)

//go:embed assert/*
var embededFiles embed.FS

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(embededFiles, "assert")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}

func New(e *echo.Echo) {
	assetHandler := http.FileServer(getFileSystem())
	e.GET("/swagger/*", echo.WrapHandler(http.StripPrefix("/swagger/", assetHandler)))
}
