package service

type PayService struct {
	//支付服务同步订单方法，订单信息是保存于redis队列中，需要异步处理待处理的订单状态（例如 收到支付宝回调后，修改订单为支付成功）
	SyncOrderByQueue func() error
}
