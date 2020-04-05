package vo

import "github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"

type MallOrderVO struct {
	model.MallOrder

	UserVO              *UserVO              `json:"user_vo"`
	MallSkuVO           *MallSkuVO           `json:"mall_sku_vo"`
	MallSpecificationVO *MallSpecificationVO `json:"mall_specification_vo"`
}
