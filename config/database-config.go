package config

import (
	"github.com/fplaraujo/agenda/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func SetupDatabaseConnection() {
	dsn := "root:root@tcp(localhost:3306)/agenda?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&model.Contact{})
	Database = db
}

func CloseDatabaseConnection() {
	dbSQL, err := Database.DB()
	if err != nil {
		panic("Failed to close connection to database!")
	}
	dbSQL.Close()
}
