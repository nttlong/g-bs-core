package main

import (
	"fmt"

	app_config "gnol.hrm.core/app_config"
	app_db "gnol.hrm.core/app_db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	appPath, err := app_config.GetAppDir()

	if err != nil {
		panic(err)
	}
	db, err := app_db.Connect(appPath, "hrm")
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
}
