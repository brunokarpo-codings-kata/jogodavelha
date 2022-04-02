package game

import (
	"github.com/brunokarpo-codings-kata/jogodavelha/player"
)

const indexTurnPlayer = 0
const indexWaitingPlayer = 1

type GameAPI interface {
	Mark(x, y int, mark string) error
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

func (g *Game) switchTurnPlayer() {
	g.players[indexTurnPlayer], g.players[indexWaitingPlayer] = g.players[indexWaitingPlayer], g.players[indexTurnPlayer]
}

func (g *Game) markField(x, y int) error {
	err := g.board.Mark(x, y, g.players[0].Mark)
	return err
}
