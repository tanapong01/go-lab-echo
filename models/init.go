package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SQLiteDB : Centralize database instance connection
var Db *gorm.DB

// InitialDatabase : Initial database function to connect with DBMS
func InitialDatabase() error {

	var err error
	dsn := "host=localhost user=postgres password=aA250544 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	autoMigrateDatabase()

	fmt.Println("Database connection successfully...")
	return nil
}

func autoMigrateDatabase() {
	Db.AutoMigrate(&StudentScore{})
}
