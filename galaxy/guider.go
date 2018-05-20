package galaxy

import (
	"fedomn/converter/calculator/roman"
	. "fedomn/converter/galaxy/handlers"
	. "fedomn/converter/galaxy/models"
)

var DefaultGuider Guider

func init() {
	DefaultGuider = initGuider()
	DefaultGuider.Use(AliasHandler{})
	DefaultGuider.Use(GoodsHandler{})
	DefaultGuider.Use(HowMuchHandler{})
	DefaultGuider.Use(HowManyUnitHandler{})
}

func initGuider() Guider {
	return Guider{
		Alias:      make(Alias),
		Goods:      make(Goods),
		Handlers:   make([]Handler, 0),
		Calculator: roman.DefaultCalculator,
	}
}
