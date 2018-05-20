package handlers

import (
	. "fedomn/converter/galaxy/models"
	"fmt"
	"regexp"
	"strings"
)

var howMuchRegexp = regexp.MustCompile(`^how much is (.+) \?$`)

type HowMuchHandler struct{}

func (HowMuchHandler) Validate(context string, g *Guider) error {
	if !howMuchRegexp.MatchString(context) {
		return NotMatchErr
	}

	findAry := howMuchRegexp.FindStringSubmatch(context)
	aliasAry := strings.Split(findAry[1], " ")
	for _, each := range aliasAry {
		if _, ok := g.Alias[AliasSymbol(each)]; !ok {
			return UnknownErr
		}
	}

	return nil
}

func (HowMuchHandler) Handle(context string, g *Guider) HandlerRsp {
	findAry := howMuchRegexp.FindStringSubmatch(context)
	aliasStr := findAry[1]

	// 计算Alias的十进制数
	res, err := calcAliasDecimal(aliasStr, g)
	if err != nil {
		return HandlerRsp{Context: context, Err: CalcErr}
	}

	answer := fmt.Sprintf("%s is %d", aliasStr, res)
	return HandlerRsp{Context: context, Res: answer}
}