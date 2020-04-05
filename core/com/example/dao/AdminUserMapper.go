package dao

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
)

type AdminUserMapper struct {
	InsertTemplete func(arg model.AdminUser) error
	UpdateTemplete func(arg model.AdminUser) error
	DeleteTemplete func(id string) error                       `mapperParams:"id"`
	SelectByPhone  func(phone string) (model.AdminUser, error) `mapperParams:"phone"`
	SelectTemplete func(id string) (model.AdminUser, error)    `mapperParams:"id"`

	SelectPageTemplete  func(phone string, enable *int, page int, size int) ([]model.AdminUser, error) `mapperParams:"phone,enable,page,size"`
	SelectCountTemplete func(phone string, enable *int) (int64, error)                                 `mapperParams:"phone,enable"`
}

func (it AdminUserMapper) New() AdminUserMapper {
	core_util.LoadMapper(&it)
	return it
}
