package board_test

import (
	"testing"

	"github.com/edipermadi/sadako/pkg/core/board"
	"github.com/stretchr/testify/require"
)

func TestPermutation(t *testing.T) {
	boards := board.Permutation()
	require.Equal(t, 362880, len(boards))
}
