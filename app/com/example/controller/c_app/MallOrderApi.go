package c_app

import (
	"errors"
	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/enum"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/easy_mvc"
)

type MallOrderApi struct {
	easy_mvc.Controller `doc:"商城订单API"`

	MallOrderService *service.MallOrderService `inject:"MallOrderService"`

	Page   func(login_id string, status *int, page int, size int) interface{} `path:"/api/user/mall/order/page" arg:"login_id,status,page:0,size:5" doc:"mall_specification_vo商品规格，mall_cover_image_vos是商品缩略图，sku_id:商品id,sku_specification_id:商品规格id,sku_num:订购数量,receive_name:收货人名称,receive_phone:收货人手机号,receive_address:收货地址,amount:下单金额,remark:下单备注" doc_arg:"login_id:_,status:-1失败0支付中1待发货2已发货3已完成4已取消"`
	Detail func(id string) interface{}                                        `path:"/api/user/mall/order/detail" arg:"id" doc:"" doc_arg:""`
	Order  func(login_id string,
		receive_name string,
		receive_phone string,
		receive_address string,
		sku_id string,
		sku_num int,
		sku_specification_id string,
		amount int,
		pay_type string,
		remark string) interface{} `path:"/api/user/mall/order/add" arg:"login_id,receive_name,receive_phone,receive_address,sku_id,sku_num,sku_specification_id,amount,pay_type,remark" doc:"下单" doc_arg:"login_id:_,sku_id:商品id,sku_specification_id:商品规格id,sku_num:订购数量,receive_name:收货人名称,receive_phone:收货人手机号,receive_address:收货地址,amount:下单金额,remark:下单备注"`
	Edit func(id string) interface{} `path:"/api/user/mall/order/edit" arg:"id" doc:"修改订单为已完成（收货成功）" doc_arg:"status:-1失败0支付中1待发货2已发货3已完成4已取消"`
}

func (it *MallOrderApi) Routers() {

	it.Page = func(login_id string, status *int, page int, size int) interface{} {
		var data, e = it.MallOrderService.Page(service.MallOrderPageDTO{
			UserId:   login_id,
			Status:   status,
			Pageable: vo.Pageable{}.New(page, size),
		})
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
		if data.Id == "" {
			return vo.ResultVO{}.NewSuccess(nil)
		}
		return vo.ResultVO{}.NewSuccess(data)
	}
	it.Order = func(login_id string, receive_name string, receive_phone string, receive_address string, sku_id string, sku_num int, sku_specification_id string, amount int, pay_type string, remark string) interface{} {
		var data, e = it.MallOrderService.Order(service.MallOrderDTO{
			SkuId:              sku_id,
			UserId:             login_id,
			SkuSpecificationId: sku_specification_id,
			ReceiveName:        receive_name,
			ReceivePhone:       receive_phone,
			ReceiveAddress:     receive_address,
			SkuNum:             sku_num,
			Amount:             amount,
			PayType:            pay_type,
			Remark:             remark,
		})
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(data)
	}

	it.Edit = func(id string) interface{} {
		var data, e = it.MallOrderService.Find(id)
		if e != nil {
			return e
		}
		if data.Id == "" {
			return errors.New("订单:" + id + "不存在!")
		}
		if data.Status == enum.MallOrderStatus_COMPLETE {
			return errors.New("订单已完成！")
		}
		data.Status = enum.MallOrderStatus_COMPLETE
		_, e = it.MallOrderService.Update(data.MallOrder)
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
