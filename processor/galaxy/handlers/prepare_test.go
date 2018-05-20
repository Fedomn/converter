package handlers_test

import (
	"fedomn/converter/calculator/roman"
	. "fedomn/converter/processor/galaxy/handlers"
	. "fedomn/converter/processor/galaxy/models"
	"os"
	"testing"
)

var aliasHandler AliasHandler
var goodsHandler GoodsHandler
var howMuchHandler HowMuchHandler
var howManyUnitHandler HowManyUnitHandler

var guilder *Guider

func TestMain(m *testing.M) {
	aliasHandler = AliasHandler{}
	goodsHandler = GoodsHandler{}
	howMuchHandler = HowMuchHandler{}
	howManyUnitHandler = HowManyUnitHandler{}

	guilder = &Guider{
		Alias:      make(Alias),
		Goods:      make(Goods),
		Handlers:   make([]Handler, 0),
		Calculator: roman.DefaultCalculator,
	}
	os.Exit(m.Run())
}

func prepareAliasData() {
	var inputs = []string{
		"glob is I",
		"prok is V",
		"pish is X",
		"tegj is L",
	}
	for _, each := range inputs {
		aliasHandler.Handle(each, guilder)
	}
}

func prepareAliasAndGoodsData() {
	prepareAliasData()

	var inputs = []string{
		"glob glob Silver is 34 Credits",
		"glob prok Gold is 57800 Credits",
		"pish pish Iron is 3910 Credits",
	}
	for _, each := range inputs {
		goodsHandler.Handle(each, guilder)
	}
}
