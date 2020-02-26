package galaxy

import (
	"sync"

	"github.com/fedomn/converter/calculator/roman"
	. "github.com/fedomn/converter/processor/galaxy/handlers"
	. "github.com/fedomn/converter/processor/galaxy/models"
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
