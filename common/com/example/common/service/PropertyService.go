package service

import "github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"

type UpdateAmountDTO struct {
	OrderId         string        `json:"order_id"`
	AddReduceAmount int           `json:"add_reduce_amount"`
	PropertyVO      vo.PropertyVO `json:"property_vo"`
}

type PropertyService struct {
	Insert       func(arg vo.PropertyVO) error
	Update       func(arg vo.PropertyVO) (int64, error)
	UpdateAmount func(arg UpdateAmountDTO) (int64, error)
	FindByUser   func(user_id string) (vo.PropertyVO, error)
	Find         func(id string) (vo.PropertyVO, error)
}
