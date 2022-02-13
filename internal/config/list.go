package config

import (
	"os"
)

type Application string

func (config Config) ListApps() ([]Application, error) {
	var apps []Application
	dirs, err := os.ReadDir(config.Applications)

	if err != nil {
		return apps, err
	}

	for _, file := range dirs {
		if !file.IsDir() {
			continue
		}

		apps = append(apps, Application(file.Name()))
	}

	return apps, nil
}
