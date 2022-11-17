package app

import (
	"github.com/giriaditya/test-gin/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() *gorm.DB {
	dsn := "root:root@tcp(localhost:3306)/test_gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	db.AutoMigrate(&repository.Book{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	return db
}
