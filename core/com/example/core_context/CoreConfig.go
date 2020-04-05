package core_context

type CoreConfig struct {
	Server  string `json:"server"`  //服务名称
	Address string `json:"address"` //ip或者域名地址
	Consul  string `json:"consul"`  //consul服务发现地址
	Port    int    `json:"port"`    //服务端口

	Mysql          string `json:"mysql"`
	Redis_url      string `json:"redis_url"`
	Redis_password string `json:"redis_password"`
	CashierUrl     string `json:"cashier_url"` //支付宝和微信支付 http 请求接口（需要按照微信官方实现接口）
}
