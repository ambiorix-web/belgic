package config

import (
	"fmt"
	"os"
)

type Applications string

func (config Config) ListApps() ([]Applications, error) {
	var apps []Applications
	dirs, err := os.ReadDir(config.Applications)

	if err != nil {
		return apps, err
	}

	for _, file := range dirs {
		fmt.Println(file.IsDir())
	}

	return apps, nil
}
