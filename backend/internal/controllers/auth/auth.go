package auth

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vinit-chauhan/tic-tac-toe/internal/database"
	"github.com/vinit-chauhan/tic-tac-toe/internal/models"
	"github.com/vinit-chauhan/tic-tac-toe/metrics"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	metrics.HttpRequestsTotal.WithLabelValues(ctx.Request.URL.Path).Inc()
	var body struct {
		Username string `binding:"required"`
		Password string `binding:"required"`
	}

	if err := ctx.ShouldBindBodyWithJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var foundUser models.User
	res := database.DB.Where("username=?", body.Username).Find(&foundUser)
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": res.Error})
		return
	}

	if foundUser.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(body.Password+body.Username)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	exp := time.Now().Add(24 * time.Hour).Unix()
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(foundUser.ID)),
		"exp": exp,
		"iat": time.Now().Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error generating token"})
		return
	}

	ctx.SetCookie("Authorization", token, int(exp-time.Now().Unix()), "/", "", true, true)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "logged in successfully",
		"token":   token,
	})
}

func SignOut(ctx *gin.Context) {
	metrics.HttpRequestsTotal.WithLabelValues(ctx.Request.URL.Path).Inc()
	ctx.SetCookie("Authorization", "", int(time.Now().Add(-1*time.Hour).Unix()), "/", "", true, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully signed out"})
}
