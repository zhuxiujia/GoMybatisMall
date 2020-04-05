package core_service

import (
	"errors"
	"github.com/zhuxiujia/GoMybatis"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/utils"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_context"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/dao"
	"time"
)

type MallClassServiceImpl struct {
	service.MallClassService `bean:"MallClassService"`

	mallBrandMapper dao.MallClassMapper
}

func (it *MallClassServiceImpl) Init() {
	it.mallBrandMapper = it.mallBrandMapper.New()

	it.Add = func(arg model.MallClass) error {
		if arg.Name == "" {
			return errors.New("名称不能为空！")
		}
		arg.Id = utils.CreateUUID()
		arg.DeleteFlag = 1
		arg.CreateTime = time.Now()
		return it.mallBrandMapper.InsertTemplete(arg)
	}

	it.Find = func(id string) (vo vo.MallClassVO, e error) {
		data, e := it.mallBrandMapper.SelectTemplete(id)
		if e != nil {
			return vo, e
		}
		vo.MallClass = data
		return vo, e
	}

	it.Page = func(arg service.MallClassPageDTO) (pageVO vo.PageVO, e error) {
		data, e := it.mallBrandMapper.SelectByCondition(arg)
		if e != nil {
			return pageVO, e
		}
		total, e := it.mallBrandMapper.CountByCondition(arg)
		if e != nil {
			return pageVO, e
		}
		pageVO = pageVO.New(arg.Pageable, total, data)
		return pageVO, e
	}

	it.Update = func(arg model.MallClass) error {
		if arg.Id == "" {
			return errors.New("id为空")
		}
		return it.mallBrandMapper.UpdateTemplete(arg)
	}

	it.Delete = func(id string) error {
		return it.mallBrandMapper.DeleteTemplete(id)
	}

	it.FindAll = func() (vos []vo.MallClassVO, e error) {
		data, e := it.mallBrandMapper.FindAll()
		if e != nil {
			return vos, e
		}
		if data != nil {
			utils.ConvertToVOs(&vos, data)
		}
		return vos, e
	}

	//finish
	GoMybatis.AopProxyService(&it.MallClassService, core_context.Engine)
	core_util.ScanInject("MallClassServiceImpl", it)
}
