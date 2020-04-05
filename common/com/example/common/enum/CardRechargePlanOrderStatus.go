package enum

import "fmt"

type CardRechargePlanOrderStatus int

func (it CardRechargePlanOrderStatus) ToInt() int {
	return int(it)
}
func (it CardRechargePlanOrderStatus) ToString() string {
	switch it {
	case -1:
		return "失败"
	case 1:
		return "已创建"
	case 2:
		return "支付中"
	case 3:
		return "预支付"
	case 4:
		return "充值中"
	case 5:
		return "已充值"
	default:
		panic("未知的充值计划类型")
	}
	return ""
}

func (it CardRechargePlanOrderStatus) New(status int) CardRechargePlanOrderStatus {
	switch status {
	case -1:
		return CardRechargePlanOrderStatus_Fail
	case 1:
		return CardRechargePlanOrderStatus_Created
	case 2:
		return CardRechargePlanOrderStatus_Paying
	case 3:
		return CardRechargePlanOrderStatus_Success
	case 4:
		return CardRechargePlanOrderStatus_Recharing
	case 5:
		return CardRechargePlanOrderStatus_Recharged
	default:
		panic("未知的充值计划类型：" + fmt.Sprint(status))
	}
	return it
}

const (
	CardRechargePlanOrderStatus_Fail      CardRechargePlanOrderStatus = -1 //支付失败或取消
	CardRechargePlanOrderStatus_Created   CardRechargePlanOrderStatus = 1  //已创建
	CardRechargePlanOrderStatus_Paying    CardRechargePlanOrderStatus = 2  //支付中
	CardRechargePlanOrderStatus_Success   CardRechargePlanOrderStatus = 3  //支付成功
	CardRechargePlanOrderStatus_Recharing CardRechargePlanOrderStatus = 4  //充值中
	CardRechargePlanOrderStatus_Recharged CardRechargePlanOrderStatus = 5  //已充值

)
