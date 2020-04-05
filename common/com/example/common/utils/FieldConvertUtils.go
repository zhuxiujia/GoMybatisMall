package utils

import (
	"reflect"
)

type Convert struct {
	mapKey       string //map key，即数据源的对比Id
	setFieldName string //需要设置数据的fieldName
	_map         map[string]*reflect.Value
}

func (it Convert) New(mapKey string, setFieldName string, Map interface{}) Convert {
	it.mapKey = mapKey
	it.setFieldName = setFieldName
	it._map = nil
	if Map != nil {
		var m = reflect.ValueOf(Map)
		loopGetValue(&m)
		var keys = m.MapKeys()
		if len(keys) == 0 {
			return it
		}
		it._map = make(map[string]*reflect.Value)
		for i := 0; i < len(keys); i++ {
			var keyv = keys[i]
			if keyv.Kind() != reflect.String {
				panic("Convert only support map[string]*")
			}
			var mapv = m.MapIndex(keyv)
			it._map[keyv.Interface().(string)] = &mapv
		}
	}
	return it
}

func ConvertField(arg interface{}, args ...Convert) error {
	if arg == nil {
		return nil
	}
	if args == nil || len(args) == 0 {
		return nil
	}
	var v = reflect.ValueOf(arg)
	if v.Kind() != reflect.Ptr && v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		panic("ConvertField arg only support struct ptr, array slice!")
	}
	var e = loopGetValue(&v)
	if e != nil {
		return nil
	}
	if v.Kind() == reflect.Struct {
		for _, convert := range args {
			getStructItemResult(v, convert)
		}
		return nil
	}
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		panic("GetArrayIds arg is not a array or slice or struct or string!")
	}
	for i := 0; i < v.Len(); i++ {
		var item = v.Index(i)
		for _, convert := range args {
			getStructItemResult(item, convert)
		}
	}
	return nil
}

//return success
func getStructItemResult(item reflect.Value, convert Convert) bool {
	if convert._map == nil {
		return false
	}
	if item.Kind() == reflect.Ptr && item.IsNil() {
		return false
	}
	var fv = item.FieldByName(convert.setFieldName)
	var kv = item.FieldByName(convert.mapKey)
	var kvStr = kv.Interface().(string)
	var mapResult = convert._map[kvStr]
	if mapResult == nil {
		return false
	}
	fv.Set(*mapResult)
	return true
}
