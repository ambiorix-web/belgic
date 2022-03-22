package internal

import (
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/devOpifex/belgic/internal/config"
)

type loadBalancer struct {
	Config   config.Config
	Backends []config.Backend
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

// Run run belgic.
func Run() {
	var lb loadBalancer
	lb.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	lb.ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	conf, err := config.Read()

	if err != nil {
		lb.ErrorLog.Fatal(err)
	}

	lb.Config = conf
	lb.RunApps()
}

// RunApps runs all the applications found in the directory.
func (lb loadBalancer) RunApps() {
	ncpus, err := strconv.Atoi(lb.Config.Backends)

	// error we assume it was set to max
	// default to max CPUs
	if err != nil {
		ncpus = runtime.GOMAXPROCS(runtime.NumCPU())
	}

	for i := 0; i < ncpus; i++ {
		var back config.Backend
		back.Rpath = lb.Config.Path
		back.RunApp()
		lb.Backends = append(lb.Backends, back)
	}

	lb.InfoLog.Printf("Running %v child processes", len(lb.Backends))
	lb.InfoLog.Printf("Running load balancer on %v", lb.Config.Port)

	lb.StartApp()
}
