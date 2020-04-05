package core_service

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/zhuxiujia/GoMybatis"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/enum"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/utils"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_context"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/dao"
	"log"
	"net/url"
	"time"
)

type MallOrderServiceImpl struct {
	service.MallOrderService `bean:"MallOrderService"`
	mallOrderMapper          dao.MallOrderMapper
	CoreConfig               *core_context.CoreConfig          `inject:"CoreConfig"`
	MallSkuService           *service.MallSkuService           `inject:"MallSkuService"`
	MallSpecificationService *service.MallSpecificationService `inject:"MallSpecificationService"`
	MallCoverImageService    *service.MallCoverImageService    `inject:"MallCoverImageService"`
	KVService                *service.KVService                `inject:"KVService"`
}

func (it *MallOrderServiceImpl) Init() {
	it.mallOrderMapper = it.mallOrderMapper.New()

	it.Add = func(arg model.MallOrder) error {
		if arg.Id == "" {
			return errors.New("id为空")
		}
		arg.DeleteFlag = 1
		arg.CreateTime = time.Now()
		return it.mallOrderMapper.InsertTemplete(arg)
	}

	it.Delete = func(id string) error {
		return it.mallOrderMapper.DeleteTemplete(id)
	}

	it.Update = func(arg model.MallOrder) (int64, error) {
		if arg.Id == "" {
			return 0, errors.New("id为空")
		}
		order, e := it.mallOrderMapper.SelectTemplete(arg.Id)
		if e != nil {
			return 0, e
		}
		if (arg.Status == enum.MallOrderStatus_ING || arg.Status == enum.MallOrderStatus_COMPLETE) && (order.Status == enum.MallOrderStatus_Paying || order.Status == enum.MallOrderStatus_Fail || order.Status == enum.MallOrderStatus_Cancel) {
			return 0, errors.New("支付中或失败的订单不可改为已发货！")
		}
		return it.mallOrderMapper.UpdateTemplete(arg)
	}

	it.Find = func(id string) (orderVO vo.MallOrderVO, e error) {
		data, e := it.mallOrderMapper.SelectTemplete(id)
		if e != nil {
			return orderVO, e
		}
		orderVO.MallOrder = data

		skuMap, e := it.MallSkuService.Finds(utils.GetIds(data, "SkuId"))
		if e != nil {
			return orderVO, e
		}
		skuSpecificationMap, e := it.MallSpecificationService.Finds(utils.GetIds(data, "SkuSpecificationId"))
		if e != nil {
			return orderVO, e
		}
		utils.ConvertField(&orderVO,
			utils.Convert{}.New("SkuId", "MallSkuVO", skuMap),
			utils.Convert{}.New("SkuSpecificationId", "MallSpecificationVO", skuSpecificationMap))
		return orderVO, e
	}

	it.Page = func(arg service.MallOrderPageDTO) (pageVO vo.PageVO, e error) {
		data, e := it.mallOrderMapper.SelectByCondition(arg)
		if e != nil {
			return pageVO, e
		}
		total, e := it.mallOrderMapper.CountByCondition(arg)
		if e != nil {
			return pageVO, e
		}

		skuMap, e := it.MallSkuService.Finds(utils.GetIds(data, "SkuId"))
		if e != nil {
			return pageVO, e
		}
		skuSpecificationMap, e := it.MallSpecificationService.Finds(utils.GetIds(data, "SkuSpecificationId"))
		if e != nil {
			return pageVO, e
		}

		imgMap, e := it.MallCoverImageService.FindBySkuId(utils.GetIds(data, "SkuId"))
		if e != nil {
			return pageVO, e
		}

		var orderVOs = []vo.MallOrderVO{}
		for _, v := range data {
			orderVOs = append(orderVOs, vo.MallOrderVO{
				MallOrder: v,
			})
		}
		utils.ConvertField(orderVOs,
			utils.Convert{}.New("SkuId", "MallSkuVO", skuMap),
			utils.Convert{}.New("SkuSpecificationId", "MallSpecificationVO", skuSpecificationMap))
		for _, v := range orderVOs {
			if v.MallSkuVO != nil {
				v.MallSkuVO.MallCoverImageVOs = imgMap[v.SkuId]
			}
		}

		pageVO = pageVO.New(arg.Pageable, total, orderVOs)
		return pageVO, nil
	}

	it.Order = func(arg service.MallOrderDTO) (orderVO vo.MallOrderVO, e error) {
		//1.商品验证
		//1.1.先判断商品数量
		sku, e := it.MallSkuService.Find(arg.SkuId)
		if e != nil {
			return orderVO, e
		}
		var remainNum = sku.RemainNum - arg.SkuNum
		e = it.check(sku, arg, remainNum)
		if e != nil {
			return orderVO, e
		}
		//3.再添加订单
		//filter user
		var order_title = "mall-"
		var order = model.MallOrder{
			Id:                 order_title + utils.CreateUUID()[len(order_title):],
			UserId:             arg.UserId,
			ReceiveName:        arg.ReceiveName,
			ReceivePhone:       arg.ReceivePhone,
			ReceiveAddress:     arg.ReceiveAddress,
			SkuId:              sku.Id,
			SkuNum:             arg.SkuNum,
			SkuSpecificationId: arg.SkuSpecificationId,
			OrderAmount:        sku.Amount * arg.SkuNum,
			PayType:            arg.PayType,
			Status:             enum.MallOrderStatus_Paying,
			Remark:             arg.Remark,
			CreateTime:         time.Now(),
			DeleteFlag:         1,
		}
		if arg.PayType == "" {
			//支付宝
			arg.PayType = "zfb"
		}
		//金额转换为 第三方支付对应金额（元/分）
		var thirdPayRealAmount = core_util.ToThirdPayAmount(arg.PayType, order.OrderAmount)
		//金额过滤
		if it.KVService != nil {
			var supers, err = it.KVService.Find("user_0.01")
			if err == nil {
				thirdPayRealAmount = core_util.FilterAmount(supers, arg.ReceivePhone, arg.PayType, thirdPayRealAmount)
			}
		}

		var zfbValues = url.Values{}
		zfbValues.Add("order_id", order.Id)
		zfbValues.Add("pay_type", order.PayType)
		zfbValues.Add("body", sku.Title)
		zfbValues.Add("Subject", sku.Title)
		zfbValues.Add("totalAmount", fmt.Sprint(thirdPayRealAmount))
		zfbValues.Add("debug", arg.Debug)

		//支付宝 支付
		var bytes, err = utils.HttpPost(it.CoreConfig.CashierUrl, zfbValues)
		if err != nil {
			var info = "支付失败：" + err.Error()
			log.Println(info)
			return orderVO, errors.New(info)
		}
		var data = string(bytes)
		if data == "" {
			return orderVO, errors.New("支付宝支付失败：订单创建失败")
		}
		order.PayLink = data
		//存储订单
		e = it.mallOrderMapper.InsertTemplete(order)
		if e != nil {
			return orderVO, e
		}
		return vo.MallOrderVO{
			MallOrder: order,
		}, nil
	}

	//订单通知
	it.SyncOrderByQueue = func(notify vo.AbstractNotifyVO) error {
		if notify.OutTradeNo == "" {
			return nil
		}
		//更新订单
		orderVO, e := it.Find(notify.OutTradeNo)
		if e != nil {
			return e
		}
		//状态拦截
		if orderVO.Status != enum.MallOrderStatus_Paying {
			log.Println("订单已完成:" + orderVO.Id)
			return nil
		}
		//第三方状态判断
		switch notify.Status {
		case enum.NotifyStatus_SUCCESS:
			orderVO.Status = enum.MallOrderStatus_Redy //success
		case enum.NotifyStatus_PAYING:
			orderVO.Status = enum.MallOrderStatus_Paying //paying
		case enum.NotifyStatus_FAIL:
			orderVO.Status = enum.MallOrderStatus_Fail //fail
		}
		num, e := it.mallOrderMapper.UpdateTemplete(orderVO.MallOrder)
		if e != nil {
			return e
		}
		if num == 0 {
			return errors.New("失败，商城订单更新失败：" + orderVO.Id)
		}
		//下单成功减库存,如果商品剩余库存为0，默认将商品下架
		if orderVO.Status == enum.MallOrderStatus_Redy {
			sku, e := it.MallSkuService.Find(orderVO.SkuId)
			if e != nil {
				return e
			}
			sku.RemainNum = sku.RemainNum - orderVO.SkuNum
			_, e = it.MallSkuService.Update(service.MallSkuAddDTO{
				MallSku: sku.MallSku,
			})
			sku, e = it.MallSkuService.Find(orderVO.SkuId)
			if sku.RemainNum == 0 {
				sku.Status = enum.MallSkuStatus_OffLine
				it.MallSkuService.Update(service.MallSkuAddDTO{
					MallSku: sku.MallSku,
				})
			}
		}
		return nil
	}

	//取消订单
	it.Cancel = func(id string) error {
		if id == "" {
			return nil
		}
		//1.先更改订单状态
		data, e := it.mallOrderMapper.SelectTemplete(id)
		if e != nil {
			return e
		}
		data.Status = enum.MallOrderStatus_Cancel
		updateNum, e := it.mallOrderMapper.UpdateTemplete(data)
		if e != nil {
			return e
		}
		if updateNum == 0 {
			return errors.New("失败，商城订单更新失败：" + data.Id)
		}
		//2 修改商品库存
		sku, e := it.MallSkuService.Find(data.SkuId)
		if e != nil {
			return e
		}
		sku.RemainNum = sku.RemainNum + data.SkuNum
		_, e = it.MallSkuService.Update(service.MallSkuAddDTO{
			MallSku: sku.MallSku,
		})
		return e
	}

	//finish
	GoMybatis.AopProxyService(&it.MallOrderService, core_context.Engine)
	core_util.ScanInject("MallOrderServiceImpl", it)
}

func (it *MallOrderServiceImpl) check(sku vo.MallSkuVO, dto service.MallOrderDTO, remainNum int) error {
	if sku.Id == "" || sku.Status == enum.MallSkuStatus_OffLine {
		return errors.New("商品已下架！")
	}
	//1.再根据商品的兑换规则来过滤,判断是否只能兑换一次
	if sku.OrderTimeLimit > 0 {
		//TODO 统计兑换次数，限制购买次数
		var notEqualStatus = int(enum.MallOrderStatus_Cancel)
		var count, _ = it.mallOrderMapper.CountByUserIdAndIntegralProductId(service.MallOrderPageDTO{
			UserId: dto.UserId,
			SkuId:  sku.Id,
			Status: &notEqualStatus,
		})
		if count > int64(sku.OrderTimeLimit) {
			return errors.New("该商品只能下单" + fmt.Sprint(sku.OrderTimeLimit) + "次！")
		}
	}
	//2.再扣除用户资金\积分

	//是否小于产品限制积分
	if float64(dto.Amount)/float64(dto.SkuNum) < float64(sku.Amount) {
		var a100, _ = decimal.NewFromString("100")
		var needAmount = float64(dto.SkuNum) / float64(sku.Amount)
		var de = decimal.NewFromFloat(needAmount).Div(a100)
		return errors.New("认购不足以下单！!认购数量：" + fmt.Sprint(dto.SkuNum) + ",需要金额:" + de.String())
	}
	if remainNum == 0 || (sku.RemainNum-dto.SkuNum) < 0 {
		return errors.New("库存不足！!认购数量：" + fmt.Sprint(dto.SkuNum))
	}

	return nil
}
