package base

import (
	"net/http"
)

// author: ArchieYao
// date: 2021/7/8 8:21 下午
// description:

type HandleEngine struct {
	router *router
}

func (handleEngine *HandleEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	handleEngine.router.handle(c)
}

const GET = "GET"
const POST = "POST"
const PUT = "PUT"
const DELETE = "DELETE"

// New 初始化
func New() *HandleEngine {
	return &HandleEngine{
		router: newRouter(),
	}
}

// Run 启动一个服务
func (handleEngine *HandleEngine) Run(addr string) error {
	return http.ListenAndServe(addr, handleEngine)
}

// addRouter 增加一个路由
func (handleEngine *HandleEngine) addRouter(method string, uri string, handler HandleFunc) {
	handleEngine.router.addRouter(method, uri, handler)
}

// GET 增加一个GET方法
func (handleEngine *HandleEngine) GET(uri string, handleFunc HandleFunc) {
	handleEngine.addRouter(GET, uri, handleFunc)
}

// POST 增加一个POST方法
func (handleEngine *HandleEngine) POST(uri string, handleFunc HandleFunc) {
	handleEngine.addRouter(POST, uri, handleFunc)
}

// DELETE 增加一个DELETE方法
func (handleEngine HandleEngine) DELETE(uri string, handleFunc HandleFunc) {
	handleEngine.addRouter(DELETE, uri, handleFunc)
}

// PUT 增加一个PUT方法
func (handleEngine HandleEngine) PUT(uri string, handleFunc HandleFunc) {
	handleEngine.addRouter(PUT, uri, handleFunc)
}
