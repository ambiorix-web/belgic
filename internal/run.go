package internal

import (
	"log"

	"github.com/devOpifex/eburon/internal/config"
)

func Run(path string) {
	config, err := config.Read(path)

	if err != nil {
		log.Fatal(err)
	}

	config.ListApps()
}
