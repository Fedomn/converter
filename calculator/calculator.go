package calculator

type (
	Calculator interface {
		CalcDecimal(symbols interface{}) (interface{}, error)
	}
)
