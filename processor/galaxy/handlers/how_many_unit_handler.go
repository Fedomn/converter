package handlers

import (
	"fmt"
	"regexp"
	"strings"

	. "github.com/fedomn/converter/processor/galaxy/models"
)

var howManyUnitRegexp = regexp.MustCompile(`^how many (\S+) is (.+) (\S+) \?$`)

type HowManyUnitHandler struct{}

func (HowManyUnitHandler) Validate(context string, g *Guider) error {
	if !howManyUnitRegexp.MatchString(context) {
		return NotMatchErr
	}

	findAry := howManyUnitRegexp.FindStringSubmatch(context)
	goodsUnitSymbol, aliasStr, goodsSymbol := findAry[1], findAry[2], findAry[3]

	// 验证商品合法性
	if _, ok := g.LoadGoods(GoodsSymbol(goodsSymbol)); !ok {
		return UnknownErr
	}

	// 验证商品单位合法性
	info, _ := g.LoadGoods(GoodsSymbol(goodsSymbol))
	if info.Unit != GoodsUnitSymbol(goodsUnitSymbol) {
		return UnknownErr
	}

	// 验证Alias合法性
	aliasAry := strings.Split(aliasStr, " ")
	for _, each := range aliasAry {
		if _, ok := g.LoadAlias(AliasSymbol(each)); !ok {
			return UnknownErr
		}
	}

	return nil
}

func (HowManyUnitHandler) Handle(context string, g *Guider) HandlerRsp {
	findAry := howManyUnitRegexp.FindStringSubmatch(context)
	goodsUnitSymbol, aliasStr, goodsSymbol := GoodsUnitSymbol(findAry[1]), findAry[2], GoodsSymbol(findAry[3])

	// 计算商品个数
	goodsNum, err := calcAliasDecimal(aliasStr, g)
	if err != nil {
		return HandlerRsp{Context: context, Err: CalcErr}
	}

	// 计算总价
	info, _ := g.LoadGoods(goodsSymbol)
	totalPrice := float64(info.Price) * float64(goodsNum)

	answer := fmt.Sprintf("%s %s is %.0f %s", aliasStr, goodsSymbol, totalPrice, goodsUnitSymbol)

	return HandlerRsp{Context: context, Res: answer}
}
