package service

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
)

type MallClassPageDTO struct {
	vo.Pageable
	Name string `json:"name"`
}

type MallClassService struct {
	Add     func(arg model.MallClass) error
	Update  func(arg model.MallClass) error
	Delete  func(id string) error
	Find    func(id string) (vo.MallClassVO, error)
	Page    func(arg MallClassPageDTO) (vo.PageVO, error)
	FindAll func() ([]vo.MallClassVO, error)
}
