package permutation

import "github.com/edipermadi/sadako/pkg/core/grid"

type Context struct {
	Grids []grid.Grid
}

func (c *Context) Permutation(g grid.Grid) {
	if g.Depth() >= 9 {
		c.Grids = append(c.Grids, g)
		return
	}

	for i, v := range g {
		if v == 0 {
			for j := 1; j <= 9; j++ {
				if !g.Contains(j) {
					g1 := g
					g1[i] = j
					c.Permutation(g1)
				}
			}
			return
		}
	}
}
