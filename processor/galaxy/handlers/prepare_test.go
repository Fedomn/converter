package handlers_test

import (
	"os"
	"sync"
	"testing"

	"github.com/fedomn/converter/calculator/roman"
	. "github.com/fedomn/converter/processor/galaxy/handlers"
	. "github.com/fedomn/converter/processor/galaxy/models"
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
		Alias:      sync.Map{},
		Goods:      sync.Map{},
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
