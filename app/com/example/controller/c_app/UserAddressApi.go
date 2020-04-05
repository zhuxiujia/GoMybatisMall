package c_app

import (
	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/easy_mvc"
)

type UserAddressApi struct {
	easy_mvc.Controller `doc:"用户地址API"`
	UserService         *service.UserService `inject:"UserService"`

	Page    func(login_id string, phone string, real_name string, page int, size int) interface{}    `path:"/api/user/address/page" arg:"login_id,phone,real_name,page:0,size:20" doc:"地址分页" doc_arg:"login_id:_"`
	Add     func(login_id string, real_name string, phone string, address_detail string) interface{} `path:"/api/user/address/add" arg:"login_id,real_name,phone,address_detail" doc:"添加地址" doc_arg:"login_id:_"`
	Update  func(id string, real_name string, phone string, address_detail string) interface{}       `path:"/api/user/address/update" arg:"id,real_name,phone,address_detail" doc:"更新地址" doc_arg:"id:用户地址id"`
	Delete  func(id string) interface{}                                                              `path:"/api/user/address/delete" arg:"id" doc:"删除地址" doc_arg:"id:用户地址id"`
	Default func(login_id string, addr_id string) interface{}                                        `path:"/api/user/address/default" arg:"login_id,addr_id" doc:"设置默认地址" doc_arg:"login_id:_,addr_id:用户地址id"`
}

func (it *UserAddressApi) Routers() {
	it.Page = func(login_id string, phone string, real_name string, page int, page_size int) interface{} {
		var arg = service.AddressPageDTO{}
		arg.Page = page
		arg.PageSize = page_size
		arg.RealName = real_name
		arg.Phone = phone
		arg.UserId = login_id
		var result, e = it.UserService.AddressPage(arg)
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(result)
	}

	it.Add = func(login_id string, real_name string, phone string, address_detail string) interface{} {
		e := it.UserService.AddAddress(vo.UserAddress{
			UserAddress: model.UserAddress{
				UserId:        login_id,
				RealName:      real_name,
				Phone:         phone,
				AddressDetail: address_detail,
			},
		})
		if e != nil {
			return e
		} else {
			return vo.ResultVO{}.NewSuccess(nil)
		}
	}

	it.Update = func(id string, real_name string, phone string, address_detail string) interface{} {

		var old, e = it.UserService.FindUserAddress(id)
		if e != nil {
			return e
		}
		old.RealName = real_name
		old.Phone = phone
		old.AddressDetail = address_detail
		e = it.UserService.UpdateAddress(vo.UserAddress{
			UserAddress: old.UserAddress,
		})
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(nil)
	}

	it.Delete = func(id string) interface{} {
		var e = it.UserService.DeleteAddress(id)
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(nil)
	}

	it.Default = func(user_id string, addr_id string) interface{} {
		var e = it.UserService.SetDefaultAddress(service.SetDefAddressDTO{
			user_id, addr_id,
		})
		if e != nil {
			return e
		} else {
			return vo.ResultVO{}.NewSuccess(nil)
		}
	}

	it.Init(it)

	app_context.Context.Provide(&inject.Object{
		Value: it,
	})
}
