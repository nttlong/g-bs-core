package app

import (
	"encoding/json"
	fmt "fmt"
	"path/filepath"
	"runtime"

	config "gnol.hrm.core/pkg/structs/config"
)

type App struct {
	Name    string
	Version string
	AppPath string
	Config  *config.Config
}

func (c *App) Start() error {
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
	c.Config = &config.Config{}
	err = c.Config.Load(c.AppPath)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		return fmt.Errorf("Error loading configuration: %w", err)
	}
	return nil
	// do something here
}
func (p *App) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshalling to JSON: %v", err)
	}

	return string(jsonData)
}
