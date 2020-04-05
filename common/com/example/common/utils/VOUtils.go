package utils

import (
	"reflect"
	"strings"
)

func ConvertToVOs(vos interface{}, fields interface{}) {
	if fields == nil {
		return
	}
	if vos == nil {
		panic("[ExtendBeanBuildUtil] childArray can not be null!")
	}
	var childArrayValue = reflect.ValueOf(vos)
	if childArrayValue.Kind() != reflect.Ptr {
		panic("[ExtendBeanBuildUtil]" + childArrayValue.String() + " is not a ptr")
	}
	if childArrayValue.Elem().Kind() != reflect.Slice {
		panic("[ExtendBeanBuildUtil]" + childArrayValue.String() + " is not a Slice")
	}

	var fatherArrayValue = reflect.ValueOf(fields)
	if fatherArrayValue.Kind() != reflect.Slice {
		panic("[ExtendBeanBuildUtil]" + fatherArrayValue.String() + " is not a Slice")
	}
	var slice = reflect.MakeSlice(childArrayValue.Type().Elem(), fatherArrayValue.Len(), fatherArrayValue.Cap())
	for i := 0; i < fatherArrayValue.Len(); i++ {
		var itemInfValue = fatherArrayValue.Index(i)
		var vosValue = slice.Index(i)
		var names = strings.Split(itemInfValue.Type().String(), ".")
		var name = names[len(names)-1]
		var f = vosValue.FieldByName(name)
		f.Set(itemInfValue)
	}
	childArrayValue.Elem().Set(slice)
}
