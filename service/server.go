package service

import (
	"gitee.com/up-zero/chan/util"
	"net/http"
)

func ListenAndServer() {
	ports := make(map[string]struct{})
	for _, v := range Srv.ChanConf.Server {
		vv := v
		if _, ok := ports[vv.Listen]; !ok {
			if vv.Listen == ":443" {
				go func() {
					err := listenAndServeTLS()
					if err != nil {
						util.Error("[LISTEN ERROR] : " + err.Error())
						panic("[LISTEN ERROR] : " + err.Error())
					}
				}()
			} else {
				go func() {
					err := http.ListenAndServe(vv.Listen, nil)
					if err != nil {
						util.Error("[LISTEN ERROR] : " + err.Error())
						panic("[LISTEN ERROR] : " + err.Error())
					}
				}()
			}
		}
		ports[v.Listen] = struct{}{}
	}
	select {}
}
