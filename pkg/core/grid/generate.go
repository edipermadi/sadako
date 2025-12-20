package grid

import (
	"github.com/edipermadi/sadako/pkg/core/column"
	"github.com/edipermadi/sadako/pkg/core/row"
)

func GenerateGrids(horizontalNeighbors []Grid, verticalNeighbors []Grid, dictionaries []Grid) []Grid {
	var result []Grid

	if dictionaries == nil {
		dictionaries = Permutation(nil)
	}

	for _, grid := range dictionaries {
		valid := true
		for number := row.NumberOne; number <= row.NumberThree && valid; number++ {
			gridMask := NewMask(grid.RowValues(number))
			for _, n := range horizontalNeighbors {
				if valid && NewMask(n.RowValues(number))&gridMask > 0 {
					valid = false
				}
			}
		}

		for number := column.NumberOne; number <= column.NumberThree && valid; number++ {
			gridMask := NewMask(grid.ColumnValues(number))
			for _, n := range verticalNeighbors {
				if valid && NewMask(n.ColumnValues(number))&gridMask > 0 {
					valid = false
				}
			}
		}

		if valid {
			result = append(result, grid)
		}
	}
	return result
}
