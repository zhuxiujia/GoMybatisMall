package core_util

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"strings"
)

var yuan, _ = decimal.NewFromString("100")

//金额转换
func ToThirdPayAmount(PayType string, OrderAmount int) string {
	//金额过滤
	var payRealAmount string
	if PayType == "zfb" {
		OrderAmountDec, _ := decimal.NewFromString(fmt.Sprint(OrderAmount))
		OrderAmountDec = OrderAmountDec.Div(yuan)
		payRealAmount = OrderAmountDec.String()
	} else {
		payRealAmount = fmt.Sprint(OrderAmount)
	}
	return payRealAmount
}

func FilterAmount(supers vo.KVVO, receive_phone string, payType string, thirdPayAmount string) string {
	if supers.Id != "" && supers.Value != "" {
		//filter
		if strings.Contains(supers.Value, receive_phone) {
			return ToThirdPayAmount(payType, 1)
		}
	}
	return thirdPayAmount
}
