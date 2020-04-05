package app_context

type AppConfig struct {
	PlatformName  string `json:"platform_name"` //平台中文名称
	SwaggerEnable bool
	ClientName    string `json:"client_name"`
	Host          string `json:"host"`
	Port          int    `json:"port"`
	ConsulUrl     string `json:"consul_url"`
	FastDFSUrl    string `json:"fast_dfs_url"`
}
