package roman

import (
	"fmt"
	"testing"

	. "github.com/fedomn/converter/util"
)

var rc = initRomanCalculator()

type RomanTest struct {
	Symbols []Symbol
	Wat     interface{}
}

func TestValidateSymbol(t *testing.T) {
	var tests = []RomanTest{
		{[]Symbol{}, rc.MakeError(EmptyErr)},
		{[]Symbol{"D"}, nil},
		{[]Symbol{"D", "J"}, rc.MakeError(SymbolErr, "J")},
	}

	for _, tt := range tests {
		got := rc.validateSymbols(tt.Symbols)
		msg := fmt.Sprintf("input : %+v", tt.Symbols)
		Equals(t, msg, tt.Wat, got)
	}
}

func TestValidateSuccessiveRepeat(t *testing.T) {
	var tests = []RomanTest{
		{[]Symbol{"D"}, nil},
		{[]Symbol{"D", "D"}, rc.MakeError(RepeatErr, "D")},
		{[]Symbol{"I", "I", "I"}, nil},
		{[]Symbol{"I", "I", "I", "I"}, rc.MakeError(RepeatErr, "I")},
	}

	for _, tt := range tests {
		got := rc.validateSuccessiveRepeat(tt.Symbols)
		msg := fmt.Sprintf("input : %+v", tt.Symbols)
		Equals(t, msg, tt.Wat, got)
	}
}

func TestValidateSubtract(t *testing.T) {
	var tests = []RomanTest{
		{[]Symbol{"V", "M"}, rc.MakeError(SubtractErr, "V", "")},
		{[]Symbol{"I", "V", "X", "M"}, rc.MakeError(SubtractErr, "X", "M")},
		{[]Symbol{"I", "V", "X", "L", "C", "D"}, nil},
	}

	for _, tt := range tests {
		got := rc.validateSubtract(tt.Symbols)
		msg := fmt.Sprintf("input : %+v", tt.Symbols)
		Equals(t, msg, tt.Wat, got)
	}
}

func TestMakeSortedCombines(t *testing.T) {
	watCombines := SortedCombines{
		{[]Symbol{"M"}, 1000},
		{[]Symbol{"C", "M"}, 900},
		{[]Symbol{"D"}, 500},
		{[]Symbol{"C", "D"}, 400},
		{[]Symbol{"C"}, 100},
		{[]Symbol{"X", "C"}, 90},
		{[]Symbol{"L"}, 50},
		{[]Symbol{"X", "L"}, 40},
		{[]Symbol{"X"}, 10},
		{[]Symbol{"I", "X"}, 9},
		{[]Symbol{"V"}, 5},
		{[]Symbol{"I", "V"}, 4},
		{[]Symbol{"I"}, 1},
	}
	gotCombines := rc.makeSortedCombines()
	Equals(t, fmt.Sprintf("compare combines"), watCombines, gotCombines)
}
