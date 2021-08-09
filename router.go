package groute

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// author: ArchieYao
// date: 2021/7/8 9:12 下午
// description:

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type router struct {
	roots        map[string]*node
	handleFunMap map[string]HandleFunc
}

type HandleFunc func(c *Context)

// newRouter 添加一个新路由
func newRouter() *router {
	return &router{
		roots:        make(map[string]*node),
		handleFunMap: make(map[string]HandleFunc),
	}
}

// parsePattern 解析uri
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// addRouter 添加一个路由
func (r *router) addRouter(method string, pattern string, handler HandleFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handleFunMap[key] = handler
}

// getRoute 获取路由
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

// handle 执行每个路由的方法
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		key := c.Method + "-" + n.pattern
		c.Params = params
		c.handlers = append(c.handlers, r.handleFunMap[key])
		//r.handleFunMap[key](c)
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s \n", c.Path)
		})
		log.Println(fmt.Sprintf("404 ERROR, uri [%s] not found.", c.Req.RequestURI))
	}
	c.Next()
}
