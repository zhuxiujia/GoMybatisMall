package core_service

import (
	"errors"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/component/redis"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	"strconv"
	"strings"
)

type StoreServiceImpl struct {
	service.StoreService `bean:"StoreService"`

	RedisTemplete *redis.RedisTemplete `inject:"RedisTemplete"`
}

func (it *StoreServiceImpl) Init() {

	//it.SaveRes = func(arg service.SaveRoleDTO) error {
	//	var json, _ = json.Marshal(arg.Map)
	//	return it.RedisTemplete.Set("role:"+arg.Account, string(json), 4*60*60)
	//}
	//
	//it.GetRes = func(account string) (m map[string]*string, e error) {
	//	data, e := it.RedisTemplete.Get("role:" + account)
	//	if e != nil {
	//		return nil, errors.New("用户没有设置角色或者未登录！请重新登录！")
	//	}
	//	if data == "" {
	//		return nil, errors.New("用户没有设置角色或者未登录！请重新登录！")
	//	}
	//	e = json.Unmarshal([]byte(data), &m)
	//	if e != nil {
	//		return nil, e
	//	}
	//	return m, e
	//}

	it.SaveImageCode = func(arg service.ImageCodeDTO) error {
		println("存储图片验证码：code = ["+arg.Code+"],phone = ", arg.Phone)
		var e = it.RedisTemplete.Set("image_code:"+arg.Phone, arg.Code, 60)
		return e
	}
	it.CheckImageCode = func(arg service.ImageCodeDTO) error {
		var result, e = it.RedisTemplete.Get("image_code:" + arg.Phone)
		if e != nil {
			if strings.Contains(e.Error(), "nil") {
				return errors.New("验证码不匹配！")
			}
			return e
		}
		println("eq:", result, arg.Code)
		if result == arg.Code || isSameNumber(result, arg.Code) {
			it.RedisTemplete.Delete("image_code:" + arg.Phone)
			return nil
		} else {
			return errors.New("验证失败！")
		}
	}

	it.ListLPop = func(key string) (result string, e error) {
		r, e := it.RedisTemplete.ListLPop(key)
		if e != nil {
			return result, e
		}
		result = r
		return result, nil
	}
	it.ListRPop = func(key string) (result string, e error) {
		r, e := it.RedisTemplete.ListRPop(key)
		if e != nil {
			return result, e
		}
		result = r
		return result, nil
	}
	it.ListLPush = func(arg service.RedisKVDTO) error {
		e := it.RedisTemplete.ListLPush(arg.Key, arg.Value)
		if e != nil {
			return e
		}
		return nil
	}
	it.ListRPush = func(arg service.RedisKVDTO) error {
		e := it.RedisTemplete.ListRPush(arg.Key, arg.Value)
		if e != nil {
			return e
		}
		return nil
	}

	//GoMybatis.AopProxyService(&it.StoreService,core_context.Engine)
	core_util.ScanInject("StoreServiceImpl", it)
}

func isSameNumber(result string, code string) bool {
	arg1, e1 := strconv.Atoi(result)
	if e1 != nil {
		return false
	}
	arg2, e2 := strconv.Atoi(code)
	if e2 != nil {
		return false
	}
	if arg2 <= arg1+10 && arg2 >= arg1-10 {
		return true
	}
	return false
}
