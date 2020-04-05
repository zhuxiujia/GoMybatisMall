package dao

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
)

type UserAddressMapper struct {
	SelectPageTemplete  func(user_id string, phone string, real_name string, page int, size int) ([]model.UserAddress, error) `mapperParams:"user_id,phone,real_name,page,size"`
	SelectCountTemplete func(user_id string, phone string, real_name string) (int64, error)                                   `mapperParams:"user_id,phone,real_name"`

	InsertTemplete func(arg model.UserAddress) error
	UpdateTemplete func(arg model.UserAddress) error
	DeleteTemplete func(id string) error                      `mapperParams:"id"`
	SelectTemplete func(id string) (model.UserAddress, error) `mapperParams:"id"`
}

func (it UserAddressMapper) New() UserAddressMapper {
	core_util.LoadMapper(&it)
	return it
}
