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
	Code string
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
	db := app.GetDB()
	db.AutoMigrate(&User{})
	//list all users
	var users []User
	db.Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}

	user := User{Name: "John", Age: 25}
	db.Create(&user)
	//print
}
