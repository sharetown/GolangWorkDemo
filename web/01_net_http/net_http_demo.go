//package main
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//)

//func sayHelloName(w http.ResponseWriter,r *http.Request) {
//    r.ParseForm()
//    fmt.Println(r.Form)
//    fmt.Println("path",r.URL.Path)
//    fmt.Println("scheme",r.URL.Scheme)
//    fmt.Println(r.Form["url_long"])
//    for k,v := range r.Form{
//        fmt.Println("key:",k)
//        fmt.Println("val:",strings.Join(v,""))
//    }
//    fmt.Fprintf(w,"Hello astaxie!")
//
//}
//
//func main() {
//    http.HandleFunc("/",sayHelloName)
//    err := http.ListenAndServe(":9090",nil)
//    if err != nil{
//        log.Fatal("ListenAndServe: ", err)
//    }
//}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello") //这个写入到 w 的是输出到客户端的
	w.Write([]byte(" World！"))
}

func main() {
	http.HandleFunc("/", sayHello)           //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
