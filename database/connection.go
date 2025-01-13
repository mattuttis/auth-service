package database

import (
	"github.com/mattuttis/auth-service/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	dsn := "zdw:zdw@tcp(127.0.0.1:3306)/zdw-auth-service"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db
	db.AutoMigrate(&models.User{})

	return db, nil
}
