package vo

import "github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"

type AdminUserVO struct {
	model.AdminUser
	AccessToken string `json:"access_token"`
}
