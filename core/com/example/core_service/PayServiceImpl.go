package core_service

import (
	"encoding/json"
	"errors"
	"github.com/zhuxiujia/GoMybatis"
	"github.com/zhuxiujia/GoMybatis/utils"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_context"
	"github.com/zhuxiujia/GoMybatisMall/core/com/example/core_util"
	"log"
	"strings"
	"time"
)

const zfb_notify = "zfb_notify"
const wx_notify = "wx_notify"

type PayServiceImpl struct {
	service.PayService `bean:"PayService"`
	DoNotify           func(arg string, notifyVO vo.AbstractNotifyVO) error
	DoSyncZfb          func(arg string) error
	DoSyncWx           func(arg string) error

	PayCallBackService *service.PayCallBackService `inject:"PayCallBackService"`
	StoreService       *service.StoreService       `inject:"StoreService"`
	MallOrderService   *service.MallOrderService   `inject:"MallOrderService"`
}

func (it *PayServiceImpl) Init() {
	//从消息队列同步总方法
	it.SyncOrderByQueue = func() error {
		var e error
		//支付宝队列
		zfb_result, e := it.StoreService.ListLPop(zfb_notify)
		zfb_error := it.DoSyncZfb(core_util.CutJsonString(zfb_result))
		if zfb_error != nil {
			//出问题，left push 回队列
			it.StoreService.ListLPush(service.RedisKVDTO{
				Key:   zfb_notify,
				Value: zfb_result,
			})
			e = errors.New(zfb_error.Error())
		}
		//微信队列
		wx_result, e := it.StoreService.ListLPop(wx_notify)
		wx_error := it.DoSyncWx(core_util.CutJsonString(wx_result))
		if wx_error != nil {
			//出问题，left push 回队列
			it.StoreService.ListLPush(service.RedisKVDTO{
				Key:   wx_notify,
				Value: zfb_result,
			})
			if e != nil {
				e = errors.New(e.Error() + wx_error.Error())
			} else {
				e = errors.New(wx_error.Error())
			}
		}
		return e
	}

	//执行支付宝通知消息
	it.DoSyncZfb = func(result string) error {
		if result == "" {
			return nil
		}
		log.Println("同步支付宝订单:开始:" + result)
		var zfb vo.ZfbNotifyVO
		e := json.Unmarshal([]byte(result), &zfb)
		if e != nil {
			//出问题，left push 回队列
			log.Println(e)
			return e
		}
		//创建抽象订单对象 支付宝
		var zfb_notifyVO = vo.NewAbstractNotifyVOZfb(zfb)
		e = it.DoNotify(result, zfb_notifyVO)
		if e != nil {
			return e
		}
		return e
	}

	//执行微信通知消息
	it.DoSyncWx = func(result string) error {
		if result == "" {
			return nil
		}
		log.Println("同步微信订单:开始:" + result)
		var wx vo.WXNotifyVO
		e := json.Unmarshal([]byte(result), &wx)
		if e != nil {
			log.Println(e)
			return e
		}
		//创建抽象订单对象 微信
		var wx_notifyVO = vo.NewAbstractNotifyVOWx(wx)
		return it.DoNotify(result, wx_notifyVO)
	}

	it.DoNotify = func(result string, notifyVO vo.AbstractNotifyVO) error {
		//记录回调记录到数据库
		it.PayCallBackService.Add(vo.PayCallBackVO{
			PayCallBack: model.PayCallBack{
				Id:         utils.CreateUUID(),
				Data:       result,
				CreateTime: time.Now(),
			},
		})
		if notifyVO.OutTradeNo == "" {
			return nil
		}
		//同步抽象订单状态
		if strings.Index(notifyVO.OutTradeNo, "mall-") == 0 {
			//Sync mall order
			return it.MallOrderService.SyncOrderByQueue(notifyVO)
		} else {
			//todo other
			return nil
		}
	}
	GoMybatis.AopProxyService(&it.PayService, core_context.Engine)
	core_util.ScanInject("PayServiceImpl", it)
}
