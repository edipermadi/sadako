package grid

import "github.com/edipermadi/sadako/pkg/core/cell"

type permutationCtx struct {
	grids []Grid
}

func Permutation() []Grid {
	var ctx permutationCtx
	ctx.permutation(EmptyGrid, cell.NumberOne, EmptyMask)
	return ctx.grids
}

func (c *permutationCtx) permutation(grid Grid, number cell.Number, mask Mask) {
	if number >= cell.NumberNine {
		c.grids = append(c.grids, grid)
		return
	}

	for value := cell.ValueOne; value <= cell.ValueNine; value++ {
		if !mask.IsSet(value) {
			c.permutation(grid.SetCellValue(number, value), number+1, mask.Set(value))
		}
	}
}
