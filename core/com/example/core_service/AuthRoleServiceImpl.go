package core_service

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/utils"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/dao"
	"time"
)

type AuthRoleServiceImpl struct {
	service.AuthRoleService `bean:"AuthRoleService"`

	authRoleMapper dao.AuthRoleMapper
}

func (it *AuthRoleServiceImpl) Init() {
	it.authRoleMapper = it.authRoleMapper.New()

	it.Delete = func(id string) error {
		return it.authRoleMapper.DeleteTemplete(id)
	}
	it.Add = func(arg model.AuthRole) error {
		arg.Id = utils.CreateUUID()
		arg.CreateTime = time.Now()
		arg.DeleteFlag = 1
		return it.authRoleMapper.InsertTemplete(arg)
	}
	it.Update = func(arg model.AuthRole) error {
		return it.authRoleMapper.UpdateTemplete(arg)
	}
	it.Find = func(id string) (roleVO vo.AuthRoleVO, e error) {
		date, e := it.authRoleMapper.SelectTemplete(id)
		if e != nil {
			return roleVO, e
		}
		roleVO.AuthRole = date
		return roleVO, e
	}
	it.Finds = func(ids []string) (vos []vo.AuthRoleVO, e error) {
		if ids == nil || len(ids) == 0 {
			return vos, e
		}
		datas, e := it.authRoleMapper.SelectByIds(ids)
		if e != nil {
			return vos, e
		}
		vos = make([]vo.AuthRoleVO, 0)
		utils.ConvertToVOs(&vos, datas)
		return vos, e
	}
	it.Page = func(arg service.AuthRolePageDTO) (pageVO vo.PageVO, e error) {
		data, e := it.authRoleMapper.SelectByCondition(arg)
		if e != nil {
			return pageVO, e
		}
		total, e := it.authRoleMapper.CountByCondition(arg)
		if e != nil {
			return pageVO, e
		}
		pageVO = pageVO.New(arg.Pageable, total, data)
		return pageVO, e
	}

	//end
	core_util.ScanInject("AuthRoleServiceImpl", it)
}
