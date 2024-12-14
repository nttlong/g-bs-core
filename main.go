package main

import (
	"fmt"

	application "gnol.hrm.core/pkg/structs/app"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	//create new user

	app := application.App{}
	err := app.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(app)
	//print
}
