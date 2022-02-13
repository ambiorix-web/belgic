package config

import (
	"os/exec"
	"path/filepath"
)

func (conf Config) RunApps(apps []Application) []exec.Cmd {
	var cmds []exec.Cmd
	for _, app := range apps {
		cmd := exec.Command(
			"Rscript",
			"--no-save",
			"--slave",
			"-f",
			filepath.Join(conf.Applications, string(app)),
		)
		cmds = append(cmds, *cmd)
	}
	return cmds
}
