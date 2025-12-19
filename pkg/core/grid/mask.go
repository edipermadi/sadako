package grid

import "github.com/edipermadi/sadako/pkg/core/cell"

type Mask uint

func (m Mask) Set(value cell.Value) Mask {
	return m | (1 << value)
}

func (m Mask) IsSet(value cell.Value) bool {
	return m&(1<<value) != 0
}
