package c_app

import (
	"errors"
	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/easy_mvc"
	"strings"
)

type UserApi struct {
	easy_mvc.Controller `doc:"用户API"`

	UserService  *service.UserService   `inject:"UserService"`
	StoreService *service.StoreService  `inject:"StoreService"`
	AppConfig    *app_context.AppConfig `inject:"AppConfig"`

	Detail       func(login_id string) interface{}                                                                 `path:"/api/user/detail" arg:"login_id"  doc:"用户详情" doc_arg:"login_id:_"`
	Login        func(phone string, pwd string, img_code string) interface{}                                       `path:"/api/login" arg:"phone,pwd,img_code" doc:"登录" doc_arg:"img_code:图片验证码" `
	Register     func(phone string, vcode string, password string, invate_code string, channel string) interface{} `path:"/api/register" arg:"phone,vcode,password,invate_code,channel"  doc:"注册" doc_arg:"invate_code:邀请码(非必填),channel:注册渠道标识" `
	UploadAvatar func(login_id string, img_uri string) interface{}                                                 `path:"/api/user/upload/avatar" arg:"login_id:_,img_uri" doc:"修改用户头像接口 需要先调用/api/upload上传图片拿到图片url" doc_arg:"login_id:_,img_uri:已上传的图片url"  `
}

func (it *UserApi) Routers() {
	it.Detail = func(login_id string) interface{} {
		var user, e = it.UserService.Find(login_id)
		if e != nil {
			return e
		}
		user.ResetImportantInfo()

		address, e := it.UserService.FindUserAddress(user.DefAddressId)
		if e != nil {
			return e
		}
		if address.Id != "" {
			user.UserAddressVO = &address
		}
		return vo.ResultVO{}.NewSuccess(user)
	}
	//登录接口
	it.Login = func(phone string, pwd string, img_code string) interface{} {
		var imageReqDTO = service.ImageCodeDTO{
			Phone: phone,
			Code:  img_code,
		}
		if imageReqDTO.Phone == "" {
			return errors.New("手机号不能为空！")
		}
		if imageReqDTO.Code == "" {
			return errors.New("图片验证码不能为空！")
		}
		var e = it.StoreService.CheckImageCode(imageReqDTO)
		if e != nil {
			return e
		}
		var loginDTO = service.LoginDTO{
			Phone: phone,
			PWD:   pwd,
		}
		loginResult, e := it.UserService.Login(loginDTO)
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(loginResult)
	}
	//注册接口
	it.Register = func(phone string, vcode string, password string, invate_code string, channel string) interface{} {
		var arg = service.RegisterDTO{
			Vcode: vcode,
			UserVO: vo.UserVO{
				User: model.User{
					Phone:       phone,
					Password:    password,
					Channel:     channel,
					InviterCode: strings.ToLower(invate_code),
				},
			},
		}
		var e = it.UserService.Register(arg)
		if e != nil {
			return e
		}
		var loginDTO = service.LoginDTO{
			Phone: phone,
			PWD:   password,
		}
		loginResult, e := it.UserService.Login(loginDTO)
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(loginResult)
	}

	it.UploadAvatar = func(login_id string, img_uri string) interface{} {
		var e error
		u, e := it.UserService.Find(login_id)
		if e != nil {
			return e
		}
		u.Avatar = img_uri
		e = it.UserService.Update(u)
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(nil)
	}
	//finish
	it.Init(it)
	app_context.Context.Provide(&inject.Object{
		Value: it,
	})
}
