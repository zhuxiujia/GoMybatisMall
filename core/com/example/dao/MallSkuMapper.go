package dao

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	_ "github.com/zhuxiujia/GoMybatisMall/core/com/example/dao/statik"
)

type MallSkuMapper struct {
	InsertTemplete    func(arg model.MallSku) error
	UpdateTemplete    func(arg model.MallSku) (int64, error)
	DeleteTemplete    func(id string) error                       `mapperParams:"id"`
	SelectTemplete    func(id string) (model.MallSku, error)      `mapperParams:"id"`
	SelectByIds       func(ids []string) ([]model.MallSku, error) `mapperParams:"ids"`
	SelectByCondition func(arg service.MallSkuPageDTO) ([]model.MallSku, error)
	CountByCondition  func(arg service.MallSkuPageDTO) (int64, error)
}

func (it MallSkuMapper) New() MallSkuMapper {
	core_util.LoadMapper(&it)
	return it
}
