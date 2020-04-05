package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/url"
	"time"
)

const TokenEncoderString = "123TEST"

const Key_login_phone = "login_phone"
const Key_login_id = "login_id"
const Key_res = "login_res"

//登录用户插入login_phone，login_id
func CheckLoginToken(req *url.Values, outOfTime *time.Duration) error {
	var access_token = req.Get("access_token")
	var data, e = VerifyToken(access_token, TokenEncoderString)
	if e != nil {
		return errors.New("登录过期，请重新登录!")
	}
	var phone = data.(jwt.MapClaims)["phone"]
	var id = data.(jwt.MapClaims)["id"]
	var res = data.(jwt.MapClaims)["res"]
	if phone == nil || id == nil {
		return errors.New("登录过期，请重新登录!")
	}
	if phone.(string) == "" || id.(string) == "" {
		return errors.New("登录过期，请重新登录!")
	}
	req.Set(Key_login_phone, phone.(string))
	req.Set(Key_login_id, id.(string))
	if res != nil {
		req.Set(Key_res, res.(string))
	}
	if outOfTime != nil {
		var create_time = data.(jwt.MapClaims)["create_time"]
		if create_time == nil {
			return nil
		}
		//2019-05-26 03:48:42.3965243 +0800 CST m=+3429.163087301
		var oldTime, e = time.Parse("2006-01-02 15:04:05", create_time.(string))
		if e != nil {
			return e
		}
		if time.Now().Sub(oldTime) > *outOfTime {
			//超时
			return errors.New("登录过期，请重新登录!")
		}
	}
	return nil
}
