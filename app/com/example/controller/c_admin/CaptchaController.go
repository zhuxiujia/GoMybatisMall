package c_admin

import (
	"errors"
	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/captcha"
	"github.com/zhuxiujia/easy_mvc"
	"log"
	"net/http"
)

type CaptchaController struct {
	easy_mvc.Controller `doc:"验证码API"`
	StoreService        *service.StoreService `inject:"StoreService"`

	GetImage func(phone string, writer http.ResponseWriter) interface{} `path:"/admin/captcha" arg:"phone,w" doc:"图形验证码接口" doc_arg:"phone:手机号,w:_"`
}

//路由
func (it *CaptchaController) Routers() {
	it.GetImage = func(phone string, writer http.ResponseWriter) interface{} {
		if phone == "" {
			return errors.New("手机号不能空！")
		}
		if len(phone) != 11 {
			return errors.New("手机号不正确！")
		}
		//生成4位随机数字
		var digs = captcha.RandomDigits(4)
		var code = toDigitString(digs)
		//save code
		var err = it.StoreService.SaveImageCode(service.ImageCodeDTO{
			Phone: phone,
			Code:  code,
		})
		if err != nil {
			return err
		}
		var img = captcha.NewImage("", digs, 80, 42)
		log.Println("产生一个验证码[" + code + "],保存于   key[" + phone + "]")
		writer.Header().Set("Content-Type", "image/png")
		img.WriteTo(writer)
		return nil
	}

	it.Init(it)

	app_context.Context.Provide(&inject.Object{
		Value: it,
	})
}

func toDigitString(bytess []byte) string {
	var zero = byte('0')
	var resultBytes = make([]byte, 0)
	for _, item := range bytess {
		resultBytes = append(resultBytes, item+zero)
	}
	return string(resultBytes)
}
