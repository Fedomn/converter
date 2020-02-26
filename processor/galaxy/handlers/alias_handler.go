package handlers

import (
	"regexp"

	. "github.com/fedomn/converter/processor/galaxy/models"
)

var aliasRegexp = regexp.MustCompile(`^(\S+) is (\S)$`)

type AliasHandler struct{}

func (AliasHandler) Validate(context string, g *Guider) error {
	if aliasRegexp.MatchString(context) {
		return nil
	}
	return NotMatchErr
}

func (AliasHandler) Handle(context string, g *Guider) HandlerRsp {
	findAry := aliasRegexp.FindStringSubmatch(context)
	g.StoreAlias(AliasSymbol(findAry[1]), AliasMapSymbol(findAry[2]))
	return HandlerRsp{Context: context}
}
