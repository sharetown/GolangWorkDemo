package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func input(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("/workspace/GolangWorkDemo/web/04_disable_multiple_input/input.gtpl")
		fmt.Println("token:", token)
		t.Execute(w, token) //写入token
	} else {
		// 请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// 验证 token 的合法性
		} else {
			// 不存在 token 报错
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // 输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) // 输出到客户端
	}
}

func main() {
	http.HandleFunc("/", input)
	http.HandleFunc("/input", input)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
