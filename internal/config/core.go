package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// Config structure.
type Config struct {
	Path     string `json:"path"` // path to application
	Port     string `json:"port"` // port to server apps on
	Backends string `json:"background"`
}

// Default configuration object.
var Default = Config{
	Path:     "/belgic",
	Port:     "8080",
	Backends: "max",
}

// getPath retrieves the path to the configuration file from
// the environment variable.
func getPathConfig() (string, error) {
	path := os.Getenv("BELGIC_CONFIG")

	if path == "" {
		return "", errors.New("BELGIC_CONFIG environment variable not set")
	}

	return path, nil
}

// Read the configuration file.
func Read() (Config, error) {
	var config Config

	path, err := getPathConfig()

	if err != nil {
		return config, err
	}

	file, err := os.Open(path)

	if err != nil {
		return config, err
	}

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		return config, err
	}

	err = json.Unmarshal(byteValue, &config)

	return config, err
}

// Create the default configuration file.
func Create(path string) error {
	file, err := json.MarshalIndent(Default, "", " ")

	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(path, "belgic.json"), file, 0644)
}

// CheckConfigPath Checks that the path to create the configuration file
// is correct.
func CheckConfigPath(path string) error {
	found, err := regexp.MatchString("\\.json$|\\.config$", path)

	if err != nil {
		return errors.New("could not check path")
	}

	if found {
		return errors.New("specify a path to a directory, not a path to a file")
	}

	return nil
}
