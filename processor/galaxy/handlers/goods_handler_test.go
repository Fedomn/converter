package handlers_test

import (
	"fmt"
	"testing"

	. "github.com/fedomn/converter/processor/galaxy/models"
	. "github.com/fedomn/converter/util"
)

func TestGoodsHandlerValidate(t *testing.T) {
	prepareAliasData()

	var tests = []struct {
		context string
		wat     error
	}{
		{"", NotMatchErr},
		{"glob xxx Silver is 34 Credits", UnknownErr},
		{"glob glob Silver is 34 Credits", nil},
	}
	for _, tt := range tests {
		got := goodsHandler.Validate(tt.context, guilder)
		msg := fmt.Sprintf("contxt: %s", tt.context)
		Equals(t, msg, tt.wat, got)
	}
}

func TestGoodsHandlerHandle(t *testing.T) {
	prepareAliasData()

	var tests = []struct {
		context     string
		goodsSymbol GoodsSymbol
		wat         interface{}
	}{
		{"glob tegj Silver is 34 Credits", "Silver", CalcErr},
		{"glob glob Silver is 34 Credits", "Silver", GoodsInfo{Symbol: "Silver", Price: 17, Unit: "Credits"}},
		{"glob prok Gold is 57800 Credits", "Gold", GoodsInfo{Symbol: "Gold", Price: 14450, Unit: "Credits"}},
		{"pish pish Iron is 3910 Credits", "Iron", GoodsInfo{Symbol: "Iron", Price: 195.5, Unit: "Credits"}},
	}
	for _, tt := range tests {
		handleRsp := goodsHandler.Handle(tt.context, guilder)
		msg := fmt.Sprintf("contxt: %s", tt.context)
		if handleRsp.Err != nil {
			Equals(t, msg, tt.wat, handleRsp.Err)
		} else {
			info, _ := guilder.LoadGoods(tt.goodsSymbol)
			Equals(t, msg, tt.wat, info)
		}
	}
}
