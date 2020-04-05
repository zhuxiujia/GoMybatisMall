package c_admin

import (
	"errors"
	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/enum"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/easy_mvc"
)

type MallOrderController struct {
	easy_mvc.Controller `doc:"(后台接口)商城订单"`

	MallOrderService         *service.MallOrderService         `inject:"MallOrderService"`
	MallSpecificationService *service.MallSpecificationService `inject:"MallSpecificationService"`

	Page        func(fuzzy string, sku_id string, status *int, time_start string, time_end string, page int, size int) interface{}                                                                                                   `path:"/admin/user/mall/order/page" arg:"fuzzy,sku_id,status,time_start,time_end,page:0,size:5" doc:"商城订单列表" doc_arg:""`
	Detail      func(id string) interface{}                                                                                                                                                                                          `path:"/admin/user/mall/order/detail" arg:"id" doc:"商城订单详情" doc_arg:""`
	Edit        func(id string, receive_name string, receive_phone string, receive_address string, express_num string, express_name string, status *int, sku_specification_id string, order_amount *int, remark *string) interface{} `path:"/admin/user/mall/order/edit" arg:"id,receive_name,receive_phone,receive_address,express_num,express_name,status,sku_specification_id,order_amount,remark" doc:"商城订单修改" doc_arg:""`
	ExpressList func() interface{}                                                                                                                                                                                                   `path:"/admin/user/mall/order/express" arg:"" doc:"商城快递列表" doc_arg:""`
	CancelOrder func(id string) interface{}                                                                                                                                                                                          `path:"/admin/user/mall/order/cancel" arg:"id" doc:"商城订单取消" doc_arg:""`
}

func (it *MallOrderController) Routers() {

	it.Page = func(fuzzy string, sku_id string, status *int, startTimeStr string, endTimeStr string, page int, size int) interface{} {
		var timeAble = vo.TimeRangeable{}.New(startTimeStr, endTimeStr)
		var startTime = timeAble.TimeStart()
		var endTime = timeAble.TimeEnd()
		var dto = service.MallOrderPageDTO{
			Fuzzy:     fuzzy,
			SkuId:     sku_id,
			Status:    status,
			StartTime: startTime,
			EndTime:   endTime,
			Pageable:  vo.Pageable{}.New(page, size),
		}
		var data, e = it.MallOrderService.Page(dto)
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(data)
	}
	it.Detail = func(id string) interface{} {
		var data, e = it.MallOrderService.Find(id)
		if e != nil {
			return e
		}
		if data.SkuId != "" {
			specMap, e := it.MallSpecificationService.FindBySkuIds([]string{data.SkuId})
			if e != nil {
				return e
			}
			if specMap != nil && len(specMap) > 0 {
				var specs = specMap[data.SkuId]
				if specs != nil && len(*specs) > 0 {
					data.MallSpecificationVO = &(*specs)[0]
				}
			}

		}
		return vo.ResultVO{}.NewSuccess(data)
	}

	it.Edit = func(id string, receive_name string, receive_phone string, receive_address string, express_num string, express_name string, status *int, sku_specification_id string, order_amount *int, remark *string) interface{} {
		var data, e = it.MallOrderService.Find(id)
		if e != nil {
			return e
		}
		if data.Id == "" {
			return errors.New("订单:" + id + "不存在!")
		}

		if data.Status == enum.MallOrderStatus_Paying && order_amount != nil {
			//支付中订单才允许改价
			data.OrderAmount = *order_amount
		}
		data.MallOrder.ReceiveName = receive_name
		data.MallOrder.ReceivePhone = receive_phone
		data.MallOrder.ReceiveAddress = receive_address

		data.MallOrder.ExpressNum = express_num
		data.MallOrder.ExpressName = express_name
		data.SkuSpecificationId = sku_specification_id

		if remark != nil {
			data.Remark = *remark
		}

		if status != nil {
			data.Status = enum.NewMallOrderStatus(*status)
		}
		_, e = it.MallOrderService.Update(data.MallOrder)
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(nil)
	}
	it.ExpressList = func() interface{} {
		var m = []string{}
		m = append(m, "中国邮政")
		m = append(m, "顺丰快递")
		m = append(m, "中通快递")
		m = append(m, "圆通快递")
		m = append(m, "申通快递")
		m = append(m, "韵达快递")
		m = append(m, "汇通快递")
		m = append(m, "天天快递")
		m = append(m, "宅急送快递")
		m = append(m, "德邦物流")
		m = append(m, "国通快递")
		m = append(m, "中铁快运")
		m = append(m, "E邮宝")
		m = append(m, "其他快递")
		return vo.ResultVO{}.NewSuccess(m)
	}

	it.CancelOrder = func(id string) interface{} {
		e := it.MallOrderService.Cancel(id)
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(nil)
	}

	it.Init(it)
	app_context.Context.Provide(&inject.Object{
		Value: it,
	})
}
