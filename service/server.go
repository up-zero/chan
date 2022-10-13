package service

import (
	"gitee.com/up-zero/chan/util"
	"net/http"
)

func ListenAndServer() {
	ports := make(map[string]struct{})
	for _, v := range Srv.ChanConf.Server {
		vv := v
		if vv.Listen == ":443" {
			go func() {
				err := http.ListenAndServeTLS(vv.ServerName+vv.Listen, vv.SslCert, vv.SslKey, nil)
				if err != nil {
					panic("[LISTEN ERROR] : " + err.Error())
				}
			}()
		} else {
			if _, ok := ports[vv.Listen]; !ok {
				go func() {
					err := http.ListenAndServe(vv.ServerName+vv.Listen, nil)
					if err != nil {
						util.Error("[LISTEN ERROR] : " + err.Error())
					}
				}()
			}
		}

		ports[v.Listen] = struct{}{}
	}
	select {}
}
