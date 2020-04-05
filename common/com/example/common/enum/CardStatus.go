package enum

import "fmt"

type CardStatus int

func (it CardStatus) New(t int) CardStatus {
	switch t {
	case 0:
		return CardStatus_Disable
	case 1:
		return CardStatus_Enable

	}
	panic("不支持卡状态! type=" + fmt.Sprint(t))
}

const (
	CardStatus_Disable CardStatus = 0
	CardStatus_Enable  CardStatus = 1
)
