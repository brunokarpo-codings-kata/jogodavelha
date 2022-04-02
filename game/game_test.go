package game

import (
	"github.com/brunokarpo-codings-kata/jogodavelha/player"
	"testing"
)

var player1 = player.Player{
	Name: "player1",
	Mark: "X",
}
var player2 = player.Player{
	Name: "player2",
	Mark: "O",
}

func TestGame_GetTheTurnPlayer(t *testing.T) {
	t.Run("get the first turn player of the game", func(t *testing.T) {
		g := Init(&player1, &player2, nil)

		p := g.GetTheTurnPlayer()

		if p != player1 {
			t.Logf("error: the first turn player should be %v, but got %v", player1, p)
			t.Fail()
		}
	})
	t.Run("should switch the turn player", func(t *testing.T) {
		g := Init(&player1, &player2, nil)

		// testing
		g.switchTurnPlayer()

		// validating
		p := g.GetTheTurnPlayer()

		if p != player2 {
			t.Logf("error: the turn player should be %v, but got %v", player2, p)
			t.Fail()
		}
	})
	t.Run("should switch the turn player twice backing to the first player", func(t *testing.T) {
		g := Init(&player1, &player2, nil)

		// testing
		g.switchTurnPlayer()
		g.switchTurnPlayer()

		// validating
		p := g.GetTheTurnPlayer()

		if p != player1 {
			t.Logf("error: the turn player should be %v, but got %v", player1, p)
			t.Fail()
		}
	})
}

//
//func TestGame_MarkField(t *testing.T) {
//	testWrapper := new(mocks.MockGameWrapper)
//	x, y := 0, 0
//	t.Run("turn player should mark a valid field", func(t *testing.T) {
//		g := Init(&player1, &player2, testWrapper)
//		testWrapper.On("Mark", x, y, player1.Mark).Return(nil).Once()
//		g.markField()
//	})
//}
