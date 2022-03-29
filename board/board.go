package board

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
