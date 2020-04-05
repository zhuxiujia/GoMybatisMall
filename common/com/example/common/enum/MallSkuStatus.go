package enum

import "fmt"

//状态
type MallSkuStatus int

const (
	MallSkuStatus_OffLine MallSkuStatus = iota
	MallSkuStatus_OnLine
)

func NewMallSkuStatus(t int) MallSkuStatus {
	switch t {
	case 0:
		return MallSkuStatus_OffLine
	case 1:
		return MallSkuStatus_OnLine
	default:
		panic("未知的商城订单类型：" + fmt.Sprint(t))
	}
}
