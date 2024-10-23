package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vinit-chauhan/tic-tac-toe/config"
	"github.com/vinit-chauhan/tic-tac-toe/internal"
	"github.com/vinit-chauhan/tic-tac-toe/pkg/logger"
)

var conf config.Config

func init() {
	var err error

	logger.SetLogLevel(logger.LevelDebug)

	logger.Debug("initializing the server", "init")
	logger.Debug("loading config file", "init")

	conf, err = config.Load("config.yml")
	if err != nil {
		logger.Panic("error loading config file", "init", err)
	}

	logger.Info("config loaded successfully", "init")
}

func main() {
	logger.Info("starting the server", "main")

	logger.Info("connecting to Database", "main")
	err := internal.ConnectDB(conf)
	if err != nil {
		panic(err)
	}

	out, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = out

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello",
		})
	})

	addr := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)

	r.Run(addr)
}
