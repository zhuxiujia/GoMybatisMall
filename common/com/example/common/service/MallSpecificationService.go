package service

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
)

type MallSpecificationDeleteDTO struct {
	SkuId string `json:"sku_id"`
	Name  string `json:"name"`
}

type MallSpecificationService struct {
	Add               func(arg model.MallSpecification) error
	Update            func(arg model.MallSpecification) error
	Delete            func(id string) error
	DeleteBySkuIdName func(arg MallSpecificationDeleteDTO) error
	Find              func(id string) (vo.MallSpecificationVO, error)
	//根据积分商品id删除所有商品规格
	DeleteBySkuId func(skuId string) error
	FindBySkuIds  func(skuIds []string) (map[string]*[]vo.MallSpecificationVO, error)
	Finds         func(ids []string) (map[string]*vo.MallSpecificationVO, error)
}
