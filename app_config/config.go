package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Server struct {
		Port int `yaml:"port"`
	}
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
}

func GetAppDir() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("Cannot get caller information")
	}

	absPath, err := filepath.Abs(filename)
	if err != nil {
		return "", fmt.Errorf("Error getting absolute path: %w", err)
	}

	return filepath.Dir(absPath), nil
}
func LoadConfig(locate string) (AppConfig, error) {
	var config AppConfig

	config_file_path := filepath.Join(locate, "config", "app.yml")
	data, err := ioutil.ReadFile(config_file_path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
func (p *AppConfig) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshalling to JSON: %v", err)
	}

	return fmt.Sprintf(string(jsonData))
}
