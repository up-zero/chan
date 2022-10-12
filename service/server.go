package service

import "net/http"

func ListenAndServer() {
	ports := make(map[string]struct{})
	for _, v := range Srv.ChanConf.Server {
		ports[v.Listen] = struct{}{}
		go http.ListenAndServe(v.ServerName+v.Listen, nil)
	}
	select {}
}
