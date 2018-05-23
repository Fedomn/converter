package models

import (
	"fedomn/converter/calculator"
	"sync"
)

type (
	AliasSymbol    string
	AliasMapSymbol string

	GoodsPrice      float64
	GoodsSymbol     string
	GoodsUnitSymbol string
	GoodsInfo       struct {
		Symbol GoodsSymbol
		Price  GoodsPrice
		Unit   GoodsUnitSymbol
	}

	Guider struct {
		Alias      sync.Map
		Goods      sync.Map
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

func (g *Guider) Process(context string) string {
	rsp := g.Handle(context)
	if rsp.Err != nil {
		return rsp.Err.Error()
	}
	return rsp.Res
}

func (g *Guider) StoreAlias(aliasSymbol AliasSymbol, mapSymbol AliasMapSymbol) {
	g.Alias.Store(aliasSymbol, mapSymbol)
}

func (g *Guider) LoadAlias(aliasSymbol AliasSymbol) (AliasMapSymbol, bool) {
	if value, ok := g.Alias.Load(aliasSymbol); ok {
		return value.(AliasMapSymbol), true
	}
	return "", false
}

func (g *Guider) StoreGoods(goodsSymbol GoodsSymbol, info GoodsInfo) {
	g.Goods.Store(goodsSymbol, info)
}

func (g *Guider) LoadGoods(goodsSymbol GoodsSymbol) (GoodsInfo, bool) {
	if value, ok := g.Goods.Load(goodsSymbol); ok {
		return value.(GoodsInfo), true
	}
	return GoodsInfo{}, false
}
