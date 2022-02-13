package config

import (
	"os"
)

// ApplicationName is the name of an application, also its path.
type ApplicationName string

// ApplicationNames is an array of names of applications, also their paths.
type ApplicationNames []ApplicationName

// ListApps lists the applications present in the directory.
func (config Config) ListApps() (ApplicationNames, error) {
	var apps ApplicationNames
	dirs, err := os.ReadDir(config.Applications)

	if err != nil {
		return apps, err
	}

	for _, file := range dirs {
		if !file.IsDir() {
			continue
		}

		apps = append(apps, ApplicationName(file.Name()))
	}

	return apps, nil
}
