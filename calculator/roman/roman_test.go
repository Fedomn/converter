package roman_test

import (
	"fmt"
	"testing"

	. "github.com/fedomn/converter/calculator/roman"
	. "github.com/fedomn/converter/util"
)

var rc = DefaultCalculator

func TestMakeError(t *testing.T) {
	var tests = []struct {
		err     ErrType
		symbols []Symbol
		wat     interface{}
	}{
		{SymbolErr, []Symbol{}, fmt.Errorf("must give one symbol")},
		{RepeatErr, []Symbol{}, fmt.Errorf("must give one symbol")},
		{SubtractErr, []Symbol{}, fmt.Errorf("must give two symbol")},
	}
	for _, tt := range tests {
		got := rc.MakeError(tt.err, tt.symbols...)
		Equals(t, fmt.Sprintf("input : %+v", tt.err), tt.wat, got)
	}
}

// 验证处理顺序
// validateSymbols -> validateSuccessiveRepeat -> validateSubtract
func TestValidate(t *testing.T) {
	var tests = []RomanTest{
		{[]Symbol{}, rc.MakeError(EmptyErr)},
		{[]Symbol{"I", "I", "I", "J"}, rc.MakeError(SymbolErr, "J")},
		{[]Symbol{"I", "I", "I", "I", "M"}, rc.MakeError(RepeatErr, "I")},
	}

	for _, tt := range tests {
		err := rc.Validate(tt.Symbols)
		msg := fmt.Sprintf("input : %+v", tt.Symbols)
		Equals(t, msg, tt.Wat, err)
	}
}

func TestR2D(t *testing.T) {
	var tests = []RomanTest{
		{[]Symbol{"I"}, DecimalValue(1)},
		{[]Symbol{"X", "L", "I", "I"}, DecimalValue(42)},
		{[]Symbol{"M", "X", "I"}, DecimalValue(1011)},
		{[]Symbol{"I", "V", "X", "L", "C", "D"}, DecimalValue(444)},
	}

	for _, tt := range tests {
		res, err := rc.R2D(tt.Symbols)
		msg := fmt.Sprintf("input roman symbol: %v", tt.Symbols)
		if err != nil {
			Equals(t, msg, tt.Wat, err)
		} else {
			Equals(t, msg, tt.Wat, res)
		}
	}
}

func TestD2R(t *testing.T) {
	var tests = []struct {
		val DecimalValue
		wat string
	}{
		{0, ""},
		{1, "I"},
		{3333, "MMMCCCXXXIII"},
		{4001, "MMMCMCI"},
	}
	for _, tt := range tests {
		symbols, err := rc.D2R(tt.val)
		Ok(t, fmt.Sprintf("decimal to roman err"), err)
		Equals(t, fmt.Sprintf("input decimal: %+v", tt.val), tt.wat, rc.ToString(symbols))
	}
}
