package handlers_test

import (
	. "fedomn/converter/galaxy/models"
	. "fedomn/converter/util"
	"fmt"
	"testing"
)

func TestAliasHandlerValidate(t *testing.T) {
	var tests = []struct {
		context string
		wat     error
	}{
		{"", NotMatchErr},
		{"glob is a I", NotMatchErr},
		{"glob is I", nil},
	}
	for _, tt := range tests {
		got := aliasHandler.Validate(tt.context, guilder)
		msg := fmt.Sprintf("contxt: %s", tt.context)
		Equals(t, msg, tt.wat, got)
	}
}

func TestAliasHandlerHandle(t *testing.T) {
	var tests = []struct {
		context        string
		aliasSymbol    AliasSymbol
		AliasMapSymbol AliasMapSymbol
	}{
		{"glob is I", "glob", "I"},
		{"prok is V", "prok", "V"},
		{"pish is X", "pish", "X"},
		{"tegj is L", "tegj", "L"},
	}
	for _, tt := range tests {
		handleRsp := aliasHandler.Handle(tt.context, guilder)
		msg := fmt.Sprintf("contxt: %s", tt.context)
		Equals(t, msg, nil, handleRsp.Err)
		Equals(t, msg, tt.AliasMapSymbol, guilder.Alias[tt.aliasSymbol])
	}
}
