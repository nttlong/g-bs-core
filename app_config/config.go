package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	once          sync.Once
	app_config    AppConfig
	app_configErr error
)

type AppConfig struct {
	Web struct {
		Bind string `yaml:"bind"`
	}
	DB struct {
		DbType   string `yaml:"DbType"`
		Host     string `yaml:"Host"`
		Database string `yaml:"Database"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
	}
}

// This function will get current excuable directory of app
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

// Load config only one run even call many time
func LoadConfig(locate string) (AppConfig, error) {
	once.Do(func() {
		config_file_path := filepath.Join(locate, "config", "app.yml")
		data, err := ioutil.ReadFile(config_file_path)
		if err != nil {
			app_configErr = err
			return
		}
		err = yaml.Unmarshal(data, &app_config)
		if err != nil {
			app_configErr = err
			return
		}
	})

	return app_config, app_configErr

}
func (p *AppConfig) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshalling to JSON: %v", err)
	}

	return fmt.Sprintf(string(jsonData))
}
