package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Index"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":9090", router))
}

/*
HttpRouter 是一个高性能、可拓展的第三方HTTP路由包。
首先使用httprouter.New()函数生成一个*Router路由对象，然后使用GET()方法注册一个适配器，最后将*Router给ListenAndServer()函数即可


*/
