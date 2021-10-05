// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/thinkeridea/go-extend/exnet"
)

func hello(w http.ResponseWriter, r *http.Request) {
	//接收客户端 request，并将 request 中带的 header 写入 response header
	for key, values := range r.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	v := "VERSION"
	version := os.Getenv(v)
	w.Header().Set(v, version)

	//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	statusCode := 200
	fmt.Println("IP:", ip, "http status code:", statusCode)
	w.WriteHeader(statusCode)
}

/**
  healthz 健康检查
*/
func healthZ(w http.ResponseWriter, r *http.Request) {
	//当访问 localhost/healthz 时，应返回200
	_, err := fmt.Fprintf(w, "ok")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/healthz", healthZ)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("httpServer listen:", err)
	}
}
