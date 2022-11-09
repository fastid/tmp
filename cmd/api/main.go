package main

import (
	"flag"
	"github.com/fastid/fastid/internal/app"
	"github.com/fastid/fastid/internal/config"
	"log"
)

func main() {
	argPathConfig := flag.String("config", "fastid.yml", "a string")

	flag.Parse()

	cfg, err := config.NewConfig(*argPathConfig)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}
