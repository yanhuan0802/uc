// main.go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/thinkeridea/go-extend/exnet"

	"github.com/yanhuan0802/uc/httpserver/metrics"
)

func hello(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	//接收客户端 request，并将 request 中带的 header 写入 response header
	for key, values := range r.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	//添加0-2s随机延迟
	delay := randInt(10, 2000)
	time.Sleep(time.Millisecond * time.Duration(delay))
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
	fmt.Println("IP:", ip, "http status code:", statusCode, "version", version)
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

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/healthz", healthZ)
	http.Handle("metrics", promhttp.Handler())

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("httpServer listen:", err)
	}
}
