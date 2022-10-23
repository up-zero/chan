package define

type ChanProcess struct {
	Pid  int    `json:"pid"`  // 程序ID
	Conf string `json:"conf"` // 配置文件路径
}
