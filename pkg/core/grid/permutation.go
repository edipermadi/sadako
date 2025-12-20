package grid

import "github.com/edipermadi/sadako/pkg/core/cell"

type CallbackFn func(g Grid) bool
type permutationCtx struct {
	grids    []Grid
	callback CallbackFn
}

func Permutation(callback CallbackFn) []Grid {
	ctx := permutationCtx{grids: nil, callback: callback}
	ctx.permutation(EmptyGrid, cell.NumberOne, EmptyMask)
	return ctx.grids
}

func (c *permutationCtx) permutation(grid Grid, number cell.Number, mask Mask) {
	if number > cell.NumberNine {
		if c.callback != nil {
			if c.callback(grid) {
				c.grids = append(c.grids, grid)
			}
		} else {
			c.grids = append(c.grids, grid)
		}

		return
	}

	for value := cell.ValueOne; value <= cell.ValueNine; value++ {
		if !mask.IsSet(value) {
			c.permutation(grid.SetCellValue(number, value), number+1, mask.Set(value))
		}
	}
}
