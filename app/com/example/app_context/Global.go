package app_context

import (
	"encoding/json"
	"fmt"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/utils"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"

	"github.com/zhuxiujia/easy_mvc"
	"log"
	"net/http"
	"reflect"
	"runtime/debug"
	"strings"
	"time"
)

var ErrorFilter = easy_mvc.HttpErrorHandle{
	Func: func(err interface{}, w http.ResponseWriter, r *http.Request) {
		if err != nil {
			debug.PrintStack()
			switch err.(type) {
			case error:
				println(err.(error).Error())
				var vo = vo.ResultVO{}.NewError(-1, err.(error).Error()+string(debug.Stack()))
				var b, _ = json.Marshal(vo)
				w.Write(b)
			default:
				var errdata = fmt.Sprint(err) + string(debug.Stack())
				println(errdata)
				var vo = vo.ResultVO{}.NewError(-1, errdata)
				var b, _ = json.Marshal(vo)
				w.Write(b)
			}
		}
	},
	Name: "ErrorFilter",
}

var ResultFilter = easy_mvc.HttpResultHandle{
	Func: func(result *interface{}, w http.ResponseWriter, r *http.Request) bool {
		if *result != nil {
			var rv = reflect.TypeOf(*result)
			if strings.Contains(rv.String(), "error") || strings.Contains(rv.String(), "Error") {
				var msg = (*result).(error).Error()
				var v = vo.ResultVO{}.NewError(-1, msg)
				var b, _ = json.Marshal(v)
				w.Write(b)
				return true
			} else {
				var b, _ = json.Marshal(r)
				w.Write(b)
			}
		}
		return false
	},
	Name: "ResultFilter",
}

type CustomHttpChan struct {
	easy_mvc.HttpChan
}

//日志链
var LogChain = easy_mvc.HttpChan{
	Name: "LogChain",
	Func: func(path string, w http.ResponseWriter, r *http.Request) bool {
		var reqInfo = "http " + r.Method + "  => " + path + " "
		reqInfo += ":arg=" + fmt.Sprint(r.Form)
		//reqInfo+=":header="+fmt.Sprint(r.Header)
		log.Println(reqInfo)
		return false
	},
}

//登录过滤
var LoginFilter = easy_mvc.HttpChan{
	Name: "LoginFilter:" + "/api/user/",
	Func: func(path string, w http.ResponseWriter, r *http.Request) bool {
		if strings.Contains(path, "/api/user/") {
			//拦截"/api/user"的接口
			var e = utils.CheckLoginToken(&r.Form, nil)
			if e != nil {
				var r = vo.ResultVO{}.NewError(-6, e.Error())
				var b, _ = json.Marshal(r)
				w.Write(b)
				return true
			}
		}
		return false
	},
}

//后台登录
var LoginFilterAdmin = easy_mvc.HttpChan{
	Name: "LoginFilter:" + "/admin/",
	Func: func(path string, w http.ResponseWriter, r *http.Request) bool {
		if strings.Contains(path, "/admin/user/") {
			//过期检查
			var outOfTime = 4 * time.Hour
			var e = utils.CheckLoginToken(&r.Form, &outOfTime)
			if e != nil {
				var r = vo.ResultVO{}.NewError(-6, e.Error())
				var b, _ = json.Marshal(r)
				w.Write(b)
				return true
			}
			//权限检查
			var res = r.Form.Get(utils.Key_res)
			var hash = utils.CountHash4(r.URL.Path)
			if strings.Contains(res, hash) {
				println("done")
			} else {
				var r = vo.ResultVO{}.NewError(-6, "当前资源无权限访问！")
				var b, _ = json.Marshal(r)
				w.Write(b)
				return true
			}

		}
		return false
	},
}
