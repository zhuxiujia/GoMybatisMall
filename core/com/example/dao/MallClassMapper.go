package dao

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	_ "github.com/zhuxiujia/GoMybatisMall/core/com/example/dao/statik"
)

type MallClassMapper struct {
	InsertTemplete func(arg model.MallClass) error
	UpdateTemplete func(arg model.MallClass) error
	DeleteTemplete func(id string) error                    `mapperParams:"id"`
	SelectTemplete func(id string) (model.MallClass, error) `mapperParams:"id"`

	SelectByCondition func(arg service.MallClassPageDTO) ([]model.MallClass, error)
	CountByCondition  func(arg service.MallClassPageDTO) (int64, error)
	FindAll           func() ([]model.MallClass, error)
}

func (it MallClassMapper) New() MallClassMapper {
	core_util.LoadMapper(&it)
	return it
}
