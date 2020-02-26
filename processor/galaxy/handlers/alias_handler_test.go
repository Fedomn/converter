package handlers_test

import (
	"fmt"
	"testing"

	. "github.com/fedomn/converter/processor/galaxy/models"
	. "github.com/fedomn/converter/util"
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
		mapSymbol, _ := guilder.LoadAlias(tt.aliasSymbol)
		Equals(t, msg, tt.AliasMapSymbol, mapSymbol)
	}
}
