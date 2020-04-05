package dao

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	_ "github.com/zhuxiujia/GoMybatisMall/core/com/example/dao/statik"
)

type AuthRoleMapper struct {
	SelectTemplete    func(id string) (model.AuthRole, error) `mapperParams:"id"`
	InsertTemplete    func(args model.AuthRole) error
	DeleteTemplete    func(id string) error `mapperParams:"id"`
	UpdateTemplete    func(arg model.AuthRole) error
	SelectByIds       func(ids []string) ([]model.AuthRole, error) `mapperParams:"ids"`
	SelectByCondition func(arg service.AuthRolePageDTO) ([]model.AuthRole, error)
	CountByCondition  func(arg service.AuthRolePageDTO) (int64, error)
}

func (it AuthRoleMapper) New() AuthRoleMapper {
	core_util.LoadMapper(&it)
	return it
}
