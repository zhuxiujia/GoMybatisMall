package vo

import "github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"

type UserAddressVO struct {
	model.UserAddress
	Def int `json:"def"` //是否默认地址
}
