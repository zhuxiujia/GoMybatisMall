package utils

import (
	"github.com/robfig/cron"
	"log"
	"reflect"
	"runtime"
)

//扫描结构体指针
func CronScan(it interface{}, c *cron.Cron) {
	var reflectV = reflect.ValueOf(it)
	if reflectV.Kind() != reflect.Ptr {
		panic("ScanStructPtr only support struct ptr!")
	}
	for {
		if reflectV.Kind() == reflect.Ptr {
			reflectV = reflectV.Elem()
		} else {
			break
		}
	}
	var structType = reflectV.Type()

	if structType.Kind() != reflect.Struct {
		panic("ScanStructPtr only support struct ptr!")
	}

	for i := 0; i < structType.NumField(); i++ {
		var item = structType.Field(i)
		var cron = item.Tag.Get("cron")
		if cron != "" && item.Type.Kind() == reflect.Func {
			var funcValue = reflectV.Field(i)
			var newFunc = func() {
				// 延迟处理的函数
				defer func() {
					// 发生宕机时，获取panic传递的上下文并打印
					err := recover()
					if err != nil {
						switch err.(type) {
						case runtime.Error: // 运行时错误
							log.Println("runtime error:", err)
						default: // 非运行时错误
							log.Println("error:", err)
						}
					}
				}()
				funcValue.Interface().(func())()
			}
			var e = c.AddFunc(cron, newFunc)
			if e != nil {
				panic(e)
			}
		}
	}
}
