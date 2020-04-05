package vo

import (
	"encoding/json"
)

type ResultVO struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (it *ResultVO) isNil() bool {
	if it.Code == 0 && it.Msg == "" {
		return true
	}
	return false
}

func (it ResultVO) NewSuccess(data interface{}) ResultVO {
	switch data.(type) {
	case error:
		var e = data.(error)
		if e != nil {
			it = it.NewError(-1, data.(error).Error())
			return it
		}
		break
	}
	it.Code = 1
	it.Msg = "成功"
	it.Data = data
	return it
}
func (it ResultVO) NewError(code int, msg string) ResultVO {
	it.Code = code
	it.Msg = msg
	return it
}

func (it *ResultVO) Error() string {
	var b, _ = json.Marshal(*it)
	return string(b)
}
