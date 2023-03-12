package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "indexHandler......")
}
func hiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hiHandler......")
}
func webHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "webHandler......")
}
func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testHandler......")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/hi", hiHandler)
	//如果我访问的地址变成"localhost:9090/hi/"的话将无法匹配到"/hi"这个路由，
	//因为"/hi"这个不以"/"结尾的路由需要精确匹配，
	//反之如果哪个路由是以"/"结尾的话，则访问的路径中只要前缀与之相同的话即表示匹配上了
	mux.HandleFunc("/test/", testHandler)
	//这里注册路由时以"/"结尾了，则只要地址的前缀与之匹配上即可匹配。例如"http://localhost:9090/test/abcdefg/hijklmn"就可以

	mux.HandleFunc("/hi/web", webHandler)

	server := &http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
