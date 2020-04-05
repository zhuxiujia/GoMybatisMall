package service

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
)

type AuthRolePageDTO struct {
	vo.Pageable
	Name string `json:"name"`
}

type AuthRoleService struct {
	Find   func(id string) (vo.AuthRoleVO, error)
	Add    func(arg model.AuthRole) error
	Update func(arg model.AuthRole) error
	Delete func(id string) error
	Finds  func(ids []string) ([]vo.AuthRoleVO, error)

	Page func(arg AuthRolePageDTO) (vo.PageVO, error)
}
