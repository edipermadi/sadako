package grid_test

import (
	"testing"

	"github.com/edipermadi/sadako/pkg/core/grid"
	"github.com/stretchr/testify/require"
)

func TestPermutation(t *testing.T) {
	grids := grid.Permutation()
	require.Equal(t, 362880, len(grids))
}

func BenchmarkPermutation(b *testing.B) {
	for b.Loop() {
		grid.Permutation()
	}
}
