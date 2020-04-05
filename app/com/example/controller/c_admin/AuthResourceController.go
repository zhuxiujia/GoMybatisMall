package c_admin

import (
	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/easy_mvc"
	"reflect"
	"strings"
)

type AuthResourceController struct {
	easy_mvc.Controller

	Page func() interface{} `path:"/admin/res/page" arg:"" doc:"资源列表" doc_arg:"" `
}

//路由
func (it *AuthResourceController) Routers() {

	it.Page = func() interface{} {
		return vo.ResultVO{}.NewSuccess(app_context.UrlArray)
	}

	it.Init(it)
	app_context.Context.Provide(&inject.Object{
		Value: it,
	})
}

func (it *AuthResourceController) Scan() {
	var objs = app_context.Context.Objects()
	for _, item := range objs {
		var v = reflect.ValueOf(item.Value)
		var vType = v.Type().Elem()
		if strings.Contains(vType.PkgPath(), "/c_admin") && strings.Contains(vType.String(), "Controller") {
			//admin path
			for i := 0; i < vType.NumField(); i++ {
				var field = vType.Field(i)
				var path = field.Tag.Get("path")
				if path != "" {
					var doc = field.Tag.Get("doc")
					if doc == "" {
						panic(path + "必须添加doc文档描述！")
					}
					println("AuthResourceController.Scan() url:" + path + ",doc:" + doc)
					app_context.UrlArray = append(app_context.UrlArray, vo.AuthResourceVO{
						Name: doc,
						Url:  path,
					})
				}
			}
		}
	}
}
