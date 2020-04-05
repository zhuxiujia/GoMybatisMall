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

type PropertyServiceImpl struct {
	service.PropertyService `bean:"PropertyService"`
	propertyMapper          dao.PropertyMapper
	propertyRecordMapper    dao.PropertyRecordMapper
}

func (it *PropertyServiceImpl) Init() {
	it.propertyMapper = it.propertyMapper.New()
	it.propertyRecordMapper = it.propertyRecordMapper.New()
	it.Insert = func(arg vo.PropertyVO) error {
		if arg.UserId == "" {
			return errors.New("Insert arg userId can not be empty!")
		}
		arg.Id = utils.CreateUUID()
		arg.CreateTime = time.Now()
		arg.DeleteFlag = 1
		return it.propertyMapper.InsertTemplete(arg.Property)
	}

	it.Update = func(arg vo.PropertyVO) (int64, error) {
		if arg.Id == "" {
			return 0, errors.New("PropertyService arg id can not be empty!")
		}
		var e = it.propertyMapper.UpdateTemplete(arg.Property)
		if e != nil {
			return 0, e
		}
		return 0, nil
	}

	it.UpdateAmount = func(arg service.UpdateAmountDTO) (i int64, e error) {
		if arg.AddReduceAmount != 0 {
			var record = model.PropertyRecord{
				Amount:     arg.AddReduceAmount,
				Id:         arg.OrderId,
				CreateTime: time.Now(),
				DeleteFlag: 1,
			}
			var e = it.propertyRecordMapper.InsertTemplete(record)
			if e != nil {
				return 0, e
			}
			arg.PropertyVO.Amount = arg.PropertyVO.Amount + arg.AddReduceAmount
			return it.Update(arg.PropertyVO)
		}
		return i, e
	}

	it.FindByUser = func(user_id string) (result vo.PropertyVO, e error) {
		property, e := it.propertyMapper.SelectByUserId(user_id)
		if e != nil {
			return result, e
		}
		result.Property = property
		return result, nil
	}
	it.Find = func(id string) (propertyVO vo.PropertyVO, e error) {
		property, e := it.propertyMapper.SelectTemplete(id)
		if e != nil {
			return propertyVO, e
		}
		propertyVO.Property = property
		return propertyVO, nil
	}
	GoMybatis.AopProxyService(&it.PropertyService, core_context.Engine)
	core_util.ScanInject("PropertyServiceImpl", it)
}
