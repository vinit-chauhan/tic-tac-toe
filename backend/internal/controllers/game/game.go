package game

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinit-chauhan/tic-tac-toe/initializers"
	"github.com/vinit-chauhan/tic-tac-toe/internal/models"
)

func fetchUserId(ctx *gin.Context) (uint, error) {
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

	ctx.JSON(http.StatusCreated, gameId)
}

func JoinGame(ctx *gin.Context) {
	gameId := ctx.Params.ByName("gameId")
	game, ok := GameState[gameId]
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": ErrGameNotFound.Error()})
		return
	}

	userId, err := fetchUserId(ctx)
	if err != nil {
		return
	}

	if game.JoinGame(int(userId)) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrGameFull.Error()})
		return
	}

	board, ok := BoardState[gameId]
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": ErrGameNotFound.Error()})
		return
	}

	ctx.JSON(http.StatusOK, board)
}

func GetGameState(ctx *gin.Context) {
	gameId := ctx.Params.ByName("gameId")
	board, ok := BoardState[gameId]
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": ErrGameNotFound.Error()})
		return
	}
	ctx.JSON(http.StatusOK, board)
}

func MakeMove(ctx *gin.Context) {
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

	if game.Winner != 0 {
		// TODO: save the game state in DB
		var user models.User
		initializers.DB.Where("ID=?", game.Winner).Find(&user)
		ctx.JSON(http.StatusOK, gin.H{"winner": user.Username})
		return
	}

	board, ok := BoardState[gameId]
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": ErrGameNotFound.Error()})
		return
	}

	ctx.JSON(http.StatusOK, board)
}
