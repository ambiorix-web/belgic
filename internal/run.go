package internal

import (
	"log"
	"os"

	"github.com/devOpifex/belgic/internal/config"
)

type loadBalancer struct {
	Config   config.Config
	Backends config.Backends
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

	backs, err := config.RunApps()

	if err != nil {
		lb.ErrorLog.Fatal(err)
	}

	lb.Backends = backs

	lb.InfoLog.Printf("Running %v child processes", len(lb.Backends))

	for _, back := range backs {
		if back.Err != nil {
			back.SetLive(false)
			back := lb.Config.RunApp()
			lb.Backends = append(lb.Backends, back)
		}
	}

	lb.InfoLog.Printf("Running load balancer on %v", lb.Config.Port)

	err = lb.StartApp()
	if err != nil {
		lb.ErrorLog.Fatal(err)
	}
}
