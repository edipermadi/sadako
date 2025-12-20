package board

import (
	"github.com/edipermadi/sadako/pkg/core/grid"
)

type CallbackFn func(b Board) (bool, bool)

func Permutation(callback CallbackFn) []Board {
	var boards []Board
	allPossibilities := grid.Permutation(nil)
	g1Possibilities := allPossibilities
	for _, g1 := range g1Possibilities {
		g2Possibilities := grid.GenerateGrids([]grid.Grid{g1}, nil, allPossibilities)
		for _, g2 := range g2Possibilities {
			g3Possibilities := grid.GenerateGrids([]grid.Grid{g1, g2}, nil, allPossibilities)
			for _, g3 := range g3Possibilities {
				g4Possibilities := grid.GenerateGrids(nil, []grid.Grid{g1}, allPossibilities)
				for _, g4 := range g4Possibilities {
					g5Possibilities := grid.GenerateGrids([]grid.Grid{g4}, []grid.Grid{g2}, allPossibilities)
					for _, g5 := range g5Possibilities {
						g6Possibilities := grid.GenerateGrids([]grid.Grid{g4, g5}, []grid.Grid{g3}, allPossibilities)
						for _, g6 := range g6Possibilities {
							g7Possibilities := grid.GenerateGrids(nil, []grid.Grid{g1, g4}, allPossibilities)
							for _, g7 := range g7Possibilities {
								g8Possibilities := grid.GenerateGrids([]grid.Grid{g7}, []grid.Grid{g2, g5}, allPossibilities)
								for _, g8 := range g8Possibilities {
									g9Possibilities := grid.GenerateGrids([]grid.Grid{g7, g8}, []grid.Grid{g3, g6}, allPossibilities)
									for _, g9 := range g9Possibilities {
										board := [9]grid.Grid{g1, g2, g3, g4, g5, g6, g7, g8, g9}
										if callback != nil {
											stop, appendBoard := callback(board)
											if appendBoard {
												boards = append(boards, board)
											}
											if stop {
												return boards
											}
										} else {
											boards = append(boards, board)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return boards
}
