package service

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"time"
)

type MallOrderPageDTO struct {
	vo.Pageable
	Fuzzy     string     `json:"fuzzy"`
	SkuId     string     `json:"sku_id"`
	UserId    string     `json:"user_id"`
	Status    *int       `json:"status"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
}

type MallOrderDTO struct {
	SkuId              string `json:"sku_id"`
	UserId             string `json:"user_id"`
	SkuSpecificationId string `json:"sku_specification_id"`

	ReceiveName    string `json:"receive_name"`
	ReceivePhone   string `json:"receive_phone"`
	ReceiveAddress string `json:"receive_address"`
	SkuNum         int    `json:"sku_num"`
	Amount         int    `json:"amount"`
	Remark         string `json:"remark"`
	PayType        string `json:"pay_type"` //支付类型  zfb支付宝  wx微信
	Debug          string `json:"debug"`
}

type MallOrderService struct {
	Add    func(arg model.MallOrder) error
	Update func(arg model.MallOrder) (int64, error)
	Cancel func(id string) error
	Delete func(id string) error
	Find   func(id string) (vo.MallOrderVO, error)
	//获取积分商城订单记录
	Page             func(arg MallOrderPageDTO) (vo.PageVO, error)
	Order            func(arg MallOrderDTO) (orderVO vo.MallOrderVO, e error)
	SyncOrderByQueue func(arg vo.AbstractNotifyVO) error
}
