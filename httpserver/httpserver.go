/**
  @File：httpserver.go
  @Author：TT
  @Time：2021/10/4 17:21

	1. 接收客户端 request，并将 request 中带的 header 写入 response header
	2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	4. 当访问 localhost/healthz 时，应返回200
*/
package main

import (
	"fmt"
	"github.com/golang/glog"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//自己创建多路复用器
	mux := http.NewServeMux()

	//http.HandleFunc("/",handler)
	mux.HandleFunc("/",indexHandler)
	mux.HandleFunc("/healthz",healthzHandler)
	//http.HandleFunc("/",indexHandler)
	//http.HandleFunc("/healthz",healthzHandler)

	server := &http.Server{
		Addr: "localhost:9999",
		Handler: mux,
		ReadTimeout: 3*time.Second,
	}

	err := server.ListenAndServe()
	if err!=nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter,r *http.Request){

	//要求一：接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header {
		w.Header().Set(k,fmt.Sprint(v))
	}

	//要求二：读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	//os.Getenv检索环境变量并返回值，如果变量是不存在的，这将是空的。
	version := os.Getenv("VERSION")
	if version=="" {
		w.Header().Set("VERSION","nil")
	}else {
		w.Header().Set("VERSION",version)
	}
	w.WriteHeader(http.StatusOK)

	//要求三：Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	glog.Info(r.RemoteAddr)//记录客户端ip
	glog.Info(http.StatusOK)//返回码？疑问？？？

}

func healthzHandler(w http.ResponseWriter,r *http.Request)  {

	//要求四：当访问 localhost/healthz 时，应返回200
	w.WriteHeader(http.StatusOK)
	glog.Info("访问healthz")//记录日志
}

