package enum

import (
	"fmt"
)

//第三方支付 订单状态
type NotifyStatus int

//直接创建 订单状态
func NewNotifyStatus(t int) NotifyStatus {
	switch t {
	case -1:
		return NotifyStatus_FAIL
	case 1:
		return NotifyStatus_SUCCESS
	case 3:
		return NotifyStatus_PAYING
	}
	panic("不支持卡状态! type=" + fmt.Sprint(t))
}

const (
	NotifyStatus_FAIL    NotifyStatus = -1 //失败
	NotifyStatus_PAYING  NotifyStatus = 0  //进行中
	NotifyStatus_SUCCESS NotifyStatus = 1  //成功
)
