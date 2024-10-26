package main

import (
	"fmt"

	"github.com/vinit-chauhan/tic-tac-toe/config"
	"github.com/vinit-chauhan/tic-tac-toe/initializers"
	"github.com/vinit-chauhan/tic-tac-toe/internal/models"
)

func init() {

	conf, err := config.Load("config.yml")
	if err != nil {
		panic(err)
	}

	fmt.Println("Config File Loaded Successfully")

	err = initializers.ConnectDB(conf)
	if err != nil {
		panic(err)
	}

	fmt.Println("DB initialized successfully")
}

func main() {
	if err := initializers.DB.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

	fmt.Println("Models migrated successfully")
}