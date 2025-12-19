package board

import (
	"github.com/edipermadi/sadako/pkg/core/column"
	"github.com/edipermadi/sadako/pkg/core/grid"
	"github.com/edipermadi/sadako/pkg/core/row"
)

type Board [9]grid.Grid

func (b *Board) SetGrid(number grid.Number, value grid.Grid) {
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

	for number := row.NumberOne; number <= row.NumberThree; number++ {
		if !isSolvedRowGrids(b.RowGrids(number)) {
			return false
		}
	}

	for number := column.NumberOne; number <= column.NumberThree; number++ {
		if !isSolvedColumnGrids(b.ColumnGrids(number)) {
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

func isSolvedRowGrids(grids []grid.Grid) bool {
	if len(grids) < 3 {
		return false
	}

	for number := row.NumberOne; number <= row.NumberThree; number++ {
		mask := grid.NewMask(grids[0].RowValues(number)) |
			grid.NewMask(grids[1].RowValues(number)) |
			grid.NewMask(grids[2].RowValues(number))
		if mask != grid.CompletedGridMask {
			return false
		}
	}
	return true
}

func isSolvedColumnGrids(grids []grid.Grid) bool {
	if len(grids) < 3 {
		return false
	}

	for number := column.NumberOne; number <= column.NumberThree; number++ {
		mask := grid.NewMask(grids[0].ColumnValues(number)) |
			grid.NewMask(grids[1].ColumnValues(number)) |
			grid.NewMask(grids[2].ColumnValues(number))
		if mask != grid.CompletedGridMask {
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
		mask := grid.NewMask(grids[0].RowValues(number)) &
			grid.NewMask(grids[1].RowValues(number)) &
			grid.NewMask(grids[2].RowValues(number))
		if mask > 0 {
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
		mask := grid.NewMask(grids[0].ColumnValues(number)) &
			grid.NewMask(grids[1].ColumnValues(number)) &
			grid.NewMask(grids[2].ColumnValues(number))
		if mask > 0 {
			return false
		}

	}
	return true
}
