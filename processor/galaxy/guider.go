package galaxy

import (
	"fedomn/converter/calculator/roman"
	. "fedomn/converter/processor/galaxy/handlers"
	. "fedomn/converter/processor/galaxy/models"
	"sync"
)

var DefaultGuider *Guider

func init() {
	DefaultGuider = initGuider()
	DefaultGuider.Use(AliasHandler{})
	DefaultGuider.Use(GoodsHandler{})
	DefaultGuider.Use(HowMuchHandler{})
	DefaultGuider.Use(HowManyUnitHandler{})
	DefaultGuider.SetCalculator(roman.DefaultCalculator)
}

func initGuider() *Guider {
	return &Guider{
		Alias:    sync.Map{},
		Goods:    sync.Map{},
		Handlers: make([]Handler, 0),
	}
}
