package database

import (
	"fmt"

	"Tugas-kedua-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
    _ "github.com/go-sql-driver/mysql"
)

var Db *gorm.DB

// InitDB initializes the database connection and sets up the database schema
func InitDB() {
    var err error
    dataSourceName := "WindowsX:12345@tcp(localhost:3306)/orders_db?parseTime=True"
    Db, err = gorm.Open("mysql", dataSourceName)

    if err != nil {
        fmt.Println(err)
        panic("failed to connect database")
    }


    // Migration to create tables for Order and Item schema
    Db.AutoMigrate(&models.Order{}, &models.Item{})
}
