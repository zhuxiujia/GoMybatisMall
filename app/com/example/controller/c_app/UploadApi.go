package c_app

import (
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/GoMybatisMall/third/com/third/fs"

	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/easy_mvc"
)

type UploadApi struct {
	easy_mvc.Controller `doc:"文件上传接口"`

	AppConfig *app_context.AppConfig `inject:"AppConfig"`

	Upload func(m easy_mvc.MultipartFile) interface{} `method:"post" path:"/api/upload" arg:"file" doc_arg:""  `
}

//路由
func (it *UploadApi) Routers() {
	it.Upload = func(m easy_mvc.MultipartFile) interface{} {
		var result, e = fs.PostGoFastDfsFile(it.AppConfig.FastDFSUrl+"/upload", m.Filename, m.File)
		if e != nil {
			println(e.Error())
			return e
		}
		result = result + "?download=0"
		return vo.ResultVO{}.NewSuccess(result)
	}

	it.Init(it)
	app_context.Context.Provide(&inject.Object{
		Value: it,
	})
}
