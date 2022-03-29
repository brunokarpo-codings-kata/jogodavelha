package board

import (
	"testing"
)

func Test_ShouldInitABoardWithAllFieldsEmpty(t *testing.T) {
	b := Board{}
	b.Init()

	for x, column := range b.fields {
		for y, cel := range column {
			if cel != "-" {
				t.Errorf(`field [%v][%v] should be initialized with empty value. Expected: %v, got: %v`, x, y, "-", cel)
			}
		}
	}
}
