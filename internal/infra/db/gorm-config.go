package db

import (
	"log"

	entity "github.com/vitorpcruz/sso/internal/domain/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(
		mysql.Open("user:pass1234567@tcp(127.0.0.1:3306)/sso?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Customer{},
		&entity.System{},
		&entity.System{},
		&entity.SystemUser{},
		&entity.SystemUserRole{},
	)
	if err != nil {
		log.Fatal("error on migrate entities", err)
	}

	return db
}
