package handlers

import (
	. "fedomn/converter/processor/galaxy/models"
	"regexp"
	"strconv"
	"strings"
)

var goodsRegexp = regexp.MustCompile(`^(.+) (\S+) is (\d+) (\S+)$`)

type GoodsHandler struct{}

func (GoodsHandler) Validate(context string, g *Guider) error {
	if !goodsRegexp.MatchString(context) {
		return NotMatchErr
	}

	findAry := goodsRegexp.FindStringSubmatch(context)
	aliasAry := strings.Split(findAry[1], " ")
	for _, each := range aliasAry {
		if _, ok := g.LoadAlias(AliasSymbol(each)); !ok {
			return UnknownErr
		}
	}

	return nil
}

func (gh GoodsHandler) Handle(context string, g *Guider) HandlerRsp {
	findAry := goodsRegexp.FindStringSubmatch(context)

	aliasStr, goodsSymbol, goodsPriceStr, goodsUnitSymbol :=
		findAry[1], GoodsSymbol(findAry[2]), findAry[3], GoodsUnitSymbol(findAry[4])

	// 计算商品个数
	goodsNum, err := calcAliasDecimal(aliasStr, g)
	if err != nil {
		return HandlerRsp{Context: context, Err: CalcErr}
	}

	// 计算商品单价
	priceFloat, _ := strconv.ParseFloat(goodsPriceStr, 64)
	totalPrice := GoodsPrice(priceFloat)

	floatStr := strconv.FormatFloat(float64(totalPrice)/float64(goodsNum), 'f', 2, 64)
	goodsUnitPrice, _ := strconv.ParseFloat(floatStr, 64)

	info := GoodsInfo{
		Symbol: goodsSymbol,
		Price:  GoodsPrice(goodsUnitPrice),
		Unit:   goodsUnitSymbol,
	}

	g.StoreGoods(goodsSymbol, info)

	return HandlerRsp{Context: context}
}
