package core_util

import (
	"github.com/rakyll/statik/fs"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_context"
	_ "github.com/zhuxiujia/GoMybatisMall/core/com/example/dao/statik"
	"io/ioutil"
	"log"
	"reflect"
)

func LoadMapper(it interface{}) {
	var v = reflect.ValueOf(it)
	if v.Kind() != reflect.Ptr {
		panic("load " + v.String() + " must use ptr!")
	}
	for {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		} else {
			break
		}
	}

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	var name = "/" + v.Type().Name() + ".xml"
	log.Println("load xml:", name)
	var f, e = statikFS.Open(name)
	if e != nil {
		panic(e)
	}
	b, e := ioutil.ReadAll(f)
	core_context.Engine.WriteMapperPtr(it, b)
}
