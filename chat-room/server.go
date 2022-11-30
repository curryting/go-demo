package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// 注册路由
	router := mux.NewRouter()
	go h.run()

	// 绑定路由方法
	router.HandleFunc("/ws", myws)

	// 监听端口
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
