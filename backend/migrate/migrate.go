package main

import (
	"fmt"

	"github.com/vinit-chauhan/tic-tac-toe/config"
	"github.com/vinit-chauhan/tic-tac-toe/internal"
	"github.com/vinit-chauhan/tic-tac-toe/internal/models"
)

func init() {

	conf, err := config.Load("config.yml")
	if err != nil {
		panic(err)
	}

	fmt.Println("Config File Loaded Successfully")

	err = internal.ConnectDB(conf)
	if err != nil {
		panic(err)
	}

	fmt.Println("DB initialized successfully")
}

func main() {
	internal.DB.AutoMigrate(&models.User{})

	fmt.Println("Models migrated successfully")
}
