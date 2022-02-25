package config

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"sync"
)

// RCommand represents a single R command.
type Backend struct {
	Port  int
	Path  string
	Rpath string
	Mu    *sync.RWMutex
}

// getR retrieves the full path to the R installation.
func getR() (string, error) {
	var p string

	p, err := exec.LookPath("R")

	if err != nil {
		return p, errors.New("could not locate R installation")
	}

	return p, nil
}

// runApp run a single application.
func (back *Backend) RunApp(stdout chan string) error {
	err := back.callApp(stdout)

	if err != nil {
		return err
	}

	return nil
}

// callApp calls R to launch an ambiorix application.
func (back *Backend) callApp(stdout chan string) error {
	rprog, err := getR()

	if err != nil {
		return err
	}

	script, port, err := makeCall(back.Rpath)

	if err != nil {
		return err
	}

	back.Port = port
	back.Path = "http://localhost:" + strconv.Itoa(port)

	go back.ExecuteCommand(rprog, script, stdout)

	return nil
}

func (back *Backend) ExecuteCommand(rprog, script string, stdout chan string) {
	cmd := exec.Command(
		rprog,
		"--no-save",
		"--slave",
		"-e",
		script,
	)

	stder, err := cmd.CombinedOutput()

	if err != nil {
		stdout <- err.Error()
		return
	}

	stdout <- string(stder)
}

// makeCall creates the R code used to launch the application.
func makeCall(base string) (string, int, error) {
	var script string

	port, err := GetFreePort()

	if err != nil {
		return script, port, err
	}

	script = "setwd('" + base + "');options(ambiorix.host = '0.0.0.0', ambiorix.port.force =" +
		fmt.Sprint(port) + ", shiny.port = " +
		fmt.Sprint(port) + ");source('app.R')"

	return script, port, nil
}
