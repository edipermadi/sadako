package grid_test

import (
	"testing"

	"github.com/edipermadi/sadako/pkg/core/grid"
	"github.com/stretchr/testify/require"
)

func TestGrid_Depth(t *testing.T) {
	require.Equal(t, 0, grid.Grid{}.Depth())
}

func TestGrid_Row(t *testing.T) {
	g := grid.Grid{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	require.Equal(t, []int{1, 2, 3}, g.Row(0))
	require.Equal(t, []int{4, 5, 6}, g.Row(1))
	require.Equal(t, []int{7, 8, 9}, g.Row(2))
}

func TestGrid_Column(t *testing.T) {
	g := grid.Grid{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	require.Equal(t, []int{1, 4, 7}, g.Column(0))
	require.Equal(t, []int{2, 5, 8}, g.Column(1))
	require.Equal(t, []int{3, 6, 9}, g.Column(2))
}
