package internal

import (
	"fmt"
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
	Stdout   chan string
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
	lb.Stdout = make(chan string)
	lb.RunApps()
}

// collect Collect messages from stdout
func collect(c chan string) {
	// this currently does not work
	// no issue on the Go side
	// it's a problem or caveat with R
	// or probably httpuv
	// it seems the app is launch
	// in an other subprocess
	for {
		fmt.Println(<-c)
	}
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
		back.RunApp(lb.Stdout)
		lb.Backends = append(lb.Backends, back)
	}

	lb.InfoLog.Printf("Running %v child processes", len(lb.Backends))
	lb.InfoLog.Printf("Running load balancer on %v", lb.Config.Port)

	go collect(lb.Stdout)
	lb.StartApp()
}
