package board

import (
	"math/rand/v2"

	"github.com/edipermadi/sadako/pkg/core/grid"
)

func Randomize() Board {
	var board Board
	allPossibilities := grid.Permutation(nil)
	g1Possibilities := allPossibilities
	for _, idxG1 := range rand.Perm(len(g1Possibilities)) {
		g1 := g1Possibilities[idxG1]
		g2Possibilities := grid.GenerateGrids([]grid.Grid{g1}, nil, allPossibilities)
		for _, g2Idx := range rand.Perm(len(g2Possibilities)) {
			g2 := g2Possibilities[g2Idx]
			g3Possibilities := grid.GenerateGrids([]grid.Grid{g1, g2}, nil, allPossibilities)
			for _, g3Idx := range rand.Perm(len(g3Possibilities)) {
				g3 := g3Possibilities[g3Idx]
				g4Possibilities := grid.GenerateGrids(nil, []grid.Grid{g1}, allPossibilities)
				for _, g4Idx := range rand.Perm(len(g4Possibilities)) {
					g4 := g4Possibilities[g4Idx]
					g5Possibilities := grid.GenerateGrids([]grid.Grid{g4}, []grid.Grid{g2}, allPossibilities)
					for _, g5Idx := range rand.Perm(len(g5Possibilities)) {
						g5 := g5Possibilities[g5Idx]
						g6Possibilities := grid.GenerateGrids([]grid.Grid{g4, g5}, []grid.Grid{g3}, allPossibilities)
						for _, g6Idx := range rand.Perm(len(g6Possibilities)) {
							g6 := g6Possibilities[g6Idx]
							g7Possibilities := grid.GenerateGrids(nil, []grid.Grid{g1, g4}, allPossibilities)
							for _, g7Idx := range rand.Perm(len(g7Possibilities)) {
								g7 := g7Possibilities[g7Idx]
								g8Possibilities := grid.GenerateGrids([]grid.Grid{g7}, []grid.Grid{g2, g5}, allPossibilities)
								for _, g8Idx := range rand.Perm(len(g8Possibilities)) {
									g8 := g8Possibilities[g8Idx]
									g9Possibilities := grid.GenerateGrids([]grid.Grid{g7, g8}, []grid.Grid{g3, g6}, allPossibilities)
									for _, g9Idx := range rand.Perm(len(g9Possibilities)) {
										g9 := g9Possibilities[g9Idx]
										board = [9]grid.Grid{g1, g2, g3, g4, g5, g6, g7, g8, g9}
										return board
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return board
}
