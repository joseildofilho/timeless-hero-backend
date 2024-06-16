package turnbased_test

import (
	"testing"

	"github.com/magiconair/properties/assert"

	turnbased "github.com/joseildofilho/timeless-hero-backend/src/domain/turn-based"
)

func TestBoard(t *testing.T) {
	t.Run("should create the board", func(t *testing.T) {
		width := uint(10)
		height := uint(11)
		board := turnbased.NewBoard2D(width, height)

		assert.Equal(t, board.Width, width, "Width should be equals")
		assert.Equal(t, board.Height, height, "Heigth should be equals")
		assert.Equal(t, len(board.Grid), int(height), "Grid height should be equals to height")
		assert.Equal(t, len(board.Grid[0]), int(width), "Grid width should be equals to width")
	})

	t.Run("should add a thing to a place", func(t *testing.T) {
		board := turnbased.NewBoard2D(5, 5)

		board.AddToPosition("Player", 1, 2)

		board.GetPlace(1, 1)
	})
}
