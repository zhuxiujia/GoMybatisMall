package dao

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
)

type UserMapper struct {
	InsertTemplete         func(arg model.User) (int64, error)
	SelectById             func(id string) (model.User, error)         `mapperParams:"id"`
	SelectByIds            func(ids []string) ([]model.User, error)    `mapperParams:"ids"`
	SelectByPhone          func(phone string) (model.User, error)      `mapperParams:"phone"`
	SelectByPhones         func(phones []string) ([]model.User, error) `mapperParams:"phones"`
	UpdateTemplete         func(arg model.User) error
	SelectByInvitationCode func(inviter_code string) (model.User, error) `mapperParams:"inviter_code"`

	SelectPageTemplete  func(arg service.UserPageDTO) ([]model.User, error) `mapperParams:"arg"`
	SelectCountTemplete func(arg service.UserPageDTO) (int64, error)        `mapperParams:"arg"`
}

func (it UserMapper) New() UserMapper {
	core_util.LoadMapper(&it)
	return it
}
