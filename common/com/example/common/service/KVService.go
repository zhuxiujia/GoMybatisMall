package service

import "github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"

type KVPageDTO struct {
	vo.Pageable
	Id     string `json:"id"`
	Remark string `json:"remark"`
}

type KVService struct {
	Add        func(arg vo.KVVO) error
	Update     func(arg vo.KVVO) error
	Delete     func(id string) error
	Find       func(id string) (vo.KVVO, error)
	Finds      func(ids []string) ([]vo.KVVO, error)
	FindIdLike func(id string) ([]vo.KVVO, error)
	Page       func(id KVPageDTO) (vo.PageVO, error)
}
