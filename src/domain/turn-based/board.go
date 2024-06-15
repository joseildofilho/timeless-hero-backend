package turnbased

type Board2D struct {
	Grid [][]Place
	Width, Height uint
}

type Place struct {
	things []any
}

func NewPlace() *Place {
	return &Place {
		things: []any{},
	}
}

func NewBoard2D(width, heigth uint) *Board2D {
	return &Board2D{
		Grid: newGrid(width, heigth),
		Width: width,
		Height: heigth,
	}
}

func newGrid(width, heigth uint) [][]Place {
	board := make([][]Place, heigth)

	for i := range 10 {
		board[i] = make([]Place, width)
	}

	return board
}
