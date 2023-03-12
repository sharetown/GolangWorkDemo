首先我们看一下我们main文件中的代码

day2-context/main.go
```go
func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
```
gee.New()返回我们的Engine类的指针，而Engine类通过重写ServeHTTP方法继承自Handler接口。http中的ListenAndServe方法的两个参数，第一个是端口号，第二个便是Handler类型的事件处理器，这里的Engine便是我们自己的Handler。

GET、POST都是添加路由的方法，第一个参数是一个string类型的，代表路由的地址。第二个参数是一个函数类型的参数，它里面可以用来处理一些事件的逻辑。它被定义在gee文件中，他的参数列表是一个类型为Context的指针，在这个函数类型作用域内，便可以调用Context结构体中的所有属性和方法。

Run方法用来启动一个web服务，它调用http.ListenAndServe，将参数中的端口号，以及r（Engine）这个处理器传入，启动了一个我们自定义的web服务

```shell
day2-context
│  main.go
└─gee
     context.go
     gee.go
     router.go
```

