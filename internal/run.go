package internal

import (
	"log"

	"github.com/devOpifex/eburon/internal/config"
)

func Run() {
	config, err := config.Read()

	if err != nil {
		log.Fatal(err)
	}

	apps, err := config.ListApps()

	if err != nil {
		log.Fatal(err)
	}

	config.RunApps(apps)
}
