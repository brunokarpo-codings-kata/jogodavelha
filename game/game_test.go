package game

import (
	"errors"
	"github.com/brunokarpo-codings-kata/jogodavelha/game/mocks"
	"github.com/brunokarpo-codings-kata/jogodavelha/player"
	"github.com/stretchr/testify/assert"
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

		assert.Equal(t, player1, p)
	})
	t.Run("should switch the turn player", func(t *testing.T) {
		g := Init(&player1, &player2, nil)

		// testing
		g.switchTurnPlayer()

		// validating
		p := g.GetTheTurnPlayer()

		assert.Equal(t, player2, p)
	})
	t.Run("should switch the turn player twice backing to the first player", func(t *testing.T) {
		g := Init(&player1, &player2, nil)

		// testing
		g.switchTurnPlayer()
		g.switchTurnPlayer()

		// validating
		p := g.GetTheTurnPlayer()

		assert.Equal(t, player1, p)
	})
}

func TestGame_MarkField(t *testing.T) {
	testWrapper := new(mocks.MockGameWrapper)
	x, y := 0, 0
	t.Run("turn player should mark a valid field", func(t *testing.T) {
		g := Init(&player1, &player2, testWrapper)
		testWrapper.On("Mark", x, y, player1.Mark).Return(nil).Once()
		err := g.markField(x, y)
		testWrapper.AssertExpectations(t)
		assert.Nil(t, err)
	})
	t.Run("turn player should not mark a invalid field", func(t *testing.T) {
		g := Init(&player1, &player2, testWrapper)
		expectedError := errors.New("field already marked or invalid")
		testWrapper.On("Mark", x, y, player1.Mark).Return(expectedError).Once()
		err := g.markField(x, y)
		testWrapper.AssertExpectations(t)
		assert.Equal(t, expectedError, err)
	})
}
