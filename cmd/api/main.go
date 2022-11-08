package main

import (
	"github.com/fastid/fastid/internal/app"
	"github.com/fastid/fastid/internal/config"
	"log"
)

func main() {
	cfg, err := config.NewConfig("configs/config.yml")
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}
