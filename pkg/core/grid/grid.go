package grid

import (
	"github.com/edipermadi/sadako/pkg/core/cell"
	"github.com/edipermadi/sadako/pkg/core/column"
	"github.com/edipermadi/sadako/pkg/core/row"
)

const Size = uint(9)

type Grid uint64

func NewGrid(values []cell.Value) Grid {
	grid := Grid(0)
	for number := cell.NumberOne; number <= cell.NumberNine; number++ {
		if len(values) > int(number) {
			grid = grid.Set(number, values[number])
		}
	}
	return grid
}
func (g Grid) CellValue(n cell.Number) cell.Value {
	if n > cell.NumberNine {
		return cell.ValueEmpty
	}

	v := cell.Value((g >> (n * 4)) & 0x0f)
	return min(v, cell.ValueNine)
}

func (g Grid) Set(n cell.Number, v cell.Value) Grid {
	n = min(cell.NumberNine, n)
	v = min(cell.ValueNine, v)
	return g | Grid(v)<<Grid(n*4)
}

func (g Grid) Contains(v cell.Value) bool {
	v = min(cell.ValueNine, v)
	for n := cell.NumberOne; n <= cell.NumberNine; n++ {
		if g.CellValue(n) == v {
			return true
		}
	}
	return false
}

func (g Grid) Row(n row.Number) []cell.Value {
	switch n {
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

func (g Grid) Column(n column.Number) []cell.Value {
	switch n {
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
