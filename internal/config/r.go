package config

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
)

// getR retrieves the full path to the R installation.
func getR() (string, error) {
	var p string

	p, err := exec.LookPath("R")

	if err != nil {
		return p, errors.New("could not locate R installation")
	}

	return p, nil
}

// RCommand represents a single R command.
type RCommand struct {
	Err  error
	Cmd  *exec.Cmd
	Port int
}

// RCommands represents an array of R commands.
type RCommands []RCommand

// RunApps runs all the applications found in the directory.
func (conf Config) RunApps() (RCommands, error) {
	var cmds RCommands
	ncpus, err := strconv.Atoi(conf.Background)

	// error we assume it was set to max
	// default to max CPUs
	if err != nil {
		ncpus = runtime.NumCPU()
	}

	for i := 0; i < ncpus; i++ {
		cmd := conf.runApp()
		cmds = append(cmds, cmd)
	}

	return cmds, nil
}

// runApp run a single application.
func (conf Config) runApp() RCommand {
	var rcmd RCommand

	cmd, port, err := conf.callApp()

	if err != nil {
		rcmd.Err = err
		return rcmd
	}
	rcmd.Port = port
	rcmd.Cmd = cmd
	rcmd.Err = rcmd.Cmd.Start()

	if rcmd.Err != nil {
		return rcmd
	}

	return rcmd
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

	script = "options(ambiorix.host = '0.0.0.0', ambiorix.port.force ='" + fmt.Sprint(port) + "');source('" + path + "')"

	return script, port, nil
}
