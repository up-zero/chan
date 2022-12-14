package service

import (
	"crypto/tls"
	"encoding/json"
	"gitee.com/up-zero/chan/define"
	"io/ioutil"
)

type Service struct {
	ChanConf     *define.ChanConf
	Server       map[string]define.Server
	Certificates []tls.Certificate
	TlsConfig    *tls.Config
}

var Srv *Service

func NewService(confPath string) {
	srv := new(Service)
	server := make(map[string]define.Server)
	conf := new(define.ChanConf)
	// conf
	confByte, err := json.Marshal(define.DefaultChanConf)
	if err != nil {
		panic("[ERROR MARSHAL] : " + err.Error())
	}
	if confPath != "" {
		confByte, err = ioutil.ReadFile(confPath)
	}
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
		if v.Listen == ":443" {
			cert, err := tls.LoadX509KeyPair(v.SslCert, v.SslKey)
			if err != nil {
				panic(err)
			}
			srv.Certificates = append(srv.Certificates, cert)
		}
		// upstream
		upstream := v.Upstream
		if upstream != nil {
			weights := 0
			for i, value := range upstream.Values {
				if value.Weight == 0 {
					upstream.Values[i].Weight = define.DefaultWeight
				}
				weights += upstream.Values[i].Weight
			}
			upstream.Weights = weights
		}
	}
	if len(srv.Certificates) > 0 {
		srv.TlsConfig = new(tls.Config)
		srv.TlsConfig.Certificates = srv.Certificates
	}
	srv.ChanConf = conf
	srv.Server = server
	Srv = srv
}
