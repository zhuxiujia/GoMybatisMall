package dao

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
)

type RegionMapper struct {
	SelectRegion func(province_id string, city_id string, area_id string) ([]model.Regin, error) `mapperParams:"province_id,city_id,area_id"`

	SelectRegionProvince func() ([]model.RegionProvince, error)
	SelectRegionCity     func() ([]model.RegionCity, error)
	SelectRegionArea     func() ([]model.RegionArea, error)
}

func (it RegionMapper) New() RegionMapper {
	core_util.LoadMapper(&it)
	return it
}
