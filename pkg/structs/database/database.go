package database

import (
	"fmt"
	"sync"

	db_config "gnol.hrm.core/pkg/structs/config"
	mysqlutils "gnol.hrm.core/pkg/structs/mysqlutils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// create map check is database has check
var (
	dbMap   = make(map[string]*gorm.DB)
	onceMap = make(map[string]*sync.Once) // To hold sync.Once for each database
	mutex   sync.Mutex                    // Mutex to protect map access
)

// create database if not exit
func createDatabasefNotExist(config db_config.Config) (*gorm.DB, error) {
	master_db, ex := gorm.Open(mysql.Open(config.MasterConnectionString), &gorm.Config{})
	if ex != nil {
		panic(ex)
	}
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.DB.Database)
	ret := master_db.Exec(sql)
	if ret.Error != nil {
		return nil, ret.Error

	}
	mDb, ex := master_db.DB()
	if ex != nil {
		return nil, ex
	}
	master_db.Commit()
	mDb.Close()
	db, ex := gorm.Open(mysql.Open(config.ConnectionString), &gorm.Config{})
	if ex != nil {
		return nil, ex
	}

	return db, nil
}

func GetDB(config db_config.Config) (*gorm.DB, error) {
	// check if config.DB.Database exists
	dbName := config.DB.Database

	// Ensure the mutex is locked to avoid race conditions
	mutex.Lock()
	defer mutex.Unlock() // Ensure the mutex is unlocked when the function exits

	// Check in dbMap if the database already exists
	if db, ok := dbMap[dbName]; ok {
		return db, nil
	}

	// Create a new once instance to ensure the database is only created once
	once, ok := onceMap[dbName]
	if !ok {
		once = &sync.Once{}
		onceMap[dbName] = once
	}

	// Use the once to ensure only one creation attempt
	var ret_db *gorm.DB
	var retErr error = nil
	once.Do(func() {
		ret_db, retErr = mysqlutils.CreateDatabasefNotExist(config)
		if ret_db != nil { // Ensure that the creation was successful
			dbMap[dbName] = ret_db
		}
	})
	if retErr != nil {
		return nil, retErr
	}
	// Double-check if the database was created after using once
	if ret_db == nil {
		return nil, fmt.Errorf("failed to create database: %s", dbName)
	}

	return dbMap[dbName], nil
}
