package groute

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// author: ArchieYao
// date: 2021/7/8 9:16 下午
// description:

type H map[string]interface{}

type Context struct {
	Writer       http.ResponseWriter
	Req          *http.Request
	Path         string
	Method       string
	Params       map[string]string
	StatusCode   int
	handlers     []HandleFunc // plugins
	pluginIdx    int          // plugin index
	handleEngine *HandleEngine
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
		//Params:    map[string]string{},
		pluginIdx: -1,
	}
}

func (c *Context) Next() {
	c.pluginIdx++
	s := len(c.handlers)
	for ; c.pluginIdx < s; c.pluginIdx++ {
		c.handlers[c.pluginIdx](c)
	}
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 响应text请求
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) FailString(format string, values ...interface{}) {
	c.String(http.StatusForbidden, format, values...)
}

func (c *Context) FailWithCustomCode(code int, obj interface{}) {
	c.JSON(code, obj)
}
