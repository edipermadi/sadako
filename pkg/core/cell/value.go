package cell

type Value uint

const (
	ValueEmpty Value = iota
	ValueOne   Value = iota
	ValueTwo   Value = iota
	ValueThree Value = iota
	ValueFour  Value = iota
	ValueFive  Value = iota
	ValueSix   Value = iota
	ValueSeven Value = iota
	ValueEight Value = iota
	ValueNine  Value = iota
)

func (v Value) IsZero() bool {
	return v == ValueEmpty
}

type Number uint

const (
	NumberOne   Number = iota
	NumberTwo   Number = iota
	NumberThree Number = iota
	NumberFour  Number = iota
	NumberFive  Number = iota
	NumberSix   Number = iota
	NumberSeven Number = iota
	NumberEight Number = iota
	NumberNine  Number = iota
)
