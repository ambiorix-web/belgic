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

type Config struct {
	Applications string `json:"applications"` // path to applications
	Port         string `json:"port"`         // port to server apps on
	User         string `json:"user"`         // user to run as
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
}

var Default = Config{
	Applications: "/eburon/apps",
	Port:         "8080",
	User:         "",
}

func getPath() (string, error) {
	path := os.Getenv("EBURON_CONFIG")

	if path == "" {
		path = "/eburon.json"
	}

	return path, nil
}

func Read() (Config, error) {
	var config Config

	path, err := getPath()

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

func Create(path string) error {
	file, err := json.MarshalIndent(Default, "", " ")

	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(path, "eburon.json"), file, 0644)
}

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
