package main

import (
	"gitee.com/up-zero/chan/service"
	"gitee.com/up-zero/chan/util"
	"net/http"
)

func main() {
	service.NewService()
	util.Info("[SERVER] RUNNING ...")
	http.HandleFunc("/", service.HttpHandle)

	// 监听配置的所有端口
	service.ListenAndServer()
}
