package board_test

import (
	"testing"

	"github.com/edipermadi/sadako/pkg/core/board"
	"github.com/edipermadi/sadako/pkg/core/cell"
	"github.com/edipermadi/sadako/pkg/core/grid"
	"github.com/stretchr/testify/require"
)

func TestBoard_IsSolved(t *testing.T) {
	b := board.Board([]grid.Grid{
		grid.NewGrid([]cell.Value{
			4, 3, 5,
			6, 8, 2,
			1, 9, 7,
		}),
		grid.NewGrid([]cell.Value{
			2, 6, 9,
			5, 7, 1,
			8, 3, 4,
		}),
		grid.NewGrid([]cell.Value{
			7, 8, 1,
			4, 9, 3,
			5, 6, 2,
		}),

		grid.NewGrid([]cell.Value{
			8, 2, 6,
			3, 7, 4,
			9, 5, 1,
		}),
		grid.NewGrid([]cell.Value{
			1, 9, 5,
			6, 8, 2,
			7, 4, 3,
		}),
		grid.NewGrid([]cell.Value{
			3, 4, 7,
			9, 1, 5,
			6, 2, 8,
		}),

		grid.NewGrid([]cell.Value{
			5, 1, 9,
			2, 4, 8,
			7, 6, 3,
		}),
		grid.NewGrid([]cell.Value{
			3, 2, 6,
			9, 5, 7,
			4, 1, 8,
		}),
		grid.NewGrid([]cell.Value{
			8, 7, 4,
			1, 3, 6,
			2, 5, 9,
		}),
	})

	require.True(t, b.IsValid())
}

func TestBoard_IsValid(t *testing.T) {
	b := board.Board([]grid.Grid{
		grid.NewGrid([]cell.Value{
			4, 3, 5,
			6, 8, 2,
			1, 9, 7,
		}),
		grid.NewGrid([]cell.Value{
			2, 6, 9,
			5, 7, 1,
			8, 3, 4,
		}),
		grid.NewGrid([]cell.Value{
			7, 8, 1,
			4, 9, 3,
			5, 6, 2,
		}),

		grid.NewGrid([]cell.Value{
			8, 2, 6,
			3, 7, 4,
			9, 5, 1,
		}),
		grid.NewGrid([]cell.Value{
			1, 9, 5,
			6, 8, 2,
			7, 4, 3,
		}),
		grid.NewGrid([]cell.Value{
			3, 4, 7,
			9, 1, 5,
			6, 2, 8,
		}),

		grid.NewGrid([]cell.Value{
			5, 1, 9,
			2, 4, 8,
			7, 6, 3,
		}),
		grid.NewGrid([]cell.Value{
			3, 2, 6,
			9, 5, 7,
			4, 1, 8,
		}),
		grid.NewGrid([]cell.Value{
			8, 7, 4,
			1, 3, 6,
			2, 5, 9,
		}),
	})

	require.True(t, b.IsValid())
}
