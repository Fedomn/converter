package models

import "fedomn/converter/calculator"

type (
	AliasSymbol    string
	AliasMapSymbol string
	Alias          map[AliasSymbol]AliasMapSymbol

	GoodsPrice      float64
	GoodsSymbol     string
	GoodsUnitSymbol string
	GoodsInfo       struct {
		Symbol GoodsSymbol
		Price  GoodsPrice
		Unit   GoodsUnitSymbol
	}
	Goods map[GoodsSymbol]GoodsInfo

	Guider struct {
		Alias      Alias
		Goods      Goods
		Handlers   []Handler
		Calculator calculator.Calculator
	}
)

func (g *Guider) ConvertAliasMapSymbol(alias []AliasSymbol) []AliasMapSymbol {
	mapSymbols := make([]AliasMapSymbol, 0)
	for _, each := range alias {
		mapSymbols = append(mapSymbols, g.Alias[each])
	}
	return mapSymbols
}

func (g *Guider) Use(handlers ...Handler) {
	g.Handlers = append(g.Handlers, handlers...)
}

func (g *Guider) Handle(context string) HandlerRsp {
	for _, handler := range g.Handlers {
		err := handler.Validate(context, g)
		if err == NotMatchErr {
			continue
		}
		if err != nil {
			return HandlerRsp{Err: err}
		}
		return handler.Handle(context, g)
	}
	return HandlerRsp{Err: UnknownErr}
}
