package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vinit-chauhan/tic-tac-toe/config"
	"github.com/vinit-chauhan/tic-tac-toe/pkg/logger"
)

var conf config.Config

func init() {
	var err error

	logger.SetLogLevel(logger.LevelDebug)

	logger.Debug("[init] initializing the server")
	logger.Debug("[init] loading config file")

	conf, err = config.Load("config.yml")
	if err != nil {
		logger.Panic("[init] error loading config file", err)
	}

	logger.Info("[init] config loaded successfully")
}

func main() {
	logger.Info("[main] Starting the server")

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
