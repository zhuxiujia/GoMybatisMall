package schedule

import (
	"github.com/facebookgo/inject"
	"github.com/robfig/cron"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/utils"
	"log"
	"strings"
)

//定时任务 调度 通知服务处理器
type OrderNotifyZfbSchedule struct {
	PayService *service.PayService `inject:"PayService"`
	DoNotify   func()              `cron:"0/10 * * * * ?"` //10秒执行一次，具体百度参考 cron表达式
}

func (it *OrderNotifyZfbSchedule) Init(cron *cron.Cron) {
	it.DoNotify = func() {
		var e = it.PayService.SyncOrderByQueue()
		if e != nil && !strings.Contains(e.Error(), "nil returned") {
			log.Println("OrderNotifyZfbSchedule同步支付宝结果失败:", e)
		}
	}
	//scan
	utils.CronScan(it, cron)
	app_context.Context.Provide(&inject.Object{
		Name:  "ZfbNotifyQueueSchedule",
		Value: it,
	})
}
