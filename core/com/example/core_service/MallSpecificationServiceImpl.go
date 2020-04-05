package core_service

import (
	"errors"
	"github.com/zhuxiujia/GoMybatis"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_context"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/dao"
	"time"
)

type MallSpecificationServiceImpl struct {
	service.MallSpecificationService `bean:"MallSpecificationService"`
	mallSpecificationMapper          dao.MallSpecificationMapper
}

func (it *MallSpecificationServiceImpl) Init() {
	it.mallSpecificationMapper = it.mallSpecificationMapper.New()

	it.Add = func(arg model.MallSpecification) error {
		if arg.Id == "" {
			return errors.New("id为空")
		}
		arg.DeleteFlag = 1
		arg.CreateTime = time.Now()
		return it.mallSpecificationMapper.InsertTemplete(arg)
	}

	it.Delete = func(id string) error {
		return it.mallSpecificationMapper.DeleteTemplete(id)
	}

	it.DeleteBySkuIdName = func(arg service.MallSpecificationDeleteDTO) error {
		return it.mallSpecificationMapper.DeleteBySkuIdAndName(arg)
	}

	it.DeleteBySkuId = func(skuId string) error {
		return it.mallSpecificationMapper.DeleteBySkuId(skuId)
	}

	it.Update = func(arg model.MallSpecification) error {
		if arg.Id == "" {
			return errors.New("id为空")
		}
		return it.mallSpecificationMapper.UpdateTemplete(arg)
	}

	it.Find = func(id string) (vo vo.MallSpecificationVO, e error) {
		data, e := it.mallSpecificationMapper.SelectTemplete(id)
		if e != nil {
			return vo, e
		}
		vo.MallSpecification = data
		return vo, nil
	}

	it.Finds = func(ids []string) (vos map[string]*vo.MallSpecificationVO, e error) {
		if ids == nil || len(ids) == 0 {
			return vos, e
		}
		vos = make(map[string]*vo.MallSpecificationVO)
		data, e := it.mallSpecificationMapper.SelectByIds(ids)
		if e != nil {
			return vos, e
		}
		if data != nil {
			for _, v := range data {
				var vo = vo.MallSpecificationVO{
					MallSpecification: v,
				}
				vos[v.Id] = &vo
			}
		}
		return vos, nil
	}

	it.FindBySkuIds = func(skuIds []string) (vos map[string]*[]vo.MallSpecificationVO, e error) {
		if skuIds == nil || len(skuIds) == 0 {
			return vos, e
		}
		vos = make(map[string]*[]vo.MallSpecificationVO)
		data, e := it.mallSpecificationMapper.SelectBySkuIds(skuIds)
		if e != nil {
			return vos, e
		}
		if data != nil {
			for _, v := range data {
				var spec = vo.MallSpecificationVO{
					MallSpecification: v,
				}
				var arr = vos[v.SkuId]
				if arr == nil {
					var a = make([]vo.MallSpecificationVO, 0)
					arr = &a
				}
				*arr = append(*arr, spec)
				vos[v.SkuId] = arr
			}
		}
		return vos, nil
	}

	//finish
	GoMybatis.AopProxyService(&it.MallSpecificationService, core_context.Engine)
	core_util.ScanInject("MallSpecificationServiceImpl", it)
}
