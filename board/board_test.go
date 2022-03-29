package board

import (
	"testing"
)

func Test_ShouldInitABoardWithAllFieldsEmpty(t *testing.T) {
	b := Board{}
	b.Init()

	for x, column := range b.fields {
		for y, cel := range column {
			if cel != emptyMark {
				t.Errorf(`field [%v][%v] should be initialized with empty value. Expected: %v, got: %v`, x, y, "-", cel)
			}
		}
	}
}

func Test_ShouldAllowMarkAValidEmptyFieldOnTheBoard(t *testing.T) {
	b := Board{}
	b.Init()
	mark := "X"

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.Mark(i, j, mark)

			if b.fields[i][j] != mark {
				t.Errorf(`The field [%v][%v] was not marked properly. Expected '%v', but got '%v'`, i, j, mark, b.fields[i][j])
			}
		}
	}
}

func Test_ShouldNotAllowToMarkAValidFieldThatAlreadyBeMarked(t *testing.T) {
	b := Board{}
	b.Init()
	mark := "X"
	mark2 := "O"

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			err := b.Mark(i, j, mark) // valid marking

			err = b.Mark(i, j, mark2) // invalid marking
			if err == nil {
				t.Errorf(`should not allow to mark the field [%v][%v] with mark [%v]. Expected error, and the field should have value '%v', but got %v`, i, j, mark2, mark, b.fields[i][j])
			}
		}
	}
}

func Test_ShouldNotAllowMarkingInvalidNegativePositionFields(t *testing.T) {
	b := Board{}
	b.Init()
	mark := "X"

	err := b.Mark(-1, -1, mark)
	if err == nil {
		t.Errorf(`should not allow to mark invalid field [-1][-1]`)
	}
}

func Test_ShouldNotAllowMarkingInvalidPositivePositionFields(t *testing.T) {
	b := Board{}
	b.Init()
	mark := "X"

	err := b.Mark(3, 3, mark)
	if err == nil {
		t.Errorf(`should not allow to mark invalid field [3][3]`)
	}
}
