package calculator

type (
	Calculator interface {
		CalcDecimal(symbols interface{}) (int, error)
	}
)
