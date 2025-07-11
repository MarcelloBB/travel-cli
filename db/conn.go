package db

import (
	"fmt"
	"travel-cli/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	var err error

	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil, err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(&model.Workspace{}, &model.Collection{}, &model.Request{})
	if err != nil {
		fmt.Println("Error running automigration:", err)
		return nil, err
	}

	return DB, nil
}
