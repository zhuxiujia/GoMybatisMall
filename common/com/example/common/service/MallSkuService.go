package service

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
)

type MallSkuAddDTO struct {
	MallSku        model.MallSku //商城商品信息
	Images         *[]string     //商城商品图片信息
	Specifications *[]string     //商城商品规格信息
}

type MallSkuPageDTO struct {
	vo.Pageable
	Title     string `json:"title"`
	Status    *int   `json:"status"`
	ClassName string `json:"class_name"`
	MinAmount *int   `json:"min_amount"`
	MaxAmount *int   `json:"max_amount"`

	Sort     string `json:"sort"`
	Order_by string `json:"order_by"`

	Tag1 string `json:"tag1"` //热销
	Tag2 string `json:"tag2"` //新品
	Tag3 string `json:"tag3"` //精选
}

type MallSkuService struct {
	Add    func(arg MallSkuAddDTO) error
	Delete func(id string) error
	Update func(arg MallSkuAddDTO) (int64, error)
	Find   func(id string) (vo.MallSkuVO, error)
	Finds  func(ids []string) (map[string]*vo.MallSkuVO, error)
	Page   func(arg MallSkuPageDTO) (vo.PageVO, error)
}
