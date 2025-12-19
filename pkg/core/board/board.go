package board

import (
	"github.com/edipermadi/sadako/pkg/core/column"
	"github.com/edipermadi/sadako/pkg/core/grid"
	"github.com/edipermadi/sadako/pkg/core/row"
)

type Board [9]grid.Grid

func (b *Board) SetGrid(number grid.Grid, value grid.Grid) {
	b[number] = value
}

func (b *Board) RowGrids(number row.Number) []grid.Grid {
	switch number {
	case row.NumberOne:
		return []grid.Grid{
			b[grid.NumberOne],
			b[grid.NumberTwo],
			b[grid.NumberThree],
		}
	case row.NumberTwo:
		return []grid.Grid{
			b[grid.NumberFour],
			b[grid.NumberFive],
			b[grid.NumberSix],
		}
	case row.NumberThree:
		return []grid.Grid{
			b[grid.NumberSeven],
			b[grid.NumberEight],
			b[grid.NumberNine],
		}
	default:
		return nil
	}
}

func (b *Board) ColumnGrids(number column.Number) []grid.Grid {
	switch number {
	case column.NumberOne:
		return []grid.Grid{
			b[grid.NumberOne],
			b[grid.NumberFour],
			b[grid.NumberSeven],
		}
	case column.NumberTwo:
		return []grid.Grid{
			b[grid.NumberTwo],
			b[grid.NumberFive],
			b[grid.NumberEight],
		}
	case column.NumberThree:
		return []grid.Grid{
			b[grid.NumberThree],
			b[grid.NumberSix],
			b[grid.NumberNine],
		}
	default:
		return nil
	}
}

func (b *Board) IsSolved() bool {
	for number := grid.NumberOne; number <= grid.NumberNine; number++ {
		if !b[number].IsSolved() {
			return false
		}
	}

	return true
}

func (b *Board) IsValid() bool {
	for number := row.NumberOne; number <= row.NumberThree; number++ {
		if !isValidRowGrids(b.RowGrids(number)) {
			return false
		}
	}

	for number := column.NumberOne; number <= column.NumberThree; number++ {
		if !isValidColumnGrids(b.ColumnGrids(number)) {
			return false
		}
	}

	return true
}

func isValidRowGrids(grids []grid.Grid) bool {
	if len(grids) < 3 {
		return false
	}

	for number := row.NumberOne; number <= row.NumberThree; number++ {
		var mask grid.Mask
		for i := 0; i < 3; i++ {
			mask |= grid.NewMask(grids[i].RowValues(number))
		}
		if mask != grid.CompletedGridMask {
			return false
		}
	}
	return true
}

func isValidColumnGrids(grids []grid.Grid) bool {
	if len(grids) < 3 {
		return false
	}

	for number := column.NumberOne; number <= column.NumberThree; number++ {
		var mask grid.Mask
		for i := 0; i < 3; i++ {
			mask |= grid.NewMask(grids[i].ColumnValues(number))
		}
		if mask != grid.CompletedGridMask {
			return false
		}
	}
	return true
}
