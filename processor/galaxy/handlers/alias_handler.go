package handlers

import (
	. "fedomn/converter/processor/galaxy/models"
	"regexp"
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
	g.Alias[AliasSymbol(findAry[1])] = AliasMapSymbol(findAry[2])
	return HandlerRsp{Context: context}
}
