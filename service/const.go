package service

import (
	"encoding/json"
	"gitee.com/up-zero/chan/define"
	"io/ioutil"
)

type Service struct {
	ChanConf *define.ChanConf
	Server   map[string]define.Server
}

var Srv *Service

func NewService() {
	srv := new(Service)
	server := make(map[string]define.Server)
	conf := new(define.ChanConf)
	// conf
	confByte, err := ioutil.ReadFile("./conf/chan.json")
	if err != nil {
		panic("[ERROR READ CONF] : " + err.Error())
	}
	err = json.Unmarshal(confByte, conf)
	if err != nil {
		panic("[ERROR UNMARSHAL] : " + err.Error())
	}
	// server
	for _, v := range conf.Server {
		server[v.ServerName+v.Listen] = v
	}
	srv.ChanConf = conf
	srv.Server = server
	Srv = srv
}
