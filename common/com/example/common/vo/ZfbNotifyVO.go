package vo

type ZfbNotifyVO struct {
	OutTradeNo     string `json:"out_trade_no"`     //商户订单号
	TradeStatus    string `json:"trade_status"`     //交易状态
	TotalAmount    string `json:"total_amount"`     //交易金额
	BuyerPayAmount string `json:"buyer_pay_amount"` //付款金额
	ReceiptAmount  string `json:"receipt_amount"`   //实收金额
	Msg            string `json:"msg"`              //处理结果的描述，信息来自于code返回结果的描述
	Subject        string `json:"subject"`          //商品标题
	Body           string `json:"body"`             //商品描述
}
