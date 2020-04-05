package vo

//微信支付通知结果
type WXNotifyVO struct {
	ReturnCode string `json:"return_code"`
	ReturnMsg  string `json:"return_msg"`

	//以下字段在return_code为SUCCESS的时候有返回
	Appid         string `json:"appid"`          //应用ID
	MchId         string `json:"mch_id"`         //商户号
	DeviceInfo    string `json:"device_info"`    //设备号
	NonceStr      string `json:"nonce_str"`      //随机字符串
	Sign          string `json:"sign"`           //签名
	ResultCode    string `json:"result_code"`    //业务结果
	ErrCode       string `json:"err_code"`       //错误代码
	ErrCodeDes    string `json:"err_code_des"`   //错误代码描述
	Openid        string `json:"openid"`         //用户标识
	IsSubscribe   string `json:"is_subscribe"`   //是否关注公众账号
	TradeType     string `json:"trade_type"`     //交易类型
	BankType      string `json:"bank_type"`      //付款银行
	TotalFee      string `json:"total_fee"`      //总金额
	FeeType       string `json:"fee_type"`       //货币种类
	CashFee       string `json:"cash_fee"`       //现金支付金额
	CashFeeType   string `json:"cash_fee_type"`  //现金支付货币类型
	TransactionId string `json:"transaction_id"` //微信支付订单号
	OutTradeNo    string `json:"out_trade_no"`   //商户订单号
	Attach        string `json:"attach"`         //商家数据包
	TimeEnd       string `json:"time_end"`       //支付完成时间
}
