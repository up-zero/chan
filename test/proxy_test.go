package test

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"testing"
)

// 启动 8080 服务，用于被代理
func TestRunServer(t *testing.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(request.RequestURI))
	})
	fmt.Println("SERVER RUNNING 8080 ......")
	http.ListenAndServe(":8080", nil)
}

// 启动 8000 服务，用于程序的入口
func TestReverseProxy(t *testing.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 反向代理核心逻辑
		backendURL, err := url.Parse("http://localhost:8080")
		if err != nil {
			t.Fatal(err)
		}
		proxyHandler := httputil.NewSingleHostReverseProxy(backendURL)
		proxyHandler.ServeHTTP(writer, request)
	})
	fmt.Println("SERVER RUNNING 8000 ......")
	http.ListenAndServe(":8000", nil)
}
