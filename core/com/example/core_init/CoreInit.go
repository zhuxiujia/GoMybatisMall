package core_init

import (
	"fmt"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/component/redis"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/utils"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_context"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_service"

	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/GoMybatis"
	"github.com/zhuxiujia/easyrpc_discovery"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func Init(config core_context.CoreConfig) {
	//init log
	var out = utils.NewOutPut("mnt/log/" + config.Server + "/all.log")
	log.SetOutput(out)

	//日志链
	core_context.Context.Provide(&inject.Object{
		Name:  "log",
		Value: &out,
	})

	//核心配置文件
	core_context.Context.Provide(&inject.Object{
		Name:  "CoreConfig",
		Value: &config,
	})

	//RedisTemplete
	var RedisTemplete = redis.RedisTemplete{}.New(config.Redis_url, config.Redis_password)
	core_context.Context.Provide(&inject.Object{
		Name:  "RedisTemplete",
		Value: &RedisTemplete,
	})

	var GoMybatisEngine = GoMybatis.GoMybatisEngine{}.New()
	//mysql链接格式为         用户名:密码@(数据库链接地址:端口)/数据库名称   例如root:123456@(***.mysql.rds.aliyuncs.com:3306)/test
	_, err := GoMybatisEngine.Open("mysql", config.Mysql) //此处请按格式填写你的mysql链接，这里用*号代替
	if err != nil {
		panic(err.Error())
	}
	core_context.Context.Provide(&inject.Object{
		Name:  "GoMybatisEngine",
		Value: &GoMybatisEngine,
	})
	core_context.Engine = &GoMybatisEngine

	var StoreServiceImpl = core_service.StoreServiceImpl{}
	StoreServiceImpl.Init()

	var UserServiceImpl = core_service.UserServiceImpl{}
	UserServiceImpl.Init()

	var PropertyServiceImpl = core_service.PropertyServiceImpl{}
	PropertyServiceImpl.Init()

	var KVServiceImpl = core_service.KVServiceImpl{}
	KVServiceImpl.Init()

	var AdminUserServiceImpl = core_service.AdminUserServiceImpl{}
	AdminUserServiceImpl.Init()

	var RegionServiceImpl = core_service.RegionServiceImpl{}
	RegionServiceImpl.Init()

	var PayCallBackServiceImpl = core_service.PayCallBackServiceImpl{}
	PayCallBackServiceImpl.Init()

	// mall
	var MallBrandServiceImpl = core_service.MallClassServiceImpl{}
	MallBrandServiceImpl.Init()

	var MallCoverImageServiceImpl = core_service.MallCoverImageServiceImpl{}
	MallCoverImageServiceImpl.Init()

	var MallOrderServiceImpl = core_service.MallOrderServiceImpl{}
	MallOrderServiceImpl.Init()

	var MallSkuServiceImpl = core_service.MallSkuServiceImpl{}
	MallSkuServiceImpl.Init()

	var MallSpecificationServiceImpl = core_service.MallSpecificationServiceImpl{}
	MallSpecificationServiceImpl.Init()

	var AuthRoleServiceImpl = core_service.AuthRoleServiceImpl{}
	AuthRoleServiceImpl.Init()

	var serviceTable = map[string]interface{}{}
	for _, v := range core_context.Context.Objects() {
		if strings.Contains(v.Name, "Service") && !strings.Contains(v.Name, "Impl") {
			serviceTable[v.Name] = v.Value
		}
	}

	//执行注入
	if err := core_context.Context.Populate(); err != nil {
		log.Println("------------------App 注解注入失败------------------")
		log.Fatal(os.Stderr, err)
		os.Exit(1)
	}
	log.Println("------------------" + config.Server + " 初始化成功------------------")

	core_context.Context.Objects()
	//远程服务信息
	easyrpc_discovery.EnableDiscoveryService(
		config.Consul,
		serviceTable,
		config.Address,
		config.Port, 30*time.Second, func(recover interface{}) string {
			if recover != nil {

			}
			var stackData = debug.Stack()
			var stack = string(stackData)
			os.Stderr.Write(stackData)
			log.Println(stack)
			var errorInfo = fmt.Sprint(recover) + stack
			return errorInfo
		})
}
