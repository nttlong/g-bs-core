package structs

import (
	"fmt"
	"os"
	"path/filepath"

	"encoding/json"

	"gopkg.in/yaml.v2"
)

type Config struct {
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
	ConnectionString       string
	MasterConnectionString string
}

func (c *Config) Load(appDir string) error {
	configFilePath := filepath.Join(appDir, "config", "app.yml")
	// Check if the config file exists
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return fmt.Errorf("config file does not exist: %s", configFilePath)
	}

	// Read the config file using os.ReadFile
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}
	// Unmarshal the YAML data into the Config struct
	if err := yaml.Unmarshal(data, c); err != nil {
		return fmt.Errorf("error unmarshalling config file: %w", err)
	}
	c.ConnectionString = c.getConnectionString()
	c.MasterConnectionString = c.getConnectionStringNoDatabase()
	return nil
}

func (p *Config) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshalling to JSON: %v", err)
	}

	return string(jsonData)
}

func (p *Config) getConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", p.DB.Username, p.DB.Password, p.DB.Host, p.DB.Database)
}
func (p *Config) getConnectionStringNoDatabase() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8mb4&parseTime=True&loc=Local", p.DB.Username, p.DB.Password, p.DB.Host)
}
