package app_context

import (
	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/easyrpc_discovery"
	"reflect"
)

var RpcServiceBeanTable = []easyrpc_discovery.RpcServiceBean{}

//只需要在这里注册服务
var StoreService service.StoreService
var UserService service.UserService
var PropertyService service.PropertyService
var KVService service.KVService
var AdminUserService service.AdminUserService
var RegionService service.RegionService
var MallClassService service.MallClassService
var MallCoverImageService service.MallCoverImageService
var MallOrderService service.MallOrderService
var MallSkuService service.MallSkuService
var MallSpecificationService service.MallSpecificationService
var AuthRoleService service.AuthRoleService
var PayService service.PayService

func init() {
	ReflectRegisterService(&StoreService)
	ReflectRegisterService(&UserService)
	ReflectRegisterService(&PropertyService)
	ReflectRegisterService(&KVService)
	ReflectRegisterService(&AdminUserService)
	ReflectRegisterService(&RegionService)

	ReflectRegisterService(&MallClassService)
	ReflectRegisterService(&MallCoverImageService)
	ReflectRegisterService(&MallOrderService)
	ReflectRegisterService(&MallSkuService)
	ReflectRegisterService(&MallSpecificationService)
	ReflectRegisterService(&AuthRoleService)
	ReflectRegisterService(&PayService)
}

func ReflectRegisterService(server interface{}) {
	var t = reflect.TypeOf(server)
	if t.Kind() != reflect.Ptr {
		panic("registerServiceByReflectName must be a ptr")
	}
	for {
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		} else {
			break
		}
	}
	registerService(t.Name(), server)
}

//注册到服务表，配置表
func registerService(remoteService string, server interface{}) {
	var serverType = reflect.TypeOf(server)
	if serverType.Kind() == reflect.Ptr {
		serverType = serverType.Elem()
	}
	var serviceName = serverType.Name()
	Context.Provide(&inject.Object{
		Name:  serviceName,
		Value: server,
	})
	RpcServiceBeanTable = append(RpcServiceBeanTable, easyrpc_discovery.RpcServiceBean{
		Service:           server,
		ServiceName:       serviceName,
		RemoteServiceName: remoteService,
	})
}
