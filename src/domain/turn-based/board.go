package turnbased

type Board2D struct {
	Grid          [][]Place
	Width, Height uint
}

func NewBoard2D(width, heigth uint) *Board2D {
	return &Board2D{
		Grid:   newGrid(width, heigth),
		Width:  width,
		Height: heigth,
	}
}

func (b Board2D) AddToPosition(aThing any, x, y uint) error {
	b.GetPlace(x, y).AddToPlace(aThing)

	return nil
}

func (b Board2D) GetPlace(x, y uint) *Place {
	return &b.Grid[y][x]
}

func newGrid(width, heigth uint) [][]Place {
	board := make([][]Place, heigth)

	for i := range heigth {
		board[i] = make([]Place, width)
	}

	return board
}
