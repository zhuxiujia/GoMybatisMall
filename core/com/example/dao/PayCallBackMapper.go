package dao

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
)

type PayCallBackMapper struct {
	InsertTemplete      func(arg model.PayCallBack) error
	SelectPageTemplete  func(page int, size int) ([]model.PayCallBack, error) `mapperParams:"page,size"`
	SelectCountTemplete func() (int64, error)
}

func (it PayCallBackMapper) New() PayCallBackMapper {
	core_util.LoadMapper(&it)
	return it
}
