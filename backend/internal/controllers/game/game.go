package game

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinit-chauhan/tic-tac-toe/internal/database"
	"github.com/vinit-chauhan/tic-tac-toe/internal/models"
	"github.com/vinit-chauhan/tic-tac-toe/metrics"
	"github.com/vinit-chauhan/tic-tac-toe/utils/logger"
)

func fetchUserId(ctx *gin.Context) (uint, error) {
	metrics.HttpRequestsTotal.WithLabelValues(ctx.Request.URL.Path).Inc()
	details, ok := ctx.Get("currentUserId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ErrUnauthorized.Error()})
		return 0, ErrUnauthorized
	}

	userId, ok := details.(uint)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrParsingUser.Error()})
		return 0, ErrParsingUser
	}

	return userId, nil
}

func StartNewGame(ctx *gin.Context) {
	metrics.HttpRequestsTotal.WithLabelValues(ctx.Request.URL.Path).Inc()

	userId, err := fetchUserId(ctx)
	if err != nil {
		return
	}

	// create a new game
	board := NewBoard()
	gameId := strconv.Itoa(int(userId ^ 0xdeadbeef))
	BoardState[gameId] = &board
	GameState[gameId] = &Game{
		ID:      gameId,
		Player1: int(userId),
	}

	// TODO: save the game state in redis
	// initializers.RedisClient.Set(context.Background(), gameId, board, 0)

	ctx.JSON(http.StatusCreated, gin.H{"game_id": gameId})
}

func JoinGame(ctx *gin.Context) {
	metrics.HttpRequestsTotal.WithLabelValues(ctx.Request.URL.Path).Inc()

	userId, err := fetchUserId(ctx)
	if err != nil {
		return
	}

	gameId := ctx.Params.ByName("gameId")
	game, ok := GameState[gameId]
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": ErrGameNotFound.Error()})
		return
	}

	logger.Debug(fmt.Sprintf("user:{%d} requested to join game:{%s}", userId, gameId), "JoinGame")

	if game.JoinGame(int(userId)) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrGameFull.Error()})
		return
	}

	board, ok := BoardState[gameId]
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": ErrGameNotFound.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"board": board})
}

func GetGameState(ctx *gin.Context) {
	metrics.HttpRequestsTotal.WithLabelValues(ctx.Request.URL.Path).Inc()
	gameId := ctx.Params.ByName("gameId")
	board, ok := BoardState[gameId]
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": ErrGameNotFound.Error()})
		return
	}

	game := GameState[gameId]
	if game.Winner != 0 {
		var user models.User
		database.DB.Where("ID=?", game.Winner).Find(&user)
		ctx.JSON(http.StatusOK, gin.H{"winner": user.Username, "board": board})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"board": board})
}

func MakeMove(ctx *gin.Context) {
	metrics.HttpRequestsTotal.WithLabelValues(ctx.Request.URL.Path).Inc()
	userId, err := fetchUserId(ctx)
	if err != nil {
		return
	}

	gameId := ctx.Params.ByName("gameId")
	game, ok := GameState[gameId]
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": ErrGameNotFound.Error()})
		return
	}

	if game.Winner != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrGameFinished.Error()})
		return
	}

	move := GameMove{}

	if err := ctx.ShouldBindBodyWithJSON(&move); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := game.MakeMove(int(userId), move); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	board, ok := BoardState[gameId]
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": ErrGameNotFound.Error()})
		return
	}

	if game.Winner != 0 {
		// TODO: save the game state in DB
		var user models.User
		database.DB.Where("ID=?", game.Winner).Find(&user)
		ctx.JSON(http.StatusOK, gin.H{"winner": user.Username, "board": board})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"board": board})
}
