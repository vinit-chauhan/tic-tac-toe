package controllers

import "github.com/gin-gonic/gin"

func GetUserInfo(ctx *gin.Context) {
	var body struct {
		UserId string `json:"user_id"`
	}

	ctx.BindJSON(&body)
}
