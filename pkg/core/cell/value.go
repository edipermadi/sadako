package cell

type Value int

const (
	Empty Value = iota
	One   Value = iota
	Two   Value = iota
	Three Value = iota
	Four  Value = iota
	Five  Value = iota
	Six   Value = iota
	Seven Value = iota
	Eight Value = iota
	Nine  Value = iota
)

func (v Value) IsZero() bool {
	return v == Empty
}
