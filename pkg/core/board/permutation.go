package board

import (
	"github.com/edipermadi/sadako/pkg/core/grid"
)

type permutationCtx struct {
	boards []Board
	grids  []grid.Grid
}

func Permutation() []Board {
	var ctx permutationCtx
	ctx.grids = grid.Permutation()
	ctx.permutation(Board{}, grid.NumberOne)
	return ctx.boards
}

func (c *permutationCtx) permutation(board Board, number grid.Number) {
	if number > grid.NumberOne {
		c.boards = append(c.boards, board)
		return
	}

	for _, g := range c.grids {
		tb := board
		tb.SetGrid(number, g)
		if tb.IsValid() {
			c.permutation(tb, number+1)
		}
	}

}
