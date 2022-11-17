package main

import (
	"flag"
	"github.com/fastid/fastid/internal/app"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/migrations"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	var pathConfig string
	var runServer bool
	var help bool
	var migrate bool
	var up bool
	var down bool

	flag.StringVar(&pathConfig, "config", "fastid.yml", "The path to the configuration file")
	flag.BoolVar(&runServer, "run", false, "Starts the FastID server")
	flag.BoolVar(&help, "help", false, "List available commands")
	flag.BoolVar(&migrate, "migrate", false, "Performs migration to the database")
	flag.BoolVar(&up, "up", false, "Updating migrations")
	flag.BoolVar(&down, "down", false, "Downgrade migrations")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(1)
	}

	cfg, err := config.New(pathConfig)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	if migrate {
		migration, err := migrations.New(cfg, cfg.DriverName)
		if err != nil {
			log.Fatalf("Migration error: %s", err)
		}

		if up {
			if err := migration.Upgrade(); err != nil {
				log.Fatalf("Migrations error: %s", err)
			}
		} else if down {
			if err := migration.Downgrade(); err != nil {
				log.Fatalf("Migrations error: %s", err)
			}
		} else {
			log.Fatalf("No action passed to perform migration")
		}
	}

	if runServer {
		app.Run(cfg)
	} else if migrate {
		return
	} else {
		flag.PrintDefaults()
	}
}
