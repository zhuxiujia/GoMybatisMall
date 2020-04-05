package enum

import (
	"fmt"
	"strings"
)

type CardType int

func (it CardType) FromCardNum(cardNum string) CardType {
	if cardNum == "" {
		panic("卡号错误，不能为空")
	}
	if strings.Index(cardNum, "100011") == 0 {
		return CardType_ZSH
	} else {
		return CardType_ZSY
	}
}

func (it CardType) New(t int) CardType {
	switch t {
	case 1:
		return CardType_ZSH
	case 2:
		return CardType_ZSY
	}
	panic("不支持卡类型! type=" + fmt.Sprint(t))
}

const (
	CardType_ZSH CardType = 1
	CardType_ZSY CardType = 2
)
