package dao

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	_ "github.com/zhuxiujia/GoMybatisMall/core/com/example/dao/statik"
)

type MallOrderMapper struct {
	InsertTemplete                    func(arg model.MallOrder) error
	UpdateTemplete                    func(arg model.MallOrder) (int64, error)
	DeleteTemplete                    func(id string) error                    `mapperParams:"id"`
	SelectTemplete                    func(id string) (model.MallOrder, error) `mapperParams:"id"`
	SelectByCondition                 func(arg service.MallOrderPageDTO) ([]model.MallOrder, error)
	CountByCondition                  func(arg service.MallOrderPageDTO) (int64, error)
	CountByUserIdAndIntegralProductId func(arg service.MallOrderPageDTO) (int64, error)
}

func (it MallOrderMapper) New() MallOrderMapper {
	core_util.LoadMapper(&it)
	return it
}
