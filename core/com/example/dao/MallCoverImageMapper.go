package dao

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	_ "github.com/zhuxiujia/GoMybatisMall/core/com/example/dao/statik"
)

type MallCoverImageMapper struct {
	InsertTemplete func(arg model.MallCoverImage) error
	UpdateTemplete func(arg model.MallCoverImage) error
	DeleteTemplete func(id string) error                                  `mapperParams:"id"`
	DeleteBySkuId  func(sku_id string) error                              `mapperParams:"sku_id"`
	SelectTemplete func(id string) (model.MallCoverImage, error)          `mapperParams:"id"`
	SelectBySkuIds func(sku_ids []string) ([]model.MallCoverImage, error) `mapperParams:"sku_ids"`
}

func (it MallCoverImageMapper) New() MallCoverImageMapper {
	core_util.LoadMapper(&it)
	return it
}
