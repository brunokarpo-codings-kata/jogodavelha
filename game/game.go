package game

import (
	"errors"
	"github.com/brunokarpo-codings-kata/jogodavelha/player"
)

const indexTurnPlayer = 0
const indexWaitingPlayer = 1

var NoWinnerError = errors.New("no winner yet")

type GameAPI interface {
	Mark(x, y int, mark string) error
	Win() (bool, string, error)
}

type Game struct {
	players [2]*player.Player
	board   GameAPI
}

func Init(player1, player2 *player.Player, board GameAPI) *Game {
	return &Game{
		players: [2]*player.Player{player1, player2},
		board:   board,
	}
}

func (g *Game) GetTheTurnPlayer() player.Player {
	return *g.players[indexTurnPlayer]
}

func (g *Game) Play(x, y int) (*player.Player, error) {
	err := g.markField(x, y)
	if err != nil {
		return nil, err
	}
	winner, err := g.winner()
	if err != nil && err != NoWinnerError {
		return nil, err
	}
	if winner != nil {
		return winner, nil
	}
	g.switchTurnPlayer()
	return nil, nil
}

func (g *Game) switchTurnPlayer() {
	g.players[indexTurnPlayer], g.players[indexWaitingPlayer] = g.players[indexWaitingPlayer], g.players[indexTurnPlayer]
}

func (g *Game) markField(x, y int) error {
	err := g.board.Mark(x, y, g.players[0].Mark)
	return err
}

func (g *Game) winner() (*player.Player, error) {
	win, mark, err := g.board.Win()
	if err != nil {
		return nil, err
	}
	if win {
		for _, p := range g.players {
			if p.Mark == mark {
				return p, nil
			}
		}
	}
	return nil, NoWinnerError
}
