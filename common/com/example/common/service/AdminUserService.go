package service

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
)

type AdminUserRegisterDTO = AdminUserLoginDTO

type AdminUserLoginDTO struct {
	Phone    string `json:"phone"`
	Pwd      string `json:"pwd"`
	ImgCode  string `json:"img_code"`
	RealName string `json:"real_name"`
}
type AdminUserPageDTO struct {
	vo.Pageable
	Phone    string `json:"phone"`
	RealName string `json:"real_name"`
	Enable   *int   `json:"enable"`
}
type AdminUserEnableDTO struct {
	Id     string `json:"id"`
	Enable *int   `json:"enable"`
}

type AdminUserService struct {
	DoLogin     func(arg AdminUserLoginDTO) (vo.AdminUserVO, error)
	DoRegister  func(arg AdminUserRegisterDTO) (vo.AdminUserVO, error)
	Page        func(arg AdminUserPageDTO) (vo.PageVO, error)
	Enable      func(arg AdminUserEnableDTO) error
	Update      func(arg model.AdminUser) error
	Find        func(id string) (vo.AdminUserVO, error)
	FindByPhone func(phone string) (vo.AdminUserVO, error)
}
