package board

import (
	"errors"
	"fmt"
)

var emptyField = "-"
var DeuVelhaError = errors.New("deu velha")

type fields [3][3]string

type Board struct {
	fields
}

func Init() *Board {
	return &Board{
		fields{
			{emptyField, emptyField, emptyField},
			{emptyField, emptyField, emptyField},
			{emptyField, emptyField, emptyField},
		},
	}
}

func (b *Board) Mark(x, y int, mark string) (err error) {
	if x < 0 || y < 0 ||
		x > 2 || y > 2 {
		err = errors.New("invalid fields")
		return
	}
	if b.fields[x][y] != emptyField {
		err = errors.New("field already marked")
		return
	}
	b.fields[x][y] = mark
	return
}

func (b *Board) PrintBoard() {
	fmt.Printf(" %v | %v | %v \n", b.fields[0][0], b.fields[0][1], b.fields[0][2])
	fmt.Println("---|---|---")
	fmt.Printf(" %v | %v | %v \n", b.fields[1][0], b.fields[1][1], b.fields[1][2])
	fmt.Println("---|---|---")
	fmt.Printf(" %v | %v | %v \n", b.fields[2][0], b.fields[2][1], b.fields[2][2])
}

type winner struct {
	win  bool
	mark string
	err  error
}

func (b *Board) Win() (win bool, mark string, err error) {

	c := make(chan winner, 4)

	go b.verifyHorizontal(c)
	go b.verifyVertical(c)
	go b.verifyDiagonal(c)
	go b.allFieldMarked(c)

	for i := 0; i < 4; i++ {
		u := <-c
		if u.win || u.err != nil {
			return u.win, u.mark, u.err
		}
	}

	return
}

func (b *Board) verifyHorizontal(c chan winner) {
	for x := 0; x < 3; x++ {
		if b.fields[x][0] != emptyField &&
			b.fields[x][0] == b.fields[x][1] &&
			b.fields[x][1] == b.fields[x][2] {
			c <- winner{true, b.fields[x][0], nil}
			return
		}
	}
	c <- winner{false, "", nil}
}

func (b *Board) verifyVertical(c chan winner) {
	for y := 0; y < 3; y++ {
		if b.fields[0][y] != emptyField &&
			b.fields[0][y] == b.fields[1][y] &&
			b.fields[1][y] == b.fields[2][y] {
			c <- winner{true, b.fields[0][y], nil}
			return
		}
	}
	c <- winner{false, "", nil}
}

func (b *Board) verifyDiagonal(c chan winner) {
	if b.fields[0][0] != emptyField &&
		b.fields[0][0] == b.fields[1][1] &&
		b.fields[1][1] == b.fields[2][2] {
		c <- winner{true, b.fields[0][0], nil}
		return
	}

	if b.fields[0][2] != emptyField &&
		b.fields[0][2] == b.fields[1][1] &&
		b.fields[1][1] == b.fields[2][0] {
		c <- winner{true, b.fields[0][2], nil}
		return
	}

	c <- winner{false, "", nil}
}

func (b *Board) allFieldMarked(c chan winner) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if b.fields[x][y] == emptyField {
				c <- winner{false, "", nil}
				return
			}
		}
	}
	c <- winner{false, "", DeuVelhaError}
}
