package core_service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/zhuxiujia/GoMybatis"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/component/password"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/utils"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_context"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/dao"
	"log"
	"time"
)

type AdminUserServiceImpl struct {
	service.AdminUserService `bean:"AdminUserService"`
	adminUserMapper          dao.AdminUserMapper
	passwordEncoder          password.BCryptPasswordEncoder
	StoreService             *service.StoreService    `inject:"StoreService"`
	AuthRoleService          *service.AuthRoleService `inject:"AuthRoleService"`
}

func (it *AdminUserServiceImpl) Init() {
	it.adminUserMapper = it.adminUserMapper.New()
	it.DoLogin = func(arg service.AdminUserLoginDTO) (result vo.AdminUserVO, e error) {
		e = it.StoreService.CheckImageCode(service.ImageCodeDTO{
			Code:  arg.ImgCode,
			Phone: arg.Phone,
		})
		if e != nil {
			return result, e
		}
		adminUser, e := it.adminUserMapper.SelectByPhone(arg.Phone)
		if e != nil {
			return result, e
		}
		if adminUser.Id == "" {
			return result, errors.New("账号不存在！")
		}
		if adminUser.Enable == 0 {
			return result, errors.New("账号处于未启用状态，请联系管理员！")
		}
		if !it.passwordEncoder.Matches(arg.Pwd, adminUser.Pwd) {
			return result, errors.New("密码错误！")
		}

		//权限
		var roleIds []string
		json.Unmarshal([]byte(adminUser.RoleIds), &roleIds)
		roles, e := it.AuthRoleService.Finds(roleIds)
		if e != nil {
			return result, e
		}
		var urlMap = map[string]*string{}
		for _, v := range roles {
			if v.ResourceIds != "" {
				var resources []string
				json.Unmarshal([]byte(v.ResourceIds), &resources)
				if resources != nil {
					for _, item := range resources {
						urlMap[item] = &item
					}
				}
			}
		}
		var res = bytes.Buffer{}
		for v, _ := range urlMap {
			res.WriteString(utils.CountHash4(v) + ",")
		}
		var resStr = res.String()

		var access_token = utils.Sign(jwt.MapClaims{
			"id":          adminUser.Id,
			"phone":       adminUser.Phone,
			"res":         resStr,
			"create_time": time.Now().Format("2006-01-02 15:04:05"),
		}, utils.TokenEncoderString)
		log.Println(adminUser.Phone + " login access_token = " + access_token)
		result.AccessToken = access_token
		result.AdminUser = adminUser
		return result, nil
	}

	it.Update = func(arg model.AdminUser) error {
		if arg.Id != "" && arg.Pwd != "" && len(arg.Pwd) <= 20 {
			arg.Pwd = it.passwordEncoder.Encode(arg.Pwd)
		}
		return it.adminUserMapper.UpdateTemplete(arg)
	}

	it.DoRegister = func(arg service.AdminUserRegisterDTO) (result vo.AdminUserVO, e error) {
		if len(arg.Phone) != 11 {
			return result, errors.New("手机号不满足11位！")
		}
		if len(arg.Pwd) < 6 {
			return result, errors.New("密码不满足6位！")
		}
		e = it.StoreService.CheckImageCode(service.ImageCodeDTO{
			Code:  arg.ImgCode,
			Phone: arg.Phone,
		})
		if e != nil {
			return result, e
		}
		adminUser, e := it.adminUserMapper.SelectByPhone(arg.Phone)
		if e != nil {
			return result, e
		}
		if adminUser.Id != "" {
			return result, errors.New("账号已存在！")
		}
		adminUser = model.AdminUser{
			Id:         utils.CreateUUID(),
			Phone:      arg.Phone,
			Pwd:        it.passwordEncoder.Encode(arg.Pwd),
			Enable:     0,
			DeleteFlag: 1,
			CreateTime: time.Now(),
		}
		e = it.adminUserMapper.InsertTemplete(adminUser)
		if e != nil {
			return result, e
		}
		var access_token = utils.Sign(jwt.MapClaims{
			"id":    adminUser.Id,
			"phone": adminUser.Phone,
		}, utils.TokenEncoderString)
		log.Println(adminUser.Phone + " login access_token = " + access_token)
		result.AccessToken = access_token
		result.AdminUser = adminUser
		return result, nil
	}

	it.Page = func(arg service.AdminUserPageDTO) (page vo.PageVO, e error) {
		data, e := it.adminUserMapper.SelectPageTemplete(arg.Phone, arg.Enable, arg.Page, arg.PageSize)
		if e != nil {
			return page, e
		}
		total, e := it.adminUserMapper.SelectCountTemplete(arg.Phone, arg.Enable)
		if e != nil {
			return page, e
		}
		page = page.New(arg.Pageable, total, data)
		return page, nil
	}

	it.Enable = func(arg service.AdminUserEnableDTO) error {
		if arg.Enable == nil {
			return errors.New("启用禁用必须传递enable")
		}
		if arg.Id == "" {
			return errors.New("启用禁用id不能为空")
		}
		var user, e = it.adminUserMapper.SelectTemplete(arg.Id)
		if e != nil {
			return e
		}
		user.Enable = *arg.Enable
		return it.adminUserMapper.UpdateTemplete(user)
	}

	it.Find = func(id string) (userVO vo.AdminUserVO, e error) {
		data, e := it.adminUserMapper.SelectTemplete(id)
		if e != nil {
			return userVO, e
		}
		userVO.AdminUser = data
		return userVO, e
	}

	it.FindByPhone = func(phone string) (userVO vo.AdminUserVO, e error) {
		data, e := it.adminUserMapper.SelectByPhone(phone)
		if e != nil {
			return userVO, e
		}
		userVO.AdminUser = data
		return userVO, e
	}

	//end
	GoMybatis.AopProxyService(&it.AdminUserService, core_context.Engine)
	core_util.ScanInject("AdminUserServiceImpl", it)
}
