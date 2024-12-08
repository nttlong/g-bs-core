package app_db

import (
	"errors"
	"fmt"
	"sync"

	app_config "gnol.hrm.core/app_config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once

func Connect(appApth string, dbName string) (*gorm.DB, error) {
	config, err := app_config.LoadConfig(appApth)
	if err != nil {
		return nil, err
	}

	// Validate dbName
	if dbName == "" {
		return nil, errors.New("dbName cannot be empty")
	}

	host := config.DB.Host
	password := config.DB.Password
	user := config.DB.Username
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8mb4&parseTime=True&loc=Local", user, password, host)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	once.Do(func() {
		sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
		ret := db.Exec(sql)
		if ret.Error != nil {
			err = ret.Error
		}
	})

	if err != nil {
		return nil, err
	}
	return db, nil
}
