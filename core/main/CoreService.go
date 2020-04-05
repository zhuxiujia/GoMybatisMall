package main

import (
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_context"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_init"
)

func main() {
	var CoreConfig = core_context.CoreConfig{
		Server:         "CoreService",
		Address:        "127.0.0.1",
		Consul:         "127.0.0.1:8500",
		Port:           1234,
		Mysql:          "root:123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local",
		Redis_url:      "127.0.0.1:6379",
		Redis_password: "",
		CashierUrl:     "", //TODO 注意这里是 支付宝预下单接口http地址，需要自行实现（Java端提供服务接口）.可以查阅支付宝官网文档
	}
	core_init.Init(CoreConfig)
}
