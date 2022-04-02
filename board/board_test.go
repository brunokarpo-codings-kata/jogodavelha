package board

import "testing"

func TestBoard_Mark(t *testing.T) {
	t.Run("should mark all valid fields", func(t *testing.T) {
		b := Init()
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				mark := "X"
				_ = b.Mark(x, y, mark)

				if b.fields[x][y] != mark {
					t.Logf("error: field [%v][%v] should have value '%v', but got: '%v'", x, y, mark, b.fields[x][y])
					t.Fail()
				}
			}
		}
	})
	t.Run("Should not mark negative indexes field", func(t *testing.T) {
		b := Init()
		for x := -1; x > -10; x-- {
			for y := -1; y > -10; y-- {
				mark := "X"

				err := b.Mark(x, y, mark)
				if err == nil {
					t.Logf("expected error while trying to mark field [%v][%v], but not occours", x, y)
					t.Fail()
				}
			}
		}
	})
	t.Run("Should not mark invalid positive fields", func(t *testing.T) {
		b := Init()
		for x := 4; x < 10; x++ {
			for y := 4; y < 10; y++ {
				mark := "X"

				err := b.Mark(x, y, mark)
				if err == nil {
					t.Logf("expected error while trying to mark field [%v][%v], but not occours", x, y)
					t.Fail()
				}
			}
		}
	})
	t.Run("Should not mark an already marked field", func(t *testing.T) {
		b := Init()

		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				mark := "X"
				_ = b.Mark(x, y, mark)

				mark2 := "O"
				err := b.Mark(x, y, mark2)
				if err == nil {
					t.Logf("expected error. the field [%v][%v] should be marked with '%v', but got '%v'", x, y, mark, b.fields[x][y])
					t.Fail()
				}
			}
		}
	})
}

func TestBoard_Win(t *testing.T) {
	t.Run("empty board should not have winner", func(t *testing.T) {
		b := Init()
		win, mark, _ := b.Win()
		if win {
			t.Log("error: the board should not have winner")
			t.Fail()
		}
		if mark == "-" {
			t.Logf("The mark '%v' is not a valid mark to win the game", mark)
			t.Fail()
		}
	})
	t.Run("should win in horizontal", func(t *testing.T) {
		mark := "X"

		for x := 0; x < 3; x++ {
			b := Init()

			for y := 0; y < 3; y++ {
				_ = b.Mark(x, y, mark)
			}

			win, m, _ := b.Win()

			if !win {
				t.Log("error: the board should have a winner")
				t.Fail()
			}
			if m != mark {
				t.Logf("The mark '%v' should won the game, but got the mark '%v'", mark, m)
			}
		}
	})
	t.Run("should win in vertical", func(t *testing.T) {
		mark := "O"

		for y := 0; y < 3; y++ {
			b := Init()

			for x := 0; x < 3; x++ {
				_ = b.Mark(x, y, mark)
			}

			win, m, _ := b.Win()

			if !win {
				t.Log("error: the board should have a winner")
				t.Fail()
			}
			if m != mark {
				t.Logf("The mark '%v' should won the game, but got the mark '%v'", mark, m)
			}
		}
	})
	t.Run("Should win in diagonal right", func(t *testing.T) {
		mark := "B"

		b := Init()

		for i := 0; i < 3; i++ {
			_ = b.Mark(i, i, mark)
		}

		win, m, _ := b.Win()

		if !win {
			t.Log("error: the board should have a winner")
			t.Fail()
		}
		if m != mark {
			t.Logf("The mark '%v' should won the game, but got the mark '%v'", mark, m)
		}
	})
	t.Run("should win in diagonal left", func(t *testing.T) {
		mark := "S"

		b := Init()

		for x, y := 0, 2; x < 3; x, y = x+1, y-1 {
			_ = b.Mark(x, y, mark)
		}

		win, m, _ := b.Win()

		if !win {
			t.Log("error: the board should have a winner")
			t.Fail()
		}
		if m != mark {
			t.Logf("The mark '%v' should won the game, but got the mark '%v'", mark, m)
		}
	})
	t.Run("Should fail if all field is filled and no winner", func(t *testing.T) {
		b := Init()
		b.fields = fields{
			{"X", "O", "O"},
			{"O", "O", "X"},
			{"X", "X", "O"},
		}

		win, mark, err := b.Win()
		if win {
			t.Log("error: the board should not have winner")
			t.Fail()
		}
		if mark != "" {
			t.Logf("The mark '%v' is not a valid mark to win the game", mark)
			t.Fail()
		}
		if err == nil {
			t.Logf("error expected when field is fully filled and no winner mark")
			t.Fail()
		}
	})
}
