package dao

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	_ "github.com/zhuxiujia/GoMybatisMall/core/com/example/dao/statik"
)

type MallSpecificationMapper struct {
	InsertTemplete       func(arg model.MallSpecification) error
	UpdateTemplete       func(arg model.MallSpecification) error
	SelectTemplete       func(id string) (model.MallSpecification, error) `mapperParams:"id"`
	DeleteTemplete       func(id string) error                            `mapperParams:"id"`
	DeleteBySkuId        func(sku_id string) error                        `mapperParams:"sku_id"`
	DeleteBySkuIdAndName func(dto service.MallSpecificationDeleteDTO) error
	SelectByIds          func(ids []string) ([]model.MallSpecification, error)     `mapperParams:"ids"`
	SelectBySkuIds       func(sku_ids []string) ([]model.MallSpecification, error) `mapperParams:"sku_ids"`
}

func (it MallSpecificationMapper) New() MallSpecificationMapper {
	core_util.LoadMapper(&it)
	return it
}
