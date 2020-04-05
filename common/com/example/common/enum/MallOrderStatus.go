package enum

import "fmt"

//商城订单状态
type MallOrderStatus int

func NewMallOrderStatus(t int) MallOrderStatus {
	switch t {
	case -1:
		return MallOrderStatus_Fail
	case 0:
		return MallOrderStatus_Paying
	case 1:
		return MallOrderStatus_Redy
	case 2:
		return MallOrderStatus_ING
	case 3:
		return MallOrderStatus_COMPLETE
	case 4:
		return MallOrderStatus_Cancel
	default:
		panic("未知的商城订单类型：" + fmt.Sprint(t))
	}
}

const (
	MallOrderStatus_Fail     MallOrderStatus = -1 //失败
	MallOrderStatus_Paying   MallOrderStatus = 0  //支付中
	MallOrderStatus_Redy     MallOrderStatus = 1  //待发货
	MallOrderStatus_ING      MallOrderStatus = 2  //已发货
	MallOrderStatus_COMPLETE MallOrderStatus = 3  //已完成
	MallOrderStatus_Cancel   MallOrderStatus = 4  //已取消

)
