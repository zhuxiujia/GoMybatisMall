package core_service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/zhuxiujia/GoMybatis"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/component/password"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/utils"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_context"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/dao"
	"log"
	"time"
)

type UserServiceImpl struct {
	service.UserService `bean:"UserService"`

	userMapper        dao.UserMapper
	userAddressMapper dao.UserAddressMapper
	passwordEncoder   password.BCryptPasswordEncoder

	PropertyService *service.PropertyService `inject:"PropertyService"`
}

func (it *UserServiceImpl) Init() {
	it.userAddressMapper = it.userAddressMapper.New()
	it.userMapper = it.userMapper.New()

	//注册
	it.Register = func(arg service.RegisterDTO) error {
		if len(arg.Password) < 6 {
			return errors.New("密码长度必须大于6位！")
		}
		//重复用户检查
		oldUser, _ := it.userMapper.SelectByPhone(arg.Phone)
		if oldUser.Id != "" {
			return errors.New("该手机号已存在,请直接登录！")
		}
		//邀请码检查
		if arg.User.InviterCode != "" {

			if len(arg.User.InviterCode) < 11 {
				user, e := it.UserService.FindByInvitationCode(arg.User.InviterCode)
				if e != nil || user.Id == "" {
					return errors.New("邀请码不正确或者邀请人不存在！")
				}
			} else if len(arg.User.InviterCode) == 11 {
				user, e := it.UserService.FindByPhone(arg.User.InviterCode)
				if e != nil || user.Id == "" {
					return errors.New("手机号不正确或者邀请人不存在！")
				}
				arg.InviterCode = user.InvitationCode //设置邀请码
			} else {
				return errors.New("邀请码不正确！")
			}
		}

		arg.Id = utils.CreateUUID()
		arg.Password = it.passwordEncoder.Encode(arg.Password)
		arg.InvitationCode = string([]byte(arg.Id)[0:6]) //邀请码为 id 前6位
		arg.TodaySignIn = 1
		arg.ClientType = 0
		arg.CreateTime = time.Now()
		arg.DeleteFlag = 1
		_, e := it.userMapper.InsertTemplete(arg.User)
		if e != nil {
			return e
		}
		//insert property
		var property = vo.PropertyVO{}
		property.UserId = arg.Id
		e = it.PropertyService.Insert(property)
		if e != nil {
			return e
		}
		return e
	}

	//登录
	it.Login = func(arg service.LoginDTO) (result service.LoginResult, e error) {
		if arg.Phone == "" {
			return result, errors.New("登录的手机号不能为空！")
		}
		if arg.PWD == "" {
			return result, errors.New("登录的密码不能为空！")
		}
		if len(arg.PWD) < 6 {
			return result, errors.New("密码长度必须大于6位！")
		}
		u, e := it.userMapper.SelectByPhone(arg.Phone)
		if e != nil {
			return result, e
		}
		if u.Id == "" {
			return result, errors.New("用户" + arg.Phone + "不存在！")
		}
		if !it.passwordEncoder.Matches(arg.PWD, u.Password) {
			return result, errors.New("登录密码不正确！")
		}
		var access_token = utils.Sign(jwt.MapClaims{
			"id":    u.Id,
			"phone": u.Phone,
		}, utils.TokenEncoderString)
		log.Println(u.Phone + " login access_token = " + access_token)
		//event
		//返回结果
		result = service.LoginResult{
			AccessToken: access_token,
		}
		return result, nil
	}

	//用户地址页码
	it.AddressPage = func(arg service.AddressPageDTO) (result vo.PageVO, e error) {
		user, e := it.userMapper.SelectById(arg.UserId)
		if e != nil {
			return result, e
		}
		data, e := it.userAddressMapper.SelectPageTemplete(arg.UserId, arg.Phone, arg.RealName, arg.Page, arg.PageSize)
		if e != nil {
			return result, e
		}
		total, e := it.userAddressMapper.SelectCountTemplete(arg.UserId, arg.Phone, arg.RealName)
		if e != nil {
			return result, e
		}
		var newUserAddressVOs = []vo.UserAddressVO{}
		for _, v := range data {
			var item = vo.UserAddressVO{
				UserAddress: v,
			}
			if item.Id == user.DefAddressId {
				item.Def = 1
			}
			newUserAddressVOs = append(newUserAddressVOs, item)
		}
		result = result.New(arg.Pageable, total, newUserAddressVOs)
		return result, nil
	}

	//添加地址
	it.AddAddress = func(arg vo.UserAddress) error {
		if arg.UserId == "" {
			return errors.New("UserId不能空！")
		}
		if arg.Phone == "" {
			return errors.New("手机号不能空！")
		}
		if arg.RealName == "" {
			return errors.New("姓名不能空！")
		}
		if arg.AddressDetail == "" {
			return errors.New("详细地址不能空！")
		}

		arg.Id = utils.CreateUUID()
		arg.DeleteFlag = 1
		arg.CreateTime = time.Now()
		return it.userAddressMapper.InsertTemplete(arg.UserAddress)
	}

	//更新地址
	it.UpdateAddress = func(arg vo.UserAddress) error {
		var e = it.userAddressMapper.UpdateTemplete(arg.UserAddress)
		return e
	}

	//删除地址
	it.DeleteAddress = func(adddressId string) error {
		var e = it.userAddressMapper.DeleteTemplete(adddressId)
		return e
	}

	//设置默认地址
	it.SetDefaultAddress = func(arg service.SetDefAddressDTO) error {
		if arg.UserId == "" || arg.AdddressId == "" {
			return errors.New("user_id,adddress_id can not be null")
		}
		var u, e = it.userMapper.SelectById(arg.UserId)
		if e != nil {
			return e
		}
		addr, e := it.userAddressMapper.SelectTemplete(arg.AdddressId)
		if e != nil {
			return e
		}
		if addr.Id == "" {
			return errors.New("地址id不存在！")
		}
		u.DefAddressId = arg.AdddressId
		u.Address = addr.RealName + "," + addr.Phone + "," + addr.AddressDetail
		return it.userMapper.UpdateTemplete(u)
	}
	it.Find = func(login_id string) (userVO vo.UserVO, e error) {
		u, e := it.userMapper.SelectById(login_id)
		if e != nil {
			return userVO, e
		}
		userVO.User = u
		return userVO, nil
	}
	it.FindByPhone = func(phone string) (result vo.UserVO, e error) {
		if phone == "" {
			return result, errors.New("手机号为空！")
		}
		u, e := it.userMapper.SelectByPhone(phone)
		if e != nil {
			return result, e
		}
		result.User = u
		return result, nil
	}
	it.Update = func(arg vo.UserVO) error {
		if arg.Id == "" {
			return errors.New("更新用户失败,用户id不能为空")
		}
		return it.userMapper.UpdateTemplete(arg.User)
	}
	it.UpdatePwd = func(arg vo.UserVO) error {
		if len(arg.Password) < 6 {
			return errors.New("密码长度必须大于6位！")
		}
		arg.Password = it.passwordEncoder.Encode(arg.Password)
		it.Update(arg)
		return nil
	}

	it.FindByInvitationCode = func(inviter_code string) (result vo.UserVO, e error) {
		if inviter_code == "" {
			return result, nil
		}
		data, e := it.userMapper.SelectByInvitationCode(inviter_code)
		if e != nil {
			return result, e
		}
		result.User = data
		return result, nil
	}

	it.FindUserAddress = func(address_id string) (vo.UserAddressVO, error) {
		var data, e = it.userAddressMapper.SelectTemplete(address_id)
		if e != nil {
			return vo.UserAddressVO{}, e
		}
		var addressVO vo.UserAddressVO
		addressVO.UserAddress = data
		return addressVO, nil
	}

	it.FindByPhones = func(phones []string) (vos []vo.UserVO, e error) {
		if phones == nil {
			return vos, errors.New("手机号不能为空！")
		}
		for index, phone := range phones {
			if phone == "" || len(phone) != 11 {
				return vos, errors.New("手机号不正确或长度不满足11位！位置：第" + fmt.Sprint(index) + "条，手机号=" + phone)
			}
		}
		users, e := it.userMapper.SelectByPhones(phones)
		var userVOs = []vo.UserVO{}
		for _, item := range users {
			userVOs = append(userVOs, vo.UserVO{User: item})
		}
		return userVOs, nil
	}

	it.Page = func(arg service.UserPageDTO) (pageVO vo.PageVO, e error) {
		data, e := it.userMapper.SelectPageTemplete(arg)
		if e != nil {
			return pageVO, e
		}
		total, e := it.userMapper.SelectCountTemplete(arg)
		pageVO = pageVO.New(arg.Pageable, total, data)
		return pageVO, e
	}

	it.Finds = func(ids []string) (vos map[string]*vo.UserVO, e error) {
		if ids == nil || len(ids) == 0 {
			return vos, nil
		}
		vos = map[string]*vo.UserVO{}
		data, e := it.userMapper.SelectByIds(ids)
		if e != nil {
			return vos, e
		}
		for _, v := range data {
			if v.Id != "" {
				vos[v.Id] = &vo.UserVO{
					User: v,
				}
			}
		}
		return vos, e
	}

	GoMybatis.AopProxyService(&it.UserService, core_context.Engine)
	core_util.ScanInject("UserServiceImpl", it)
}
