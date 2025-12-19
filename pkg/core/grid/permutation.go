package grid

type permutationCtx struct {
	grids []Grid
}

func Permutation() []Grid {
	var ctx permutationCtx
	ctx.permutation(Grid{}, 0, 0)
	return ctx.grids
}

func (c *permutationCtx) permutation(g Grid, depth int, mask uint32) {
	if depth >= 9 {
		c.grids = append(c.grids, g)
		return
	}

	for i := 1; i <= 9; i++ {
		if mask&(1<<uint(i)) == 0 {
			g1 := g
			g1[depth] = i
			c.permutation(g1, depth+1, mask|(1<<i))
		}
	}
}
