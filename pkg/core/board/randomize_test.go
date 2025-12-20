package board_test

import (
	"testing"

	"github.com/edipermadi/sadako/pkg/core/board"
	"github.com/stretchr/testify/require"
)

func TestRandomize(t *testing.T) {
	b := board.Randomize()
	require.NotEmpty(t, b)
	for _, g := range b {
		t.Logf("%v", g.Values())
	}
}
