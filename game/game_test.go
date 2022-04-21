package game

import (
	"errors"
	"github.com/brunokarpo-codings-kata/jogodavelha/player"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockGameWrapper struct {
	mock.Mock
}

func (w *MockGameWrapper) Mark(x, y int, mark string) error {
	args := w.Called(x, y, mark)
	return args.Error(0)
}

func (w *MockGameWrapper) Win() (bool, string, error) {
	args := w.Called()
	return args.Bool(0), args.String(1), args.Error(2)
}

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

func TestGame_markField(t *testing.T) {
	testWrapper := new(MockGameWrapper)
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

func TestGame_winner(t *testing.T) {
	t.Run("should return player 1 win the game", func(t *testing.T) {
		testWrapper := new(MockGameWrapper)
		g := Init(&player1, &player2, testWrapper)
		testWrapper.On("Win").Return(true, player1.Mark, nil).Once()
		winner, err := g.winner()
		assert.Equal(t, &player1, winner)
		assert.Nil(t, err)
	})
	t.Run("should return player 2 win the game", func(t *testing.T) {
		testWrapper := new(MockGameWrapper)
		g := Init(&player1, &player2, testWrapper)
		testWrapper.On("Win").Return(true, player2.Mark, nil).Once()
		winner, err := g.winner()
		assert.Equal(t, &player2, winner)
		assert.Nil(t, err)
	})
	t.Run("should return error when game has no winner yet", func(t *testing.T) {
		testWrapper := new(MockGameWrapper)
		g := Init(&player1, &player2, testWrapper)
		expectedError := errors.New("no winner yet")
		testWrapper.On("Win").Return(false, "", nil).Once()
		winner, err := g.winner()
		assert.Nil(t, winner)
		assert.Equal(t, expectedError, err)
	})
	t.Run("should return board error when win method emit error", func(t *testing.T) {
		testWrapper := new(MockGameWrapper)
		g := Init(&player1, &player2, testWrapper)
		expectedError := errors.New("board error")
		testWrapper.On("Win").Return(false, "", expectedError).Once()
		winner, err := g.winner()
		assert.Nil(t, winner)
		assert.Equal(t, expectedError, err)
	})
}

func TestGame_Play(t *testing.T) {
	x, y := 0, 0
	matchMark := mock.MatchedBy(func(m string) bool { return true })
	t.Run("should play marking a field and no winner in the end of turn", func(t *testing.T) {
		testWrapper := new(MockGameWrapper)
		g := Init(&player1, &player2, testWrapper)
		testWrapper.On("Mark", x, y, matchMark).Return(nil).Once()
		testWrapper.On("Win").Return(false, "", nil).Once()
		winner, err := g.Play(x, y)
		testWrapper.AssertExpectations(t)
		assert.Nil(t, winner)
		assert.Nil(t, err)
		turnPlayer := g.GetTheTurnPlayer()
		assert.Equal(t, player2, turnPlayer)
	})
	t.Run("should play marking a field and return the turn player as winner", func(t *testing.T) {
		testWrapper := new(MockGameWrapper)
		g := Init(&player1, &player2, testWrapper)
		testWrapper.On("Mark", x, y, matchMark).Return(nil).Once()
		testWrapper.On("Win").Return(true, player1.Mark, nil).Once()
		turnPlayer := g.GetTheTurnPlayer()
		winner, err := g.Play(x, y)
		testWrapper.AssertExpectations(t)
		assert.Equal(t, turnPlayer.Name, winner.Name)
		assert.Equal(t, turnPlayer.Mark, winner.Mark)
		assert.Nil(t, err)
	})
	t.Run("should keep the turn if the field already been marked", func(t *testing.T) {
		testWrapper := new(MockGameWrapper)
		g := Init(&player1, &player2, testWrapper)
		markError := errors.New("field already marked")
		testWrapper.On("Mark", x, y, matchMark).Return(markError)
		actualPlayer := g.GetTheTurnPlayer()
		winner, err := g.Play(x, y)
		testWrapper.AssertExpectations(t)
		assert.Nil(t, winner)
		assert.Equal(t, markError, err)
		assert.Equal(t, actualPlayer, g.GetTheTurnPlayer())
	})
	t.Run("should finish the game if 'deu velha'", func(t *testing.T) {
		testWrapper := new(MockGameWrapper)
		g := Init(&player1, &player2, testWrapper)
		winError := errors.New("deu velha")
		testWrapper.On("Mark", x, y, matchMark).Return(nil).Once()
		testWrapper.On("Win").Return(false, "", winError).Once()
		winner, err := g.Play(x, y)
		testWrapper.AssertExpectations(t)
		assert.Nil(t, winner)
		assert.Equal(t, winError, err)
	})
}
