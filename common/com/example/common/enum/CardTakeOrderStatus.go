package enum

type CardTakeOrderStatus int

func (it CardTakeOrderStatus) ToInt() int {
	return int(it)
}

const (
	CardTakeOrderStatus_UnPay   CardTakeOrderStatus = -1 //未支付
	CardTakeOrderStatus_Disable CardTakeOrderStatus = 0  //已支付，等待绑卡
	CardTakeOrderStatus_Enable  CardTakeOrderStatus = 1  //已生效
)
