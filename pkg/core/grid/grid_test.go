package grid_test

import (
	"testing"

	"github.com/edipermadi/sadako/pkg/core/cell"
	"github.com/edipermadi/sadako/pkg/core/column"
	"github.com/edipermadi/sadako/pkg/core/grid"
	"github.com/edipermadi/sadako/pkg/core/row"
	"github.com/stretchr/testify/require"
)

func TestGrid_RowValues(t *testing.T) {
	g := grid.NewGrid(
		[]cell.Value{
			1, 2, 3,
			4, 5, 6,
			7, 8, 9,
		},
	)
	require.Equal(t, []cell.Value{1, 2, 3}, g.RowValues(row.NumberOne))
	require.Equal(t, []cell.Value{4, 5, 6}, g.RowValues(row.NumberTwo))
	require.Equal(t, []cell.Value{7, 8, 9}, g.RowValues(row.NumberThree))
}

func TestGrid_ColumnValues(t *testing.T) {
	g := grid.NewGrid([]cell.Value{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})
	require.Equal(t, []cell.Value{1, 4, 7}, g.ColumnValues(column.NumberOne))
	require.Equal(t, []cell.Value{2, 5, 8}, g.ColumnValues(column.NumberTwo))
	require.Equal(t, []cell.Value{3, 6, 9}, g.ColumnValues(column.NumberThree))
}

func TestGrid_CellValue(t *testing.T) {
	g := grid.NewGrid([]cell.Value{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})

	value := cell.ValueOne
	for number := cell.NumberOne; number <= cell.NumberNine; number++ {
		require.Equal(t, value, g.CellValue(number))
		value += 1
	}
}

func TestGrid_IsSolved(t *testing.T) {
	type testCase struct {
		Title         string
		GivenGrid     grid.Grid
		ExpectedValue bool
	}

	testCases := []testCase{
		{
			Title: "Solved",
			GivenGrid: grid.NewGrid([]cell.Value{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			}),
			ExpectedValue: true,
		},
		{
			Title:         "Empty",
			GivenGrid:     grid.Grid(0),
			ExpectedValue: false,
		},
		{
			Title: "Partial",
			GivenGrid: grid.NewGrid([]cell.Value{
				1, 2, 3,
				4, 5, 6,
				7, 8, 0,
			}),
			ExpectedValue: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			require.Equal(t, tc.ExpectedValue, tc.GivenGrid.IsSolved())
		})
	}
}

func TestGrid_ContainsHasCellWithValue(t *testing.T) {
	g := grid.NewGrid([]cell.Value{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})

	for value := cell.ValueOne; value <= cell.ValueNine; value++ {
		require.True(t, g.HasCellWithValue(value))
	}

}

func TestGrid_Equals(t *testing.T) {
	g1 := grid.NewGrid([]cell.Value{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})

	g2 := grid.NewGrid([]cell.Value{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})

	require.True(t, g1.Equals(g2))
}

func TestGrid_IsEmpty(t *testing.T) {
	g := grid.NewGrid([]cell.Value{})
	require.True(t, g.IsEmpty())
}

func TestGrid_SetCellValue(t *testing.T) {
	g := grid.Grid(0)

	value := cell.ValueOne
	for number := cell.NumberOne; number <= cell.NumberNine; number++ {
		g = g.SetCellValue(number, value)
		require.Equal(t, value, g.CellValue(number))
		value++
	}
}

func TestGrid_IsValid(t *testing.T) {
	type testCase struct {
		Title         string
		GivenGrid     grid.Grid
		ExpectedValue bool
	}

	testCases := []testCase{
		{
			Title: "Solved",
			GivenGrid: grid.NewGrid([]cell.Value{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			}),
			ExpectedValue: true,
		},
		{
			Title:         "Empty",
			GivenGrid:     grid.Grid(0),
			ExpectedValue: true,
		},
		{
			Title: "Partial",
			GivenGrid: grid.NewGrid([]cell.Value{
				1, 2, 3,
				4, 5, 6,
				7, 8, 0,
			}),
			ExpectedValue: true,
		},
		{
			Title: "Duplicate",
			GivenGrid: grid.NewGrid([]cell.Value{
				1, 2, 3,
				4, 5, 6,
				7, 7, 0,
			}),
			ExpectedValue: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			require.Equal(t, tc.ExpectedValue, tc.GivenGrid.IsValid())
		})
	}
}

func TestNewMask(t *testing.T) {
	mask := grid.NewMask([]cell.Value{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})
	require.Equal(t, grid.CompletedGridMask, mask)
}
