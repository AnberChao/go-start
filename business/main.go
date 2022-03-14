package main

import (
	"net/http"

	"go-start/framework"
)

func main() {
	server := &http.Server{
		//自定义的请求核心处理函数
		Handler: NewCore(),

		Addr: ":8080",
	}
	server.ListenAndServe()
}
