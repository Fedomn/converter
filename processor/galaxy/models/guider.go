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

func (g *Guider) SetCalculator(c calculator.Calculator) {
	g.Calculator = c
}

func (g *Guider) Process(context string) (input, output string) {
	rsp := g.Handle(context)
	if rsp.Err != nil {
		return context, rsp.Err.Error()
	}
	return context, rsp.Res
}
