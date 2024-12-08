package main

import (
	"fmt"

	app_config "gnol.hrm.core/app_config"
)

func main() {
	app_dir, err := app_config.GetAppDir()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(app_dir)
	config, err := app_config.LoadConfig(app_dir)
	if err != nil {
		panic(err)

	}
	fmt.Println(config.String())
}
