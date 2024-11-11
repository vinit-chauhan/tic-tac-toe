package game

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrParsingUser  = errors.New("error parsing user")

	ErrGameNotFound = errors.New("game not found")
	ErrGameFull     = errors.New("game already full")
	ErrGameFinished = errors.New("game already finished")

	ErrInvalidMove        = errors.New("invalid move")
	ErrCannotJoinOwnGame  = errors.New("cannot join own game")
	ErrWaitForOpponent    = errors.New("wait for opponent")
	ErrPlayerNotPermitted = errors.New("player not permitted")
)
