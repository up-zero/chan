package service

import (
	"net/http"
)

type handler struct{}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	HttpHandle(w, r)
}

func listenAndServeTLS() error {
	var handle = new(handler)
	srv := &http.Server{
		Addr:    ":443",
		Handler: handle,
	}
	defer srv.Close()
	srv.TLSConfig = Srv.TlsConfig
	err := srv.ListenAndServeTLS("", "")
	return err
}
