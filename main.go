package main

import (
	"gitee.com/up-zero/chan/service"
	"gitee.com/up-zero/chan/util"
	"net/http"
)

func main() {
	util.Info("[SERVER] RUNNING ...")
	http.HandleFunc("/", service.HttpHandle)
	http.ListenAndServe(":80", nil)
}
