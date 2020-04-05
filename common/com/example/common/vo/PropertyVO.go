package vo

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
)

type PropertyVO struct {
	model.Property

	LastRecharge       int    `json:"last_recharge"`
	NextRechargeDate   string `json:"next_recharge_date"`
	NextRechargeAmount int    `json:"next_recharge_amount"`
}
