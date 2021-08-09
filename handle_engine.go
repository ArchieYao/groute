package groute

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// author: ArchieYao
// date: 2021/7/8 8:21 下午
// description:

type (
	HandleEngine struct {
		*RouterGroup // 嵌套类型，HandleEngine 等同于拥有 RouterGroup 的全部属性
		router       *router
		groups       []*RouterGroup
	}

	RouterGroup struct {
		prefix  string       // 分组前缀
		plugins []HandleFunc // 插件
		parent  *RouterGroup
		engine  *HandleEngine // HandleEngine
	}
)

func (handleEngine *HandleEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var plugins []HandleFunc
	for _, group := range handleEngine.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) { // 通过前缀匹配
			plugins = append(plugins, group.plugins...)
		}
	}
	c := newContext(w, r)
	c.handlers = plugins
	c.handleEngine = handleEngine
	handleEngine.router.handle(c)
}

// New 初始化
func New() *HandleEngine {
	engine := &HandleEngine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	engine.Use(ReqTimeCostLog(), Recovery())
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) Use(handleFunc ...HandleFunc) {
	group.plugins = append(group.plugins, handleFunc...)
}

func (group *RouterGroup) addRouter(method string, comp string, handler HandleFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRouter(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handleFunc HandleFunc) {
	group.addRouter(GET, pattern, handleFunc)
}

func (group *RouterGroup) POST(pattern string, handleFunc HandleFunc) {
	group.addRouter(POST, pattern, handleFunc)
}

func (group *RouterGroup) PUT(pattern string, handleFunc HandleFunc) {
	group.addRouter(PUT, pattern, handleFunc)
}

func (group *RouterGroup) DELETE(patter string, handleFunc HandleFunc) {
	group.addRouter(DELETE, patter, handleFunc)
}

// Run 启动一个服务
func (handleEngine *HandleEngine) Run(addr string) error {
	err := http.ListenAndServe(addr, handleEngine)
	if err != nil {
		log.Fatalln(fmt.Sprintf("server run at [%s] failed", addr))
	}
	return err
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
