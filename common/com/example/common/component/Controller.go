package component

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/utils"
	"net/http"
)

type Controller struct {
}

func (controller Controller) HandleFunc(path string, f func(writer http.ResponseWriter, request *http.Request) interface{}) {
	http.HandleFunc(path, utils.JsonFunc(f))
}
