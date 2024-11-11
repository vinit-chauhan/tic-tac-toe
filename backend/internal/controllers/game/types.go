package game

import "time"

var BoardState = make(map[string]*Board)
var GameState = make(map[string]*Game)

type Game struct {
	ID         string
	Player1    int
	Player2    int
	GameTime   int64
	Winner     int
	LastMoveBy int
}

type GameMove struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

func (g *Game) JoinGame(player int) error {
	if g.Player2 != 0 {
		return ErrGameFull
	}

	if g.Player1 == player {
		return ErrCannotJoinOwnGame
	}

	g.Player2 = player
	g.GameTime = time.Now().Unix()

	return nil
}

func (g *Game) MakeMove(player int, move GameMove) error {
	moveType := 0

	if g.LastMoveBy == player {
		return ErrWaitForOpponent
	}

	if player == g.Player1 {
		moveType = 1
	} else if player == g.Player2 {
		moveType = -1
	}

	if moveType == 0 {
		return ErrPlayerNotPermitted
	}

	board := BoardState[g.ID]

	if board[move.Row*3+move.Column] != 0 {
		return ErrInvalidMove
	}

	board[move.Row*3+move.Column] = moveType

	winner := board.CheckWinner()
	if winner == 0 {
		if board.IsFull() {
			g.Winner = -1
		}
	} else {
		if winner == 1 {
			g.Winner = g.Player1
		} else {
			g.Winner = g.Player2
		}
	}

	g.LastMoveBy = player
	return nil
}

type Board [9]int

func NewBoard() Board {
	return Board{0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func (b *Board) IsFull() bool {
	for _, v := range b {
		if v == 0 {
			return false
		}
	}
	return true
}

func (b *Board) CheckWinner() int {
	for i := 0; i < 3; i++ {
		if b[i] == b[i+3] && b[i] == b[i+6] && b[i] != 0 {
			return b[i]
		}
		if b[i*3] == b[i*3+1] && b[i*3] == b[i*3+2] && b[i*3] != 0 {
			return b[i*3]
		}
	}
	if b[0] == b[4] && b[0] == b[8] && b[0] != 0 {
		return b[0]
	}
	if b[2] == b[4] && b[2] == b[6] && b[2] != 0 {
		return b[2]
	}
	return 0
}
