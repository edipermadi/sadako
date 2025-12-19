package grid

import (
	"github.com/edipermadi/sadako/pkg/core/cell"
	"github.com/edipermadi/sadako/pkg/core/column"
	"github.com/edipermadi/sadako/pkg/core/row"
)

type Grid uint64

func NewGrid(values []cell.Value) Grid {
	grid := Grid(0)
	for number := cell.NumberOne; number <= cell.NumberNine; number++ {
		if len(values) > int(number) {
			grid = grid.SetCellValue(number, values[number])
		}
	}
	return grid
}

func (g Grid) CellValue(n cell.Number) cell.Value {
	n = min(cell.NumberNine, n)
	v := cell.Value((g >> (n * 4)) & 0x0f)
	return min(v, cell.ValueNine)
}

func (g Grid) SetCellValue(n cell.Number, v cell.Value) Grid {
	n = min(cell.NumberNine, n)
	v = min(cell.ValueNine, v)
	return g | Grid(v)<<Grid(n*4)
}

func (g Grid) HasCellWithValue(v cell.Value) bool {
	v = min(cell.ValueNine, v)
	for n := cell.NumberOne; n <= cell.NumberNine; n++ {
		if g.CellValue(n) == v {
			return true
		}
	}
	return false
}

func (g Grid) RowValues(number row.Number) []cell.Value {
	switch number {
	case row.NumberOne:
		return []cell.Value{
			g.CellValue(cell.NumberOne),
			g.CellValue(cell.NumberTwo),
			g.CellValue(cell.NumberThree),
		}
	case row.NumberTwo:
		return []cell.Value{
			g.CellValue(cell.NumberFour),
			g.CellValue(cell.NumberFive),
			g.CellValue(cell.NumberSix),
		}
	case row.NumberThree:
		return []cell.Value{
			g.CellValue(cell.NumberSeven),
			g.CellValue(cell.NumberEight),
			g.CellValue(cell.NumberNine),
		}
	default:
		return nil
	}
}

func (g Grid) ColumnValues(number column.Number) []cell.Value {
	switch number {
	case column.NumberOne:
		return []cell.Value{
			g.CellValue(cell.NumberOne),
			g.CellValue(cell.NumberFour),
			g.CellValue(cell.NumberSeven),
		}
	case column.NumberTwo:
		return []cell.Value{
			g.CellValue(cell.NumberTwo),
			g.CellValue(cell.NumberFive),
			g.CellValue(cell.NumberEight),
		}
	case column.NumberThree:
		return []cell.Value{
			g.CellValue(cell.NumberThree),
			g.CellValue(cell.NumberSix),
			g.CellValue(cell.NumberNine),
		}
	default:
		return nil
	}
}

func (g Grid) IsSolved() bool {
	var mask Mask
	for number := cell.NumberOne; number <= cell.NumberNine; number++ {
		value := g.CellValue(number)
		if value == cell.ValueEmpty || mask.IsSet(value) {
			return false
		}
		mask = mask.Set(value)
	}
	return true
}

func (g Grid) IsEmpty() bool {
	return g&0xfffffffff == 0
}

func (g Grid) Equals(v Grid) bool {
	return g&0xfffffffff == v&0xfffffffff
}

func (g Grid) IsValid() bool {
	var mask Mask
	for number := cell.NumberOne; number <= cell.NumberNine; number++ {
		value := g.CellValue(number)
		if value > cell.ValueEmpty {
			if mask.IsSet(value) {
				return false
			}
			mask = mask.Set(value)
		}
	}
	return true
}

func NewMask(values []cell.Value) Mask {
	var mask Mask
	for _, v := range values {
		if !v.IsZero() {
			mask = mask.Set(v)
		}
	}
	return mask
}
