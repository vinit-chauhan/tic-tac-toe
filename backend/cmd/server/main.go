package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/vinit-chauhan/tic-tac-toe/config"
	"github.com/vinit-chauhan/tic-tac-toe/initializers"
	"github.com/vinit-chauhan/tic-tac-toe/internal/controllers/auth"
	"github.com/vinit-chauhan/tic-tac-toe/internal/controllers/user"
	"github.com/vinit-chauhan/tic-tac-toe/internal/middleware"
	"github.com/vinit-chauhan/tic-tac-toe/utils/logger"
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

func Todo(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "implementation missing"})
}

func main() {
	logger.Info("starting the server", "main")

	logger.Info("connecting to Database", "main")
	err := initializers.ConnectDB(conf)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(logger.LogsDir + "/gin.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = out

	if os.Getenv("SECRET") == "" {
		panic("env(SECRET) not set")
	}

	r := gin.Default()

	// Auth endpoints
	r.POST("/auth/signup/", user.CreateUser)
	r.POST("/auth/login/", auth.Login)
	r.POST("/auth/signout/", auth.SignOut)

	// User routes
	r.GET("/user/profile", middleware.CheckAuth, user.GetUserInfo)

	addr := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)

	r.Run(addr)
}
