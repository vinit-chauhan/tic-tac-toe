package middleware

import "github.com/gin-gonic/gin"

func CheckAuth(ctx *gin.Context) {

	ctx.Next()
}
