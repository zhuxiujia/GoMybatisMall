package core_util

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"testing"
	"time"
)

//根据油卡面值拆单,拆分为500,100,1000
func TestCreatePlanOrder_test(t *testing.T) {
	var arg = vo.CardRechargeOrderVO{
		CardRechargeOrder: model.CardRechargeOrder{
			Id:            "1",
			UserId:        "",
			ProductId:     "",
			CardId:        "",
			RedPacketId:   "",
			TotalAmount:   0,
			RealAmount:    0,
			EachAmount:    1200 * 100,
			RemainAmount:  0,
			TotalPayTime:  4,
			RemainPayTime: 0,
			PayedTime:     0,
			ExtraAmount:   0,
			Status:        0,
			PayLink:       "",
			PayType:       "",
			PayRemark:     "",
			Version:       0,
			CreateTime:    time.Time{},
			DeleteFlag:    1,
		},
		UserVO: &vo.UserVO{
			User: model.User{
				Id:         "0",
				Phone:      "18969542172",
				DeleteFlag: 1,
			},
			UserAddressVO: nil,
			PropertyVO:    nil,
		},
	}
	var plans = CreatePlanOrder(arg)
	for _, item := range plans {
		println(item.Amount)
	}
}
