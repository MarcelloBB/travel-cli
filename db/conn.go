package db

import (
	"fmt"
	"travel-cli/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to database.db")
	fmt.Println("Running automigration...")

	err = db.AutoMigrate(&model.Workspace{})
	if err != nil {
		fmt.Println("Error running automigration:", err)
		panic(err)
	}

	return db, nil
}
