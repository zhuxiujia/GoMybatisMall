package c_app

import (
	"errors"
	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/captcha"
	"github.com/zhuxiujia/easy_mvc"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type CaptchaApi struct {
	easy_mvc.Controller `doc:"验证码API"`
	StoreService        *service.StoreService `inject:"StoreService"`

	GetImage func(phone string, writer http.ResponseWriter) interface{} `path:"/api/captcha" arg:"phone,w" doc:"图形验证码接口" doc_arg:"phone:手机号,w:_"`
	Slide    func(phone string, writer http.ResponseWriter) interface{} `path:"/api/captcha/slide" arg:"phone,w" doc:"图形验证码接口" doc_arg:"phone:手机号,w:_"`
}

//路由
func (it *CaptchaApi) Routers() {
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

	it.Slide = func(phone string, writer http.ResponseWriter) interface{} {
		if phone == "" {
			return errors.New("手机号不能空！")
		}
		if len(phone) != 11 {
			return errors.New("手机号不正确！")
		}
		i := 20 + rand.Intn(80)
		//save code
		var err = it.StoreService.SaveImageCode(service.ImageCodeDTO{
			Phone: phone,
			Code:  strconv.Itoa(i),
		})
		if err != nil {
			return err
		}
		log.Println("产生一个验证码[" + strconv.Itoa(i) + "],保存于   key[" + phone + "]")
		img := image.NewNRGBA(image.Rect(0, 0, 100, 10))
		for y := 0; y < 10; y++ {
			for x := 0; x < 100; x++ {
				if x > i {
					img.Set(x, y, color.RGBA{uint8(255), uint8(255), uint8(255), 255})
				} else {
					img.Set(x, y, color.RGBA{uint8(255), uint8(0), uint8(0), 255})
				}
			}
		}
		writer.Header().Set("Content-Type", "image/png")
		png.Encode(writer, img)
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
