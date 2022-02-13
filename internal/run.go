package internal

import (
	"log"
	"os"

	"github.com/devOpifex/eburon/internal/app"
	"github.com/devOpifex/eburon/internal/config"
)

// Run run eburon.
func Run() {
	config, err := config.Read()

	config.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	config.ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	if err != nil {
		config.ErrorLog.Fatal(err)
	}

	apps, err := config.ListApps()

	if err != nil {
		config.ErrorLog.Fatal(err)
	}

	cmds, err := config.RunApps(apps)

	if err != nil {
		config.ErrorLog.Fatal(err)
	}

	for _, cmd := range cmds {
		if cmd.Err != nil {
			config.ErrorLog.Fatal(cmd.Err)
		}

		config.InfoLog.Printf("%v is running on port %v", cmd.Application, cmd.Port)
	}

	err = app.StartApp(config, cmds)
	if err != nil {
		config.ErrorLog.Fatal(err)
	}
}
