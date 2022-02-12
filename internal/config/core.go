package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type Config struct {
	Applications string `json:"applications"` // path to applications
	Port         string `json:"port"`         // port to server apps on
	User         string `json:"user"`         // user to run as
}

var Default = Config{
	Applications: "/eburon/apps",
	Port:         "8080",
	User:         "",
}

func exists(path string) bool {
	_, err := os.Stat(path)

	return !errors.Is(err, os.ErrNotExist)
}

func Read(path string) (Config, error) {
	var config Config

	isThere := exists(path)

	if !isThere {
		return config, errors.New("config file not found")
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
	isThere := exists(path)

	if isThere {
		return errors.New("config file already exists")
	}

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
