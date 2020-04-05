package utils

import (
	"errors"
	"reflect"
)

//arg array or struct
func GetIds(arg interface{}, field string) []string {
	if arg == nil {
		return nil
	}
	var t = reflect.TypeOf(arg)
	var v = reflect.ValueOf(arg)
	var e = loopGetValue(&v)
	if e != nil {
		return nil
	}

	e = loopGetType(&t)
	if e != nil {
		return nil
	}

	if t.Kind() == reflect.String {
		var vValue = v.Interface().(string)
		if vValue != "" {
			return []string{vValue}
		}
		return nil
	} else if t.Kind() == reflect.Struct {
		var str = ""
		var success = getItemResult(v, field, &str)
		if success {
			return []string{str}
		}
		return nil
	}
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		panic("GetIds()  arg is not a array or slice or struct or string!")
	}
	var result = make([]string, 0)
	for i := 0; i < v.Len(); i++ {
		var item = v.Index(i)
		var str = ""
		var success = getItemResult(item, field, &str)
		if success == false {
			continue
		}
		result = append(result, str)
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func loopGetValue(v *reflect.Value) error {
	if v == nil {
		return errors.New("loopGetValue v is nil!")
	}
	for {
		if v.Kind() == reflect.Ptr {
			if v.Kind() == reflect.Ptr && v.IsNil() {
				return errors.New("loopGetValue v is nil!")
			}
			*v = v.Elem()
		} else {
			break
		}
	}
	return nil
}

func loopGetType(v *reflect.Type) error {
	if v == nil {
		return errors.New("loopGetValue v is nil!")
	}
	for {
		if (*v).Kind() == reflect.Ptr {
			if (*v).Kind() == reflect.Ptr {
				return errors.New("loopGetValue v is nil!")
			}
			*v = (*v).Elem()
		} else {
			break
		}
	}
	return nil
}

//return success
func getItemResult(item reflect.Value, field string, result *string) bool {
	if item.Kind() == reflect.Ptr && item.IsNil() {
		return false
	}
	var fieldValue = item.FieldByName(field)
	var e = loopGetValue(&fieldValue)
	if e != nil {
		return false
	}
	var itemResult = fieldValue.Interface().(string)
	if itemResult == "" {
		return false
	}
	*result = itemResult
	return true
}
