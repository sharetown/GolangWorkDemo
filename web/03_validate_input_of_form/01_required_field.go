//校验必填字段
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func input(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//以GET请求访问时加载模板input.gtpl
		t, _ := template.ParseFiles("/workspace/GolangWorkDemo/web/03_validate_input_of_form/input.gtpl")
		log.Println(t.Execute(w, nil)) //将模板渲染后传递给客户端
		//模板语言中声明了一个表单，其中<form action="/input" method="post">
		//表示：客户端提交时会发送POST请求到"/input"路由
	} else {
		//以POST请求访问时校验数据
		r.ParseForm() //解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		// 注意:如果没有调用 ParseForm 方法，下面无法获取表单的数据
        //默认情况下，Handler 里面是不会自动解析 form 的，必须显式的调用 r.ParseForm() 后，你才能对这个表单数据进行操作

        //校验必填字段
		if len(r.Form["username"][0]) == 0 {
			fmt.Fprintf(w, "username字段必填")
		}
	}

}

func main() {
	http.HandleFunc("/", input) //注册路由
	http.HandleFunc("/input", input)
	err := http.ListenAndServe(":9090", nil) //开启监听服务
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
