package handlers_test

import (
	. "fedomn/converter/galaxy/models"
	. "fedomn/converter/util"
	"fmt"
	"testing"
)

func TestHowManyHandlerValidate(t *testing.T) {
	prepareAliasAndGoodsData()

	var tests = []struct {
		context string
		wat     error
	}{
		{"", NotMatchErr},
		{"how much is pish xxx glob glob ?", UnknownErr},
		{"how much is pish tegj glob glob ?", nil},
	}
	for _, tt := range tests {
		got := howMuchHandler.Validate(tt.context, guilder)
		msg := fmt.Sprintf("contxt: %s", tt.context)
		Equals(t, msg, tt.wat, got)
	}
}

func TestHowManyHandlerHandle(t *testing.T) {
	prepareAliasAndGoodsData()

	var tests = []struct {
		context string
		wat     interface{}
	}{
		{"how much is glob tegj ?", CalcErr},
		{"how much is pish tegj glob glob ?", "pish tegj glob glob is 42"},
	}
	for _, tt := range tests {
		handleRsp := howMuchHandler.Handle(tt.context, guilder)
		msg := fmt.Sprintf("contxt: %s", tt.context)
		if handleRsp.Err != nil {
			Equals(t, msg, tt.wat, handleRsp.Err)
		} else {
			Equals(t, msg, tt.wat, handleRsp.Res)
		}
	}
}
