package c_app

import (
	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/easy_mvc"
)

//用户资产
type UserPropertyApi struct {
	easy_mvc.Controller `doc:"用户资产属性API"`
	PropertyService     *service.PropertyService          `inject:"PropertyService"`
	Property            func(login_id string) interface{} `path:"/api/user/property/detail" arg:"login_id" doc:"用户资产接口" doc_arg:"login_id:_"`
}

func (it *UserPropertyApi) Router() {
	it.Property = func(login_id string) interface{} {
		var property, e = it.PropertyService.FindByUser(login_id)
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(property)
	}
	it.Init(it)
	app_context.Context.Provide(&inject.Object{
		Value: it,
	})
}
