package app_init

import (
	"fmt"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/controller/c_admin"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/controller/c_app"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/controller/schedule"

	"github.com/facebookgo/inject"
	"github.com/robfig/cron"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/utils"
	"github.com/zhuxiujia/easy_mvc"
	"github.com/zhuxiujia/easy_mvc/easy_swagger"
	"github.com/zhuxiujia/easyrpc_discovery"
	"log"
	"net/http"
	"os"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprintln(w, "404 not fund")
		return
	}
	var data, _ = time.Now().MarshalText()
	fmt.Fprintln(w, "welcome ! app now is start up time:"+string(data))
}

func Init(appConfig app_context.AppConfig) {
	//init log
	var out = utils.NewOutPut("mnt/log/" + appConfig.ClientName + "/all.log")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(out)

	app_context.Context.Provide(&inject.Object{
		Name:  "AppConfig",
		Value: &appConfig,
	})

	var cronEngine = cron.New()

	easyrpc_discovery.EnableDiscoveryClient(
		nil,
		appConfig.ConsulUrl,
		appConfig.ClientName,
		appConfig.Host,
		appConfig.Port,
		5*time.Second,
		&easyrpc_discovery.RpcConfig{RetryTime: 1},
		app_context.RpcServiceBeanTable, true)

	http.HandleFunc("/info", IndexHandler)

	//错误处理链
	easy_mvc.RegisterGlobalErrorHandleChan(&app_context.ErrorFilter)

	//返回结果链
	easy_mvc.RegisterGlobalResultHandleChan(&app_context.ResultFilter)

	//日志链
	easy_mvc.RegisterGlobalHttpChan(&app_context.LogChain)

	//过滤链 app登录检查
	easy_mvc.RegisterGlobalHttpChan(&app_context.LoginFilter)

	//后台登录检查
	easy_mvc.RegisterGlobalHttpChan(&app_context.LoginFilterAdmin)

	var UserController = c_admin.UserController{}
	UserController.Routers()

	var CaptchaController = c_admin.CaptchaController{}
	CaptchaController.Routers()

	var MallClassController = c_admin.MallClassController{}
	MallClassController.Routers()

	var MallSkuController = c_admin.MallSkuController{}
	MallSkuController.Routers()

	var MallOrderController = c_admin.MallOrderController{}
	MallOrderController.Routers()

	var MallSpecificationController = c_admin.MallSpecificationController{}
	MallSpecificationController.Routers()

	var AuthResourceController = c_admin.AuthResourceController{}
	AuthResourceController.Routers()

	var AuthRoleController = c_admin.AuthRoleController{}
	AuthRoleController.Routers()
	//app

	var app_user = c_app.UserApi{}
	app_user.Routers()

	var app_cap = c_app.CaptchaApi{}
	app_cap.Routers()

	var UserAddressController = c_app.UserAddressApi{}
	UserAddressController.Routers()

	var userPropertyController = c_app.UserPropertyApi{}
	userPropertyController.Router()

	var UploadApi = c_app.UploadApi{}
	UploadApi.Routers()

	var mallOrderApi = c_app.MallOrderApi{}
	mallOrderApi.Routers()

	var mallSkuApi = c_app.MallSkuApi{}
	mallSkuApi.Routers()

	var zfbNotify = schedule.OrderNotifyZfbSchedule{}
	zfbNotify.Init(cronEngine)

	//基本地址，用于查看服务是否启动
	http.HandleFunc("/", IndexHandler)

	if appConfig.SwaggerEnable {
		easy_swagger.EnableSwagger(appConfig.Host+":"+fmt.Sprint(appConfig.Port), easy_swagger.SwaggerConfig{
			SecurityDefinitionConfig: &easy_swagger.SecurityDefinitionConfig{
				SecurityDefinition: easy_swagger.SecurityDefinition{
					ApiKey: easy_swagger.ApiKey{
						Type: "apiKey",
						Name: "access_token",
						In:   "query",
					},
				},
				Path: "/api/user",
			},
		})
	}
	//执行IOC控制反转注入
	if err := app_context.Context.Populate(); err != nil {
		log.Println("------------------App IOC注入失败,请检查app_context中配置是否加入------------------")
		log.Fatal(os.Stderr, err)
		os.Exit(1)
	}

	//扫描权限
	AuthResourceController.Scan()

	//启动cron表达式调度引擎
	cronEngine.Start()

	log.Println("------------------" + appConfig.ClientName + " 初始化完成------------------")
	var httpHost = ""
	if appConfig.Host != "127.0.0.1" && appConfig.Host != "0.0.0.0" && appConfig.Host != "localhost" {
		httpHost = appConfig.Host
	}
	//调用golang官方http 启动http服务
	http.ListenAndServe(fmt.Sprint(httpHost, ":", appConfig.Port), nil)
}
