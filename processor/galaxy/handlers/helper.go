package handlers

import (
	. "fedomn/converter/processor/galaxy/models"
	"strings"
)

func convertAliasMapSymbol(alias []AliasSymbol, g *Guider) []AliasMapSymbol {
	mapSymbols := make([]AliasMapSymbol, 0)
	for _, each := range alias {
		value, _ := g.LoadAlias(AliasSymbol(each))
		mapSymbols = append(mapSymbols, value)
	}
	return mapSymbols
}

func calcAliasDecimal(aliasStr string, g *Guider) (int, error) {
	aliasAry := strings.Split(aliasStr, " ")

	alias := make([]AliasSymbol, 0)
	for _, each := range aliasAry {
		alias = append(alias, AliasSymbol(each))
	}

	// 计算组合十进制数
	mapSymbols := convertAliasMapSymbol(alias, g)
	return g.Calculator.CalcDecimal(mapSymbols)
}
