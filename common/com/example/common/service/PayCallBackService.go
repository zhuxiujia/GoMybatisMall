package service

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
)

type PayCallBackPageDTO struct {
	vo.Pageable
}

type PayCallBackService struct {
	Add  func(act vo.PayCallBackVO) error
	Page func(req PayCallBackPageDTO) (vo.PayCallBackPageVO, error)
}
