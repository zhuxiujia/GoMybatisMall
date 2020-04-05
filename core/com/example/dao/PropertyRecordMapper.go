package dao

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
)

type PropertyRecordMapper struct {
	InsertTemplete func(arg model.PropertyRecord) error
}

func (it PropertyRecordMapper) New() PropertyRecordMapper {
	core_util.LoadMapper(&it)
	return it
}
