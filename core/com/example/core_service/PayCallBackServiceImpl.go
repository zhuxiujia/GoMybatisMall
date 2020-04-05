package core_service

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/dao"
	"time"
)

type PayCallBackServiceImpl struct {
	service.PayCallBackService `bean:"PayCallBackService"`

	PayCallBackMapper dao.PayCallBackMapper
}

func (it *PayCallBackServiceImpl) Init() {
	it.PayCallBackMapper = it.PayCallBackMapper.New()
	it.Page = func(req service.PayCallBackPageDTO) (vo vo.PayCallBackPageVO, e error) {
		data, e := it.PayCallBackMapper.SelectPageTemplete(req.Page, req.PageSize)
		if e != nil {
			return vo, e
		}
		total, e := it.PayCallBackMapper.SelectCountTemplete()
		if e != nil {
			return vo, e
		}
		vo = vo.New(req.Pageable, total, data)
		return vo, e
	}
	it.Add = func(act vo.PayCallBackVO) error {
		act.CreateTime = time.Now()
		act.DeleteFlag = 1
		return it.PayCallBackMapper.InsertTemplete(act.PayCallBack)
	}

	core_util.ScanInject("PayCallBackServiceImpl", it)
}
