package internal

import (
	"fmt"

	"github.com/vinit-chauhan/tic-tac-toe/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(conf config.Config) error {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		conf.Database.Host,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.DBName,
		conf.Database.Port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}
