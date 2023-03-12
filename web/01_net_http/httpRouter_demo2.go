package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/default", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("default get"))
	})
	router.POST("/default", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("default get"))
	})
	//精确匹配
	//router.GET("/user/name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//	w.Write([]byte("user name:" + p.ByName("name")))
	//})
	//匹配所有
	router.GET("/user/*name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("user name:" + p.ByName("name")))
	})
	http.ListenAndServe(":8080", router)
}
