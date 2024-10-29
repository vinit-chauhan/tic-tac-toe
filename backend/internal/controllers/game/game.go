package game

import "github.com/gin-gonic/gin"

func StartNewGame(ctx *gin.Context) {}

func GetGameState(ctx *gin.Context) {}

func MakeMove(ctx *gin.Context) {
	checkIsCompleted(Board{})
}

func checkIsCompleted(state Board) {}
