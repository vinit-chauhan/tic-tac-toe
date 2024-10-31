package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vinit-chauhan/tic-tac-toe/internal/controllers/auth"
	"github.com/vinit-chauhan/tic-tac-toe/internal/controllers/user"
	"github.com/vinit-chauhan/tic-tac-toe/internal/middleware"
)

func SetRoutes(r *gin.Engine) {

	// Auth endpoints
	r.POST("/auth/login", auth.Login)
	r.POST("/auth/logout", auth.SignOut)

	// User routes
	r.POST("/users", user.CreateUser)
	r.GET("/profile", middleware.CheckAuth, user.GetUserInfo)

}
