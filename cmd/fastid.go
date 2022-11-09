package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/fastid/fastid/internal/app"
	"github.com/fastid/fastid/internal/config"
	"log"
	"os"
)

var (
	//go:embed LICENSE
	dir embed.FS
)

func main() {
	//argPathConfig := flag.String("config", "fastid.yml", "a string")

	var pathConfig string
	var runServer bool
	var help bool
	var migrate bool
	var license bool

	flag.StringVar(&pathConfig, "config", "fastid.yml", "The path to the configuration file")
	flag.BoolVar(&runServer, "run", false, "Starts the FastID server")
	flag.BoolVar(&help, "help", false, "List available commands")
	flag.BoolVar(&migrate, "migrate", false, "Performs migration to the database")
	flag.BoolVar(&license, "license", false, "Show current license")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(1)
	}

	cfg, err := config.NewConfig(pathConfig)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	if runServer {
		app.Run(cfg)
	} else if license {
		entries, _ := dir.ReadFile("LICENSE")
		fmt.Println(string(entries))
	}
}
