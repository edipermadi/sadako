package permutation_test

import (
	"testing"

	"github.com/edipermadi/sadako/pkg/core/grid"
	"github.com/edipermadi/sadako/pkg/core/permutation"
	"github.com/stretchr/testify/require"
)

func TestPermutation(t *testing.T) {
	ctx := permutation.Context{}
	ctx.Permutation(grid.Grid{})
	require.Equal(t, 362880, len(ctx.Grids))
}
