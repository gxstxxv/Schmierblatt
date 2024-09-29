package app

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		AssetsPath         string `yaml:"assets_path"`
		LogFile            string `yaml:"log_file"`
		LogLevel           string `yaml:"log_level"`
		LastOpenedFilePath string `yaml:"last_opened_file_path"`
	} `yaml:"app"`
}

var config Config

func LoadConfig(configPath string) error {
	f, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}

	// Convert relative paths to absolute paths
	exeDir := filepath.Dir(configPath)
	config.App.AssetsPath = filepath.Join(exeDir, config.App.AssetsPath)
	config.App.LogFile = filepath.Join(exeDir, config.App.LogFile)
	config.App.LastOpenedFilePath = filepath.Join(exeDir, config.App.LastOpenedFilePath)

	return nil
}
