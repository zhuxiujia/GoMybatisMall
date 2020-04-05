package service

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
)

type MallCoverImageService struct {
	Add           func(arg model.MallCoverImage) error
	Update        func(arg model.MallCoverImage) error
	Delete        func(id string) error
	DeleteBySkuId func(skuId string) error
	Find          func(id string) (vo.MallCoverImageVO, error)
	FindBySkuId   func(skuIds []string) (map[string]*[]vo.MallCoverImageVO, error)
}
