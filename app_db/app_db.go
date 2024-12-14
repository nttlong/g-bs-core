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
	err = nil

	once.Do(func() {
		master_db, ex := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
		if err != nil {
			err = fmt.Errorf("failed to connect database: %w", ex)
			return
		}
		sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
		ret := master_db.Exec(sql)
		if ret.Error != nil {
			err = ret.Error
		}
		mDb, ex := master_db.DB()
		if err != nil {
			err = fmt.Errorf("failed to get master database connection: %w", ex)
			return
		}
		master_db.Commit()
		mDb.Close()

	})

	if err != nil {
		return nil, err
	}
	db_connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbName)
	db, err := gorm.Open(mysql.Open(db_connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	return db, nil
}

// lis all databases in the server
func ListDatabases(appApth string) ([]string, error) {
	config, err := app_config.LoadConfig(appApth)
	if err != nil {
		return nil, err
	}
	host := config.DB.Host
	password := config.DB.Password
	user := config.DB.Username
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8mb4&parseTime=True&loc=Local", user, password, host)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	var dbs []string
	sql := "SHOW DATABASES"
	rows, err := db.Raw(sql).Rows()
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var dbName string
		err = rows.Scan(&dbName)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		dbs = append(dbs, dbName)
	}
	return dbs, nil
}
