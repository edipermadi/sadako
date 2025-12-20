package grid_test

import (
	"testing"

	"github.com/edipermadi/sadako/pkg/core/grid"
	"github.com/stretchr/testify/require"
)

func TestGenerateGrids(t *testing.T) {
	permutations := grid.Permutation(nil)
	g1 := permutations[0]
	possibilities := grid.GenerateGrids([]grid.Grid{g1}, nil, permutations)
	require.NotEmpty(t, possibilities)
	require.Equal(t, 12096, len(possibilities))

	g2 := possibilities[0]
	possibilities1 := grid.GenerateGrids([]grid.Grid{g1, g2}, nil, permutations)
	require.NotEmpty(t, possibilities1)
	require.Equal(t, 216, len(possibilities1))
}
