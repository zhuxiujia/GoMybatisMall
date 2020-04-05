package main

import (
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_init"
)

func main() {

	var appConfig = app_context.AppConfig{
		PlatformName:  "dev",
		SwaggerEnable: true,
		ClientName:    "App",
		Host:          "127.0.0.1",
		Port:          8000,
		ConsulUrl:     "127.0.0.1:8500",
		FastDFSUrl:    "http://127.0.0.1:8080",
	}

	app_init.Init(appConfig)
}
