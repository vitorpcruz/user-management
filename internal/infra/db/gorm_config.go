package db

import (
	"log"

	"github.com/vitorpcruz/sso/config"
	entity "github.com/vitorpcruz/sso/internal/domain/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := config.MySqlDSN()
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Customer{},
		&entity.System{},
	)
	if err != nil {
		log.Fatal("error on migrate entities", err)
	}

	return db
}
