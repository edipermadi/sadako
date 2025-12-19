package grid_test

import (
	"testing"

	"github.com/edipermadi/sadako/pkg/core/cell"
	"github.com/edipermadi/sadako/pkg/core/column"
	"github.com/edipermadi/sadako/pkg/core/grid"
	"github.com/edipermadi/sadako/pkg/core/row"
	"github.com/stretchr/testify/require"
)

func TestGrid_Row(t *testing.T) {
	g := grid.NewGrid(
		[]cell.Value{
			1, 2, 3,
			4, 5, 6,
			7, 8, 9,
		},
	)
	require.Equal(t, []cell.Value{1, 2, 3}, g.Row(row.NumberOne))
	require.Equal(t, []cell.Value{4, 5, 6}, g.Row(row.NumberTwo))
	require.Equal(t, []cell.Value{7, 8, 9}, g.Row(row.NumberThree))
}

func TestGrid_Column(t *testing.T) {
	g := grid.NewGrid([]cell.Value{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})
	require.Equal(t, []cell.Value{1, 4, 7}, g.Column(column.NumberOne))
	require.Equal(t, []cell.Value{2, 5, 8}, g.Column(column.NumberTwo))
	require.Equal(t, []cell.Value{3, 6, 9}, g.Column(column.NumberThree))
}
