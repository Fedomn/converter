package handlers_test

import (
	. "fedomn/converter/galaxy/models"
	. "fedomn/converter/util"
	"fmt"
	"testing"
)

func TestHowManyUnitHandlerValidate(t *testing.T) {
	prepareAliasAndGoodsData()

	var tests = []struct {
		context string
		wat     error
	}{
		{"", NotMatchErr},
		{"how many Credits is glob prok xxx ?", UnknownErr},
		{"how many Credits is glob prok Silver ?", nil},
	}
	for _, tt := range tests {
		got := howManyUnitHandler.Validate(tt.context, guilder)
		msg := fmt.Sprintf("contxt: %s", tt.context)
		Equals(t, msg, tt.wat, got)
	}
}

func TestHowManyUnitHandlerHandle(t *testing.T) {
	prepareAliasAndGoodsData()

	var tests = []struct {
		context string
		wat     interface{}
	}{
		{"how many Credits is glob tegj Silver ?", CalcErr},
		{"how many Credits is glob prok Silver ?", "glob prok Silver is 68 Credits"},
		{"how many Credits is glob prok Gold ?", "glob prok Gold is 57800 Credits"},
		{"how many Credits is glob prok Iron ?", "glob prok Iron is 782 Credits"},
	}
	for _, tt := range tests {
		handleRsp := howManyUnitHandler.Handle(tt.context, guilder)
		msg := fmt.Sprintf("contxt: %s", tt.context)
		if handleRsp.Err != nil {
			Equals(t, msg, tt.wat, handleRsp.Err)
		} else {
			Equals(t, msg, tt.wat, handleRsp.Res)
		}
	}
}
