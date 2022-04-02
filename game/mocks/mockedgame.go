package mocks

import "github.com/stretchr/testify/mock"

type MockGameWrapper struct {
	mock.Mock
}

func (w *MockGameWrapper) Mark(x, y int, mark string) error {
	args := w.Called(x, y, mark)
	return args.Get(0).(error)
}
