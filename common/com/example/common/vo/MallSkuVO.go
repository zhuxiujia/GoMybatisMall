package vo

import "github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"

type MallSkuVO struct {
	model.MallSku
	MallCoverImageVOs    *[]MallCoverImageVO    `json:"mall_cover_image_vos"`
	MallSpecificationVOs *[]MallSpecificationVO `json:"mall_specification_vos"`
}
