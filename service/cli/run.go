package cli

import (
	"gitee.com/up-zero/chan/service"
	"gitee.com/up-zero/chan/util"
	"net/http"
	"strings"
)

func Run(args []string) {
	var confPath string
	if len(args) >= 4 && strings.ToUpper(args[2]) == "-T" {
		confPath = args[3]
	}
	service.NewService(confPath)
	util.Info("[SERVER] RUNNING ...")
	http.HandleFunc("/", service.HttpHandle)
	// 保存 ProcessID
	util.SaveProcessId()
	// 监听配置的所有端口
	service.ListenAndServer()
}
