package c_admin

import (
	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/utils"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/easy_mvc"
)

type MallSpecificationController struct {
	easy_mvc.Controller `doc:"(后台接口)商品规格"`

	MallSpecificationService *service.MallSpecificationService `inject:"MallSpecificationService"`

	Add    func(name string, sku_id string) interface{} `path:"/admin/user/mall/sku/specification/add" arg:"name,sku_id" doc:"商品规格添加" doc_arg:""`
	Delete func(sku_id string, name string) interface{} `path:"/admin/user/mall/sku/specification/delete" arg:"sku_id,name" doc:"商品规格删除" doc_arg:""`
}

func (it *MallSpecificationController) Routers() {

	it.Add = func(name string, sku_id string) interface{} {
		if name == "" || sku_id == "" {
			return vo.ResultVO{}.NewSuccess(nil)
		}
		var e = it.MallSpecificationService.Add(model.MallSpecification{
			Id:    utils.CreateUUID(),
			Name:  name,
			SkuId: sku_id,
		})
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(nil)
	}
	it.Delete = func(sku_id string, name string) interface{} {
		var e = it.MallSpecificationService.DeleteBySkuIdName(service.MallSpecificationDeleteDTO{
			SkuId: sku_id,
			Name:  name,
		})
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(nil)
	}

	it.Init(it)
	app_context.Context.Provide(&inject.Object{
		Value: it,
	})
}
