package dao

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
)

type KVMapper struct {
	InsertTemplete func(arg model.KV) error
	UpdateTemplete func(arg model.KV) error
	DeleteTemplete func(id string) error                         `mapperParams:"id"`
	SelectTemplete func(id string) (model.KV, error)             `mapperParams:"id"`
	SelectIdLike   func(id_like_name string) ([]model.KV, error) `mapperParams:"id"`

	SelectByIds func(ids []string) ([]model.KV, error) `mapperParams:"ids"`

	SelectPageTemplete  func(id string, remark string, page int, size int) ([]model.KV, error) `mapperParams:"id,remark,page,size"`
	SelectCountTemplete func(id string, remark string) (int64, error)                          `mapperParams:"id,remark"`
}

func (it KVMapper) New() KVMapper {
	core_util.LoadMapper(&it)
	return it
}
