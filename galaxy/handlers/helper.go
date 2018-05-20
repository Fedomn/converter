package handlers

import (
	. "fedomn/converter/galaxy/models"
	"strings"
)

func calcAliasDecimal(aliasStr string, g *Guider) (int, error) {
	aliasAry := strings.Split(aliasStr, " ")

	alias := make([]AliasSymbol, 0)
	for _, each := range aliasAry {
		alias = append(alias, AliasSymbol(each))
	}

	// 计算组合十进制数
	mapSymbols := g.ConvertAliasMapSymbol(alias)
	return g.Calculator.CalcDecimal(mapSymbols)
}
