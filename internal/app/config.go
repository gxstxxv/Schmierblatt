package app

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		AssetsPath string `yaml:"assets_path"`
		LogFile    string `yaml:"log_file"`
		LogLevel   string `yaml:"log_level"`
	} `yaml:"app"`
}

var config Config

func LoadConfig() error {
	f, err := os.Open("config.yaml")
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}

	return nil
}
