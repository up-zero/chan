package define

type ChanConf struct {
	Server []Server `json:"server"`
}

type Server struct {
	Listen     string `json:"listen"`      // 监听地址
	ServerName string `json:"server_name"` // 域名
	SslCert    string `json:"ssl_cert"`    // ssl 证书
	SslKey     string `json:"ssl_key"`     // ssl key
	Location   []struct {
		Path      string   `json:"path"`       // 父级路由地址
		Index     []string `json:"index"`      // 首页包含的地址
		Root      string   `json:"root"`       // 静态资源文件地址
		ProxyPass string   `json:"proxy_pass"` // 反向代理地址，与静态代理互斥，同时配置时只有一种生效
	} `json:"location"` // 请求地址
}
