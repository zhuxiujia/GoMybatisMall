package core_service

import (
	"errors"
	"github.com/zhuxiujia/GoMybatis"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/utils"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_context"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/dao"
	"time"
)

type KVServiceImpl struct {
	service.KVService `bean:"KVService"`
	dao.KVMapper
}

func (it *KVServiceImpl) Init() {
	it.KVMapper = it.KVMapper.New()
	it.Add = func(arg vo.KVVO) error {
		if arg.Id == "" {
			arg.Id = utils.CreateUUID()
		}
		arg.KV.CreateTime = time.Now()
		return it.KVMapper.InsertTemplete(arg.KV)
	}

	it.Update = func(arg vo.KVVO) error {
		if arg.Id == "" {
			return errors.New("id不能为空！")
		}
		return it.KVMapper.UpdateTemplete(arg.KV)
	}

	it.Delete = func(id string) error {
		if id == "" {
			return errors.New("id不能为空！")
		}
		return it.KVMapper.DeleteTemplete(id)
	}

	it.Find = func(id string) (result vo.KVVO, e error) {
		if id == "" {
			return result, errors.New("id不能为空！")
		}
		r, e := it.KVMapper.SelectTemplete(id)
		if e != nil {
			return result, e
		}
		result.KV = r
		return result, nil
	}

	it.Finds = func(ids []string) (kvvos []vo.KVVO, e error) {
		kvvos = []vo.KVVO{}
		r, e := it.KVMapper.SelectByIds(ids)
		if e != nil {
			return kvvos, e
		}
		if r != nil {
			for _, item := range r {
				kvvos = append(kvvos, vo.KVVO{
					KV: item,
				})
			}
		}
		return kvvos, e
	}

	it.FindIdLike = func(id string) (result []vo.KVVO, e error) {
		if id == "" {
			return result, errors.New("id不能为空！")
		}
		r, e := it.KVMapper.SelectIdLike(id)
		if e != nil {
			return result, e
		}
		if r != nil {
			for _, item := range r {
				result = append(result, vo.KVVO{
					KV: item,
				})
			}
		}
		return result, nil
	}

	it.Page = func(arg service.KVPageDTO) (result vo.PageVO, e error) {
		data, e := it.KVMapper.SelectPageTemplete(arg.Id, arg.Remark, arg.Page, arg.PageSize)
		if e != nil {
			return result, e
		}
		total, e := it.KVMapper.SelectCountTemplete(arg.Id, arg.Remark)
		if e != nil {
			return result, e
		}
		result = result.New(arg.Pageable, total, data)
		return result, nil
	}

	//finish
	GoMybatis.AopProxyService(&it.KVService, core_context.Engine)
	core_util.ScanInject("KVServiceImpl", it)
}
