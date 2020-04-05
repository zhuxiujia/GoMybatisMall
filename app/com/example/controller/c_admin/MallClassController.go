package c_admin

import (
	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/easy_mvc"
)

type MallClassController struct {
	easy_mvc.Controller `doc:"(后台接口)商品分类"`

	MallClassService *service.MallClassService `inject:"MallClassService"`

	Page   func(name string, page int, size int) interface{}         `path:"/admin/user/mall/class/page" arg:"name,page:0,size:5" doc:"获取积分商城商品 分类列表信息" doc_arg:""`
	All    func() interface{}                                        `path:"/admin/user/mall/class/all" arg:"" doc:"获取积分商城商品 分类列表信息" doc_arg:""`
	Add    func(name string, logo_img string) interface{}            `path:"/admin/user/mall/class/add" arg:"name,logo_img" doc:"添加积分商城产品_分类" doc_arg:""`
	Detail func(id string) interface{}                               `path:"/admin/user/mall/class/detail" arg:"name" doc:"详情" doc_arg:""`
	Update func(id string, name string, logo_img string) interface{} `path:"/admin/user/mall/class/update" arg:"id,name,logo_img" doc:"更新" doc_arg:""`
	Delete func(id string) interface{}                               `path:"/admin/user/mall/class/delete" arg:"id" doc:"删除" doc_arg:""`
}

func (it *MallClassController) Routers() {

	it.Page = func(name string, page int, size int) interface{} {
		var data, e = it.MallClassService.Page(service.MallClassPageDTO{
			Name:     name,
			Pageable: vo.Pageable{}.New(page, size),
		})
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(data)
	}

	it.Add = func(name string, logo_img string) interface{} {
		e := it.MallClassService.Add(model.MallClass{
			Name:    name,
			LogoImg: logo_img,
		})
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(nil)
	}

	it.Detail = func(id string) interface{} {
		data, e := it.MallClassService.Find(id)
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(data)
	}

	it.Update = func(id string, name string, logo_img string) interface{} {
		data, e := it.MallClassService.Find(id)
		if e != nil {
			return e
		}
		data.Name = name
		data.LogoImg = logo_img
		e = it.MallClassService.Update(data.MallClass)
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(nil)
	}
	it.Delete = func(id string) interface{} {
		e := it.MallClassService.Delete(id)
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(nil)
	}
	it.All = func() interface{} {
		var data, e = it.MallClassService.FindAll()
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(data)
	}

	it.Init(it)
	app_context.Context.Provide(&inject.Object{
		Value: it,
	})
}
