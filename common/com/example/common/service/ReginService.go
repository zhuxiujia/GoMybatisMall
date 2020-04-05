package service

import "github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"

type RegionService struct {
	Find func(arg string) ([]vo.ReginVO, error)
}
