package dao

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
)

type PropertyMapper struct {
	SelectTemplete func(id string) (model.Property, error) `mapperParams:"id"`
	InsertTemplete func(arg model.Property) error
	UpdateTemplete func(arg model.Property) error
	SelectByUserId func(user_id string) (model.Property, error) `mapperParams:"user_id"`
}

func (it PropertyMapper) New() PropertyMapper {
	core_util.LoadMapper(&it)
	return it
}
