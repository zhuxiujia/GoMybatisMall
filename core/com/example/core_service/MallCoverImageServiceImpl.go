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

type MallCoverImageServiceImpl struct {
	service.MallCoverImageService `bean:"MallCoverImageService"`
	mallCoverImageMapper          dao.MallCoverImageMapper
}

func (it *MallCoverImageServiceImpl) Init() {
	it.mallCoverImageMapper = it.mallCoverImageMapper.New()

	it.Add = func(arg model.MallCoverImage) error {
		if arg.Img == "" {
			return nil
		}
		if arg.Id == "" {
			return errors.New("id为空")
		}
		arg.DeleteFlag = 1
		arg.CreateTime = time.Now()
		return it.mallCoverImageMapper.InsertTemplete(arg)
	}

	it.Delete = func(id string) error {
		return it.mallCoverImageMapper.DeleteTemplete(id)
	}

	it.DeleteBySkuId = func(skuId string) error {
		return it.mallCoverImageMapper.DeleteBySkuId(skuId)
	}

	it.Update = func(arg model.MallCoverImage) error {
		if arg.Id == "" {
			return errors.New("id为空")
		}
		return it.mallCoverImageMapper.UpdateTemplete(arg)
	}

	it.Find = func(id string) (vo vo.MallCoverImageVO, e error) {
		data, e := it.mallCoverImageMapper.SelectTemplete(id)
		if e != nil {
			return vo, e
		}
		vo.MallCoverImage = data
		return vo, e
	}

	it.FindBySkuId = func(skuIds []string) (vos map[string]*[]vo.MallCoverImageVO, e error) {
		if skuIds == nil || len(skuIds) == 0 {
			return vos, e
		}
		datas, e := it.mallCoverImageMapper.SelectBySkuIds(skuIds)
		if e != nil {
			return vos, e
		}
		if datas != nil {
			vos = make(map[string]*[]vo.MallCoverImageVO)
			for _, v := range datas {
				var mallCoverImageVO = vo.MallCoverImageVO{
					MallCoverImage: v,
				}
				var arr = vos[v.SkuId]
				if arr == nil {
					var a = make([]vo.MallCoverImageVO, 0)
					arr = &a
				}
				*arr = append(*arr, mallCoverImageVO)
				vos[v.SkuId] = arr
			}
		}
		return vos, e
	}

	//finish
	GoMybatis.AopProxyService(&it.MallCoverImageService, core_context.Engine)
	core_util.ScanInject("MallCoverImageServiceImpl", it)
}
