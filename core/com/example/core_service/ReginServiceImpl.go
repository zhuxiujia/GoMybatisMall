package core_service

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/dao"
)

type RegionServiceImpl struct {
	service.RegionService `bean:"RegionService"`
	dao.RegionMapper
}

func (it *RegionServiceImpl) Init() {
	it.RegionMapper = it.RegionMapper.New()
	it.Find = func(arg string) (result []vo.ReginVO, e error) {
		//获取省
		//获取市
		//获取区
		//var regionvos = []vo.ReginVO{}
		provinces, e := it.RegionMapper.SelectRegionProvince()
		if e != nil {
			return result, e
		}
		citys, e := it.RegionMapper.SelectRegionCity()
		if e != nil {
			return result, e
		}
		areas, e := it.RegionMapper.SelectRegionArea()
		if e != nil {
			return result, e
		}
		for _, v := range provinces {
			//市
			var citys_child = []vo.ReginVO{}
			for _, c := range citys {
				if v.ProvinceId == c.ProvinceId {
					var area_child = []vo.ReginVO{}
					for _, a := range areas {
						if c.CityId == a.CityId {
							area_child = append(area_child, vo.ReginVO{
								Value: a.AreaId,
								Label: a.Area,
							})
						}
					}
					citys_child = append(citys_child, vo.ReginVO{
						Value:    c.CityId,
						Label:    c.City,
						Children: area_child,
					})
				}
			}
			//省
			result = append(result, vo.ReginVO{
				Value:    v.ProvinceId,
				Label:    v.Province,
				Children: citys_child,
			})
		}
		return result, nil
	}

	core_util.ScanInject("RegionServiceImpl", it)
}
