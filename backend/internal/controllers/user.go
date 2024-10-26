package controllers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/vinit-chauhan/tic-tac-toe/initializers"
	"github.com/vinit-chauhan/tic-tac-toe/internal/models"
)

func GetUserInfo(ctx *gin.Context) {
	user, ok := ctx.Get("currentUser")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func CreateUser(ctx *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email"`
	}

	if err := ctx.ShouldBindBodyWithJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var u models.User
	initializers.DB.Where("username=?", body.Username).Find(&u)
	if u.ID != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user already exist"})
		return
	}

	pHash, err := bcrypt.GenerateFromPassword([]byte(body.Password+body.Username), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: body.Username,
		Password: string(pHash),
		Email:    body.Email,
	}

	res := initializers.DB.Create(&user)
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": res.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": user.ID,
	})
}
