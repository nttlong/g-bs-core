package mysqlutils

import (
	"fmt"

	db_config "gnol.hrm.core/pkg/structs/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDatabasefNotExist(config db_config.Config) (*gorm.DB, error) {
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
