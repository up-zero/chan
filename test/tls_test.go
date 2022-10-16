package test

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"
)

type handler struct{}

var Handler = handler{}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO WORLD !!"))
}

func TestRunMoreTlsServer(t *testing.T) {
	srv := http.Server{
		Addr:    ":443",
		Handler: Handler,
	}
	defer srv.Close()
	// tls config
	cert1, err := tls.LoadX509KeyPair("..\\conf\\ssl\\test.getcharzp.cn\\test.getcharzp.cn_bundle.pem",
		"..\\conf\\ssl\\test.getcharzp.cn\\test.getcharzp.cn.key")
	if err != nil {
		panic(err)
	}
	cert2, err := tls.LoadX509KeyPair("..\\conf\\ssl\\test2.getcharzp.cn\\test2.getcharzp.cn_bundle.pem",
		"..\\conf\\ssl\\test2.getcharzp.cn\\test2.getcharzp.cn.key")
	if err != nil {
		panic(err)
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{
			cert1, cert2,
		},
	}
	srv.TLSConfig = tlsConfig
	fmt.Println("RUN SERVER")
	err = srv.ListenAndServeTLS("", "")
	t.Fatal(err)
}
