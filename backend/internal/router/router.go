package router

import (
	"github.com/gin-gonic/gin"

	"github.com/vinit-chauhan/tic-tac-toe/internal/controllers/auth"
	"github.com/vinit-chauhan/tic-tac-toe/internal/controllers/user"
	"github.com/vinit-chauhan/tic-tac-toe/internal/middleware"
	"github.com/vinit-chauhan/tic-tac-toe/utils/logger"
)

func SetRoutes(r *gin.Engine) {

	logger.Info("setting up CORS middleware", "main")
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Auth endpoints
	r.POST("/auth/login", auth.Login)
	r.POST("/auth/logout", auth.SignOut)

	// User routes
	r.POST("/users", user.CreateUser)
	r.GET("/profile", middleware.CheckAuth, user.GetUserInfo)

}
