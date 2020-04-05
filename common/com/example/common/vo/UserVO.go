package vo

import "github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"

type UserVO struct {
	model.User

	UserAddressVO *UserAddressVO `json:"user_address_vo"`

	PropertyVO *PropertyVO `json:"property_vo"`
}
