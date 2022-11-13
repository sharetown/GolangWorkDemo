package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

//使用类型转换来校验输入的数是不是合法的数字
func inputForAtoI(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("/workspace/GolangWorkDemo/web/03_validate_input_of_form/input.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()

		getint, err := strconv.Atoi(r.Form.Get("number"))
		if err != nil {
			// 数字转化出错了，那么可能就不是数字
			fmt.Fprintf(w, "inputForAtoI:数字转化出错")
			log.Println("inputForAtoI:数字转化出错")
		}

		// 接下来就可以判断这个数字的大小范围了
		// if getint > 100 {
		// 	// 太大了
		// }
		fmt.Printf("%v Type of %T", getint, getint)
	}
}

//使用正则表达式来校验输入的数是不是合法的数字
func inputForMatch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("/workspace/GolangWorkDemo/web/03_validate_input_of_form/input.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()

		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("number")); !m {
			// 数字转化出错了，那么可能就不是数字
			fmt.Fprintf(w, "inputForMatch:数字转化出错")
			log.Println("inputForMatch:数字转化出错")
		}

		// 接下来就可以判断这个数字的大小范围了
		// if getint > 100 {
		// 	// 太大了
		// }
		fmt.Printf("%v Type of %T", r.Form.Get("number"), r.Form.Get("number"))
	}
}

func main() {
	http.HandleFunc("/", inputForAtoI)
	// http.HandleFunc("/input", inputForAtoI)
	http.HandleFunc("/input", inputForMatch)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

//更多校验表单输入合法性的示例可见：https://learnku.com/docs/build-web-application-with-golang/042-validation-form-input/3176
