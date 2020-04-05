package c_admin

import (
	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/utils"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/easy_mvc"
	"time"
)

type AuthRoleController struct {
	easy_mvc.Controller
	AuthRoleService *service.AuthRoleService `inject:"AuthRoleService"`

	Page   func(name string, page int, size int) interface{}             `path:"/admin/user/auth/page" arg:"name,page:0,size:5" doc:"管理员角色列表" doc_arg:"" `
	Add    func(name string, resource_ids string) interface{}            `path:"/admin/user/auth/add" arg:"name,resource_ids" doc:"管理员角色添加" doc_arg:"" `
	Update func(id string, name string, resource_ids string) interface{} `path:"/admin/user/auth/update" arg:"id,name,resource_ids" doc:"管理员角色修改" doc_arg:"" `
	Delete func(id string) interface{}                                   `path:"/admin/user/auth/delete" arg:"id" doc:"管理员角色删除" doc_arg:"" `
}

//路由
func (it *AuthRoleController) Routers() {

	it.Page = func(name string, page int, size int) interface{} {
		var data, e = it.AuthRoleService.Page(service.AuthRolePageDTO{
			Pageable: vo.Pageable{}.New(page, size),
			Name:     name,
		})
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(data)
	}

	it.Add = func(name string, resource_ids string) interface{} {
		var e = it.AuthRoleService.Add(model.AuthRole{
			Id:          utils.CreateUUID(),
			Name:        name,
			ResourceIds: resource_ids,
			Version:     0,
			CreateTime:  time.Now(),
			DeleteFlag:  1,
		})
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(nil)
	}

	it.Update = func(id string, name string, resource_ids string) interface{} {
		data, e := it.AuthRoleService.Find(id)
		if e != nil {
			return e
		}
		data.Name = name
		data.ResourceIds = resource_ids
		e = it.AuthRoleService.Update(data.AuthRole)
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(nil)
	}

	it.Delete = func(id string) interface{} {
		e := it.AuthRoleService.Delete(id)
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
