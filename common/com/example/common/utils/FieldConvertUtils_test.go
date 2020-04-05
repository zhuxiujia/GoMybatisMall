package utils

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"testing"
)

func TestConvertField(t *testing.T) {

	var skuMap = map[string]*vo.MallSkuVO{}
	skuMap["1"] = &vo.MallSkuVO{
		MallSku: model.MallSku{
			Id:    CreateUUID(),
			Title: "fdsa",
		},
	}

	var orderVOs = []vo.MallOrderVO{{
		MallOrder: model.MallOrder{
			Id:     CreateUUID(),
			SkuId:  "1",
			Remark: "asf",
		},
	}}

	//ConvertField(orderVOs,Convert{}.New("SkuId","MallSkuVO",skuMap))
	ConvertField(&orderVOs[0], Convert{}.New("SkuId", "MallSkuVO", skuMap))

	println(orderVOs[0].MallSkuVO.Title)
	println("done")

}
