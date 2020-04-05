package vo

import "github.com/zhuxiujia/GoMybatisMall/common/com/example/common/enum"

//抽象订单对象，代替支付宝和微信等的支付对象
type AbstractNotifyVO struct {
	OutTradeNo string            `json:"out_trade_no"` //订单号
	Msg        string            `json:"msg"`          //处理结果的描述，信息来自于code返回结果的描述
	Status     enum.NotifyStatus `json:"status"`       //状态,-1失败，0支付中，1成功
}

//创建抽象订单对象-支付宝
func NewAbstractNotifyVOZfb(arg ZfbNotifyVO) AbstractNotifyVO {
	//第三方状态判断
	var stats = enum.NotifyStatus_FAIL
	switch arg.TradeStatus {
	case "TRADE_SUCCESS":
		stats = enum.NotifyStatus_SUCCESS
	case "WAIT_BUYER_PAY":
		stats = enum.NotifyStatus_PAYING
	default:
		stats = enum.NotifyStatus_FAIL
	}
	var notifyVO = AbstractNotifyVO{
		OutTradeNo: arg.OutTradeNo,
		Msg:        arg.Msg,
		Status:     stats,
	}
	return notifyVO
}

//TODO 创建抽象订单对象-微信支付
func NewAbstractNotifyVOWx(arg WXNotifyVO) AbstractNotifyVO {
	//第三方状态判断
	var stats enum.NotifyStatus
	switch arg.ResultCode {
	case "SUCCESS":
		stats = enum.NotifyStatus_SUCCESS
	case "FAIL":
		stats = enum.NotifyStatus_FAIL
	default:
		stats = enum.NotifyStatus_PAYING
	}
	var notifyVO = AbstractNotifyVO{
		OutTradeNo: arg.OutTradeNo,
		Msg:        arg.ReturnMsg,
		Status:     stats,
	}
	return notifyVO
}
