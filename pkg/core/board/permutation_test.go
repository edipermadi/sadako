package board_test

import (
	"testing"

	"github.com/edipermadi/sadako/pkg/core/board"
	"github.com/stretchr/testify/require"
)

func TestPermutation(t *testing.T) {
	counter := 0
	boards := board.Permutation(func(b board.Board) (bool, bool) {
		counter++
		return counter > 2, true
	})
	require.NotEmpty(t, boards)
	for _, b := range boards {
		t.Logf("%v", b)
	}
}
