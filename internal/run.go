package internal

import (
	"log"
	"os"

	"github.com/devOpifex/belgic/internal/app"
	"github.com/devOpifex/belgic/internal/config"
)

type loadBalancer struct {
	App      app.Application
	Config   config.Config
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

// Run run belgic.
func Run() {
	var lb loadBalancer
	lb.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	lb.ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	config, err := config.Read()

	if err != nil {
		lb.ErrorLog.Fatal(err)
	}

	lb.Config = config

	cmds, err := config.RunApps()

	if err != nil {
		lb.ErrorLog.Fatal(err)
	}

	for _, cmd := range cmds {
		if cmd.Err != nil {
			lb.ErrorLog.Fatal(cmd.Err)
		}

		lb.InfoLog.Printf("%v is running on port %v", cmd.Name, cmd.Port)
	}

	err = app.StartApp(config, cmds)
	if err != nil {
		lb.ErrorLog.Fatal(err)
	}
}
