package gee

import (
	"log"
	"net/http"
)

// HandlerFunc是我们定义的一个函数类型，他的形参是Context类型的指针，在这个函数类型里，便可以调用Context里的属性与方法
type HandlerFunc func(*Context)

// Engine重写ServeHTTP实现Handler接口，并且里面还有一个router类型的指针，router用来帮助我们处理路由，是从前一天的gee中封装剥离出去的
type Engine struct {
	router *router
}

// New是gee.Engine的构造函数
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRoute(method, pattern, handler)
}

//GET 定义添加 GET 请求的方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 定义添加 POST 请求的方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 定义了启动一个http服务器的方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// http.ListenAndServe底层会处理这个方法，里面是要实现的逻辑
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
