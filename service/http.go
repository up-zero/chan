package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func HttpHandle(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	fmt.Println("RUN")
	if !strings.Contains(host, ":") {
		if _, ok := Srv.Server[host+":80"]; ok {
			host = host + ":80"
		} else if _, ok := Srv.Server[host+":443"]; ok {
			host = host + ":443"
		} else {
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte("502 Bad Gateway"))
			return
		}
	}

	for _, v := range Srv.Server[host].Location {
		if len(r.RequestURI) >= len(v.Path) && v.Path == r.RequestURI[:len(v.Path)] {
			if v.ProxyPass != "" { // 代理服务
				backendURL, err := url.Parse(v.ProxyPass)
				if err != nil {
					w.WriteHeader(http.StatusBadGateway)
					w.Write([]byte("502 Bad Gateway"))
					return
				}
				proxyHandler := httputil.NewSingleHostReverseProxy(backendURL)
				proxyHandler.ServeHTTP(w, r)
				return
			} else { // 静态服务
				fileDir := v.Root
				fmt.Println(r.RequestURI)
				filePath := fileDir + string(os.PathSeparator) + r.RequestURI[len(v.Path):]
				// index
				if r.RequestURI == v.Path {
					for _, indexName := range v.Index {
						nowFilePath := filePath + indexName
						data, err := ioutil.ReadFile(nowFilePath)
						if err != nil {
							if errors.Is(err, os.ErrNotExist) {
								continue
							}
						}
						w.Write(data)
						return
					}
				}
				index := strings.Index(filePath, "?")
				if index > 0 {
					filePath = filePath[:index]
				}
				fmt.Println(filePath)
				data, err := ioutil.ReadFile(filePath)
				if err != nil {
					w.WriteHeader(http.StatusNotFound)
					w.Write([]byte("404 Not Found"))
					return
				}

				w.Write(data)
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}
