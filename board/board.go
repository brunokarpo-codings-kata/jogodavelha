package board

import (
	"errors"
)

type Board struct {
	fields [3][3]string
}

func (b *Board) Init() {
	b.fields = [3][3]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}
}

func (b *Board) Mark(x, y int, mark string) error {
	if b.fields[x][y] != "-" {
		return errors.New("you could not mark a field that already is marked")
	}
	b.fields[x][y] = mark
	return nil
}
