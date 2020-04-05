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

type MallSkuServiceImpl struct {
	service.MallSkuService `bean:"MallSkuService"`
	mallSkuMapper          dao.MallSkuMapper

	MallCoverImageService    *service.MallCoverImageService    `inject:"MallCoverImageService"`
	MallSpecificationService *service.MallSpecificationService `inject:"MallSpecificationService"`
}

func (it *MallSkuServiceImpl) Init() {
	it.mallSkuMapper = it.mallSkuMapper.New()

	it.Add = func(arg service.MallSkuAddDTO) error {
		if arg.MallSku.Id == "" {
			return errors.New("id为空")
		}
		if arg.MallSku.RemainNum == 0 {
			return errors.New("剩余库存不可为0！")
		}
		if arg.Specifications == nil || len(*arg.Specifications) == 0 {
			return errors.New("必须添加至少一个规格！")
		}
		if (*arg.Specifications)[0] == "" {
			return errors.New("必须添加至少一个规格！")
		}
		arg.MallSku.RemainNum = arg.MallSku.TotalNum
		arg.MallSku.DeleteFlag = 1
		arg.MallSku.CreateTime = time.Now()

		if arg.Images != nil && len(*arg.Images) != 0 {
			for _, v := range *arg.Images {
				var e = it.MallCoverImageService.Add(model.MallCoverImage{
					Id:    utils.CreateUUID(),
					Img:   v,
					SkuId: arg.MallSku.Id,
				})
				if e != nil {
					return e
				}
			}
		}
		if arg.Specifications != nil && len(*arg.Specifications) != 0 {
			for _, v := range *arg.Specifications {
				var e = it.MallSpecificationService.Add(model.MallSpecification{
					Id:    utils.CreateUUID(),
					SkuId: arg.MallSku.Id,
					Name:  v,
				})
				if e != nil {
					return e
				}
			}
		}
		return it.mallSkuMapper.InsertTemplete(arg.MallSku)
	}

	it.Delete = func(id string) error {
		return it.mallSkuMapper.DeleteTemplete(id)
	}

	it.Update = func(arg service.MallSkuAddDTO) (int64, error) {
		if arg.MallSku.Id == "" {
			return 0, errors.New("id为空")
		}
		var updateNum, e = it.mallSkuMapper.UpdateTemplete(arg.MallSku)
		if e != nil {
			return 0, e
		}
		if arg.Images != nil && len(*arg.Images) != 0 {
			it.MallCoverImageService.DeleteBySkuId(arg.MallSku.Id)
			for _, v := range *arg.Images {
				var e = it.MallCoverImageService.Add(model.MallCoverImage{
					Id:    utils.CreateUUID(),
					Img:   v,
					SkuId: arg.MallSku.Id,
				})
				if e != nil {
					return updateNum, e
				}
			}
		}
		if arg.Specifications != nil && len(*arg.Specifications) != 0 {
			it.MallSpecificationService.DeleteBySkuId(arg.MallSku.Id)
			for _, v := range *arg.Specifications {
				var e = it.MallSpecificationService.Add(model.MallSpecification{
					Id:    utils.CreateUUID(),
					SkuId: arg.MallSku.Id,
					Name:  v,
				})
				if e != nil {
					return updateNum, e
				}
			}
		}
		return updateNum, e
	}

	it.Find = func(id string) (mallSkuVO vo.MallSkuVO, e error) {
		if id == "" {
			return mallSkuVO, errors.New("Find() arg id can not be empty!")
		}
		data, e := it.mallSkuMapper.SelectTemplete(id)
		if e != nil {
			return mallSkuVO, e
		}
		mallSkuVO.MallSku = data

		var skuIds = utils.GetIds(data, "Id")
		//set image data
		mallCoverImages, e := it.MallCoverImageService.FindBySkuId(skuIds)
		if e != nil {
			return mallSkuVO, e
		}
		mallSpecs, e := it.MallSpecificationService.FindBySkuIds(skuIds)
		if e != nil {
			return mallSkuVO, e
		}
		utils.ConvertField(&mallSkuVO,
			utils.Convert{}.New("Id", "MallCoverImageVOs", mallCoverImages),
			utils.Convert{}.New("Id", "MallSpecificationVOs", mallSpecs))
		return mallSkuVO, e
	}

	it.Finds = func(ids []string) (vos map[string]*vo.MallSkuVO, e error) {
		if ids == nil || len(ids) == 0 {
			return vos, e
		}
		vos = make(map[string]*vo.MallSkuVO)
		data, e := it.mallSkuMapper.SelectByIds(ids)
		if e != nil {
			return vos, e
		}
		if data != nil {
			for _, v := range data {
				var vo = vo.MallSkuVO{
					MallSku: v,
				}
				vos[v.Id] = &vo
			}
		}
		return vos, nil
	}

	it.Page = func(arg service.MallSkuPageDTO) (pageVO vo.PageVO, e error) {
		data, e := it.mallSkuMapper.SelectByCondition(arg)
		if e != nil {
			return pageVO, e
		}
		total, e := it.mallSkuMapper.CountByCondition(arg)
		if e != nil {
			return pageVO, e
		}
		var skuIds = utils.GetIds(data, "Id")
		//set image data
		mallCoverImages, e := it.MallCoverImageService.FindBySkuId(skuIds)
		if e != nil {
			return pageVO, e
		}
		mallSpecs, e := it.MallSpecificationService.FindBySkuIds(skuIds)
		if e != nil {
			return pageVO, e
		}
		var dataVOs []vo.MallSkuVO
		utils.ConvertToVOs(&dataVOs, data)

		utils.ConvertField(dataVOs,
			utils.Convert{}.New("Id", "MallCoverImageVOs", mallCoverImages),
			utils.Convert{}.New("Id", "MallSpecificationVOs", mallSpecs),
		)
		pageVO = pageVO.New(arg.Pageable, total, dataVOs)
		return pageVO, nil
	}

	//finish
	GoMybatis.AopProxyService(&it.MallSkuService, core_context.Engine)
	core_util.ScanInject("MallSkuServiceImpl", it)
}
