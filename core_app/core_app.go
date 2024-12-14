package core_app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

// struct for app configuration
type AppConfiguration struct {
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

// declare a struct of type CoreApp
type CoreApp struct {
	Name      string
	Version   string
	AppPath   string
	AppConfig AppConfiguration
}

// load configuration from file

func (c *CoreApp) LoadConfig() error {
	// Use CoreApp.AppPath to get the path of the config file
	configFilePath := filepath.Join(c.AppPath, "config", "app.yml")

	// Check if the config file exists
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return fmt.Errorf("config file does not exist: %s", configFilePath)
	}

	// Read the config file
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	// Initialize app_config with default values if necessary
	appConfig := AppConfiguration{}

	// Unmarshal the YAML data into the AppConfiguration struct
	if err := yaml.Unmarshal(data, &appConfig); err != nil {
		return fmt.Errorf("error unmarshalling config file: %w", err)
	}

	// Here you can validate appConfig fields and set default values if needed

	// Assign the loaded config to CoreApp's AppConfig
	c.AppConfig = appConfig
	return nil
}

// create startup of application via CoreApp struct
func (c *CoreApp) Start() error {
	// calculate and set the AppPath
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf("Cannot get caller information")
	}

	absPath, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("Error getting absolute path: %w", err)
	}
	c.AppPath = filepath.Dir(absPath)
	err = c.LoadConfig()
	if err != nil {
		return fmt.Errorf("Error loading configuration: %w", err)
	}
	return nil
	// do something here
}
func (p *CoreApp) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshalling to JSON: %v", err)
	}

	return fmt.Sprintf(string(jsonData))
}
