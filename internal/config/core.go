package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

// Config structure.
type Config struct {
	Applications string `json:"applications"` // path to applications
	Port         string `json:"port"`         // port to server apps on
	User         string `json:"user"`         // user to run as
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
}

// Default configuration object.
var Default = Config{
	Applications: "/eburon/apps",
	Port:         "8080",
	User:         "",
}

// getPath retrieves the path to the configuration file from
// the environment variable.
func getPathConfig() (string, error) {
	path := os.Getenv("EBURON_CONFIG")

	if path == "" {
		path = "/eburon.json"
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

	return ioutil.WriteFile(filepath.Join(path, "eburon.json"), file, 0644)
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
