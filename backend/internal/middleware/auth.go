package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vinit-chauhan/tic-tac-toe/initializers"
	"github.com/vinit-chauhan/tic-tac-toe/internal/models"
)

func CheckAuth(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header not set"})
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		return
	}

	token, err := jwt.Parse(authToken[1], func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		return
	}

	var user models.User
	initializers.DB.Where("ID=?", claims["id"]).Find(&user)

	if user.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user unauthorized"})
	}

	ctx.Set("currentUser", user)

	ctx.Next()
}
