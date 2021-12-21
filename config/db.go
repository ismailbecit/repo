package config

import (
	"repo/modal"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connet() *gorm.DB {
	dsn := "root:aea76026@(127.0.0.1:3306)/repo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Veri Tabanına Bağlanılamadı")
	}
	db.AutoMigrate(&modal.Category{})

	return db
}
