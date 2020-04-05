package model

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/enum"
	"time"
)

type MallOrder struct {
	Id                 string               `json:"id" gm:"id"`
	UserId             string               `json:"user_id"`
	ReceiveName        string               `json:"receive_name"`
	ReceivePhone       string               `json:"receive_phone"`
	ReceiveAddress     string               `json:"receive_address"`
	SkuId              string               `json:"sku_id"`
	SkuNum             int                  `json:"sku_num"`
	SkuSpecificationId string               `json:"sku_specification_id"`
	ExpressNum         string               `json:"express_num"`
	ExpressName        string               `json:"express_name"`
	OrderAmount        int                  `json:"order_amount"`
	Status             enum.MallOrderStatus `json:"status"`
	PayLink            string               `json:"pay_link"` //支付链接
	PayType            string               `json:"pay_type"` //支付类型  zfb支付宝  wx微信
	Remark             string               `json:"remark"`

	Version    int       `json:"version" gm:"version"`
	CreateTime time.Time `json:"create_time"`
	DeleteFlag int       `json:"delete_flag" gm:"logic"`
}
