/**
  @File：glog.go
  @Author：TT
  @Time：2021/10/4 17:30
*/
package main

import (
	"flag"
	"github.com/golang/glog"
	"io"
	"log"
	"net/http"
)

func main() {

	flag.Set("v","4")
	glog.V(2).Info("Starting http server...")
	mux := http.ServeMux{}
	mux.HandleFunc("/",rootHangler)
	err := http.ListenAndServe(":9998",&mux)
	if err!=nil {
		log.Fatal(err)
	}

}

func rootHangler(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"okkk!")
}