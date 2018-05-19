package roman

import (
	"container/list"
	"fedomn/converter/util"
	"fmt"
	"reflect"
	"sort"
)

// Roman计算系统
// Roman数组 <=> 十进制数

type (
	Symbol       string
	DecimalValue int
	Romans       map[Symbol]DecimalValue

	SuccessiveRepeatTimes int
	CanSuccessiveRepeated map[Symbol]SuccessiveRepeatTimes
	CanSubtracted         map[Symbol][]Symbol
	Limiter               struct {
		CanSuccessiveRepeated CanSuccessiveRepeated
		CanSubtracted         CanSubtracted
	}

	Combine struct {
		Symbols []Symbol
		Value   DecimalValue
	}
	SortedCombines []Combine

	Calculator struct {
		Romans         Romans
		Limiter        Limiter
		SortedCombines SortedCombines
	}
)

var DefaultCalculator Calculator

func init() {
	DefaultCalculator = initRomanCalculator()
	DefaultCalculator.SortedCombines = DefaultCalculator.makeSortedCombines()
}

func initRomanCalculator() Calculator {
	return Calculator{
		Romans: Romans{
			"I": 1,
			"V": 5,
			"X": 10,
			"L": 50,
			"C": 100,
			"D": 500,
			"M": 1000,
		},
		Limiter: Limiter{
			CanSuccessiveRepeated{
				"I": 3,
				"V": 1,
				"X": 3,
				"L": 1,
				"C": 3,
				"D": 1,
				"M": 3,
			},
			CanSubtracted{
				"I": {"V", "X"},
				"X": {"L", "C"},
				"C": {"D", "M"},
			},
		},
	}
}

// Roman to Decimal

type ErrType int

const (
	EmptyErr ErrType = iota
	SymbolErr
	RepeatErr
	SubtractErr
)

func (c Calculator) MakeError(t ErrType, symbols ...Symbol) error {
	num := len(symbols)
	switch t {
	case SymbolErr, RepeatErr:
		if num < 1 {
			return fmt.Errorf("must give one symbol")
		}
	case SubtractErr:
		if num < 2 {
			return fmt.Errorf("must give two symbol")
		}
	}

	switch t {
	case EmptyErr:
		return fmt.Errorf("empty symbols")
	case SymbolErr:
		return fmt.Errorf("unknown symbol :%+v", symbols[0])
	case RepeatErr:
		times, ok := c.Limiter.CanSuccessiveRepeated[symbols[0]]
		if ok {
			return fmt.Errorf("%v be repeated times > %d in succession", symbols[0], times)
		} else {
			return fmt.Errorf("%v can never be repeated", symbols[0])
		}
	case SubtractErr:
		_, ok := c.Limiter.CanSubtracted[symbols[0]]
		if ok {
			return fmt.Errorf("%v can never be substracted by %v", symbols[0], symbols[1])
		} else {
			return fmt.Errorf("%v can never be substracted", symbols[0])
		}
	}
	return fmt.Errorf("unknown input parameter: %+v %+v", t, symbols)
}

func (c Calculator) validateSymbols(symbols []Symbol) error {
	if len(symbols) == 0 {
		return c.MakeError(EmptyErr)
	}
	for _, symbol := range symbols {
		if _, ok := c.Romans[symbol]; !ok {
			return c.MakeError(SymbolErr, symbol)
		}
	}
	return nil
}

func (c Calculator) validateSuccessiveRepeat(symbols []Symbol) error {
	currentSymbol := Symbol("")
	currentTimes := SuccessiveRepeatTimes(0)
	for _, symbol := range symbols {
		if symbol != currentSymbol {
			currentSymbol = symbol
			currentTimes = 1
		} else {
			currentTimes++
		}
		if currentTimes > c.Limiter.CanSuccessiveRepeated[currentSymbol] {
			return c.MakeError(RepeatErr, currentSymbol)
		}
	}
	return nil
}

func (c Calculator) validateSubtract(symbols []Symbol) error {
	startIdx := 0
	endIdx := 1
	for endIdx < len(symbols) {
		startSymbol := symbols[startIdx]
		endSymbol := symbols[endIdx]

		// Generally, symbols are placed in order of value, starting with the largest values
		if c.Romans[startSymbol] >= c.Romans[endSymbol] {
			startIdx += 1
			endIdx += 1
			continue
		}

		// When smaller values precede larger values, the smaller values are subtracted from the larger values
		if util.Contains(c.Limiter.CanSubtracted[startSymbol], endSymbol) {
			startIdx += 2
			endIdx += 2
			continue
		}
		return c.MakeError(SubtractErr, startSymbol, endSymbol)
	}
	return nil
}

func (c Calculator) Validate(symbols []Symbol) error {
	if err := c.validateSymbols(symbols); err != nil {
		return err
	}
	if err := c.validateSuccessiveRepeat(symbols); err != nil {
		return err
	}
	if err := c.validateSubtract(symbols); err != nil {
		return err
	}
	return nil
}

func (c Calculator) R2D(symbols []Symbol) (DecimalValue, error) {
	if err := c.Validate(symbols); err != nil {
		return 0, err
	}

	l := list.New()
	for _, each := range symbols {
		l.PushBack(each)
	}

	resVal := DecimalValue(0)

	for l.Len() > 0 {
		startNode := l.Front()
		startSymbol := startNode.Value.(Symbol)
		startVal := c.Romans[startSymbol]

		endNode := startNode.Next()
		if endNode == nil {
			resVal += startVal
			return resVal, nil
		}
		endSymbol := endNode.Value.(Symbol)
		endVal := c.Romans[endSymbol]

		if startVal >= endVal {
			resVal += startVal
			l.Remove(startNode)
		} else {
			resVal += endVal - startVal
			l.Remove(startNode)
			l.Remove(endNode)
		}
	}

	return resVal, nil
}

func (c Calculator) CalcDecimal(symbols interface{}) (interface{}, error) {
	symbolsVal := reflect.ValueOf(symbols)
	if symbolsVal.Kind() != reflect.Slice {
		return 0, fmt.Errorf("symbols type must be slice")
	}

	romanSymbols := make([]Symbol, symbolsVal.Len())
	for i := 0; i < symbolsVal.Len(); i++ {
		symbol := symbolsVal.Index(i)
		symbolStr := symbol.String()
		if symbol.Kind() != reflect.String {
			return 0, fmt.Errorf("symbol: %v type must be string", symbolStr)
		}
		romanSymbols[i] = Symbol(symbolStr)
	}

	return c.R2D(romanSymbols)
}

// Decimal to Roman

func (s SortedCombines) Len() int {
	return len(s)
}

func (s SortedCombines) Less(i, j int) bool {
	return s[i].Value >= s[j].Value
}

func (s SortedCombines) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (c Calculator) makeSortedCombines() SortedCombines {
	sortedCombines := make(SortedCombines, 0)
	for symbol, val := range c.Romans {
		sortedCombines = append(sortedCombines, Combine{[]Symbol{symbol}, val})
		for _, subSymbol := range c.Limiter.CanSubtracted[symbol] {
			combineAry := []Symbol{symbol, subSymbol}
			combineVal, _ := c.R2D(combineAry)
			sortedCombines = append(sortedCombines, Combine{combineAry, combineVal})
		}
	}
	sort.Sort(sortedCombines)
	return sortedCombines
}

func (c Calculator) D2R(val DecimalValue) ([]Symbol, error) {
	resSymbols := make([]Symbol, 0)
	idx := 0
	for val > 0 {
		combine := c.SortedCombines[idx]
		for val >= combine.Value {
			err := c.validateSuccessiveRepeat(append(resSymbols, combine.Symbols...))
			if err != nil {
				break
			}
			resSymbols = append(resSymbols, combine.Symbols...)
			val -= combine.Value
		}
		idx++
	}
	return resSymbols, nil
}

func (c Calculator) ToString(symbols []Symbol) string {
	res := ""
	for _, symbol := range symbols {
		res += string(symbol)
	}
	return res
}
