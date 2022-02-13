package config

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
)

func getRscript() (string, error) {
	var p string

	p, err := exec.LookPath("R")

	if err != nil {
		return p, errors.New("could not locate R installation")
	}

	return p, nil
}

type RCommand struct {
	Application Application
	Err         error
	Cmd         *exec.Cmd
	Port        int
}

type RCommands []RCommand

func (conf Config) RunApps(apps []Application) (RCommands, error) {
	var cmds RCommands

	for _, app := range apps {
		cmd := conf.runApp(app)
		cmds = append(cmds, cmd)
	}
	return cmds, nil
}

func (conf Config) runApp(app Application) RCommand {
	var rcmd RCommand
	rcmd.Application = app

	cmd, port, err := conf.callApp(app)

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

func (conf Config) callApp(app Application) (*exec.Cmd, int, error) {
	var cmd *exec.Cmd
	var port int
	rprog, err := getRscript()

	if err != nil {
		return cmd, port, err
	}

	script, port, err := makeCall(conf.Applications, app)

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

func makeCall(base string, app Application) (string, int, error) {
	var script string

	path := filepath.Join(base, string(app))
	port, err := GetFreePort()

	if err != nil {
		return script, port, err
	}

	script = "options(ambiorix.host = '0.0.0.0', ambiorix.port ='" + fmt.Sprint(port) + "');source('" + path + "')"

	return script, port, nil
}
