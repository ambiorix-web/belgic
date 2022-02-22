package config

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
)

// RCommand represents a single R command.
type Backend struct {
	Err  error
	Cmd  *exec.Cmd
	Port int
	Path string
	Mu   sync.RWMutex
	Live bool
}

// RCommands represents an array of R commands.
type Backends []Backend

// getR retrieves the full path to the R installation.
func getR() (string, error) {
	var p string

	p, err := exec.LookPath("R")

	if err != nil {
		return p, errors.New("could not locate R installation")
	}

	return p, nil
}

// RunApps runs all the applications found in the directory.
func (conf Config) RunApps() (Backends, error) {
	var backs Backends
	ncpus, err := strconv.Atoi(conf.Backends)

	// error we assume it was set to max
	// default to max CPUs
	if err != nil {
		ncpus = runtime.NumCPU()
	}

	for i := 0; i < ncpus; i++ {
		back := conf.RunApp()
		backs = append(backs, back)
	}

	return backs, nil
}

// runApp run a single application.
func (conf Config) RunApp() Backend {
	var back Backend

	cmd, port, err := conf.callApp()

	if err != nil {
		back.Err = err
		return back
	}
	back.Port = port
	back.Path = "http://localhost:" + strconv.Itoa(port)
	back.Cmd = cmd
	back.Err = back.Cmd.Start()
	back.Live = true

	if back.Err != nil {
		return back
	}

	return back
}

// callApp calls R to launch an ambiorix application.
func (conf Config) callApp() (*exec.Cmd, int, error) {
	var cmd *exec.Cmd
	var port int
	rprog, err := getR()

	if err != nil {
		return cmd, port, err
	}

	script, port, err := makeCall(conf.Path)

	if err != nil {
		return cmd, port, err
	}

	cmd = exec.Command(
		rprog,
		"--no-save",
		"--slave",
		"-e",
		script,
	)

	return cmd, port, nil
}

// makeCall creates the R code used to launch the application.
func makeCall(base string) (string, int, error) {
	var script string

	path := filepath.Join(base, "app.R")
	port, err := GetFreePort()

	if err != nil {
		return script, port, err
	}

	script = "options(ambiorix.host = '0.0.0.0', ambiorix.port.force =" +
		fmt.Sprint(port) + ", shiny.port = " +
		fmt.Sprint(port) + ");source('" + path + "')"

	return script, port, nil
}
