package base

import (
	"fmt"
	"reflect"
	"testing"
)

// author: ArchieYao
// date: 2021/7/15 4:11 下午
// description:

func TestRouter(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	fmt.Println(ok)
	ok = reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	fmt.Println(ok)
	reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	fmt.Println(ok)
}

func TestGetRouter(t *testing.T) {
	testRouter := newTestRouter()
	route, m := testRouter.getRoute(GET, "/hello/geek")
	if route == nil {
		t.Fatalf("ERROR")
	}
	fmt.Println(route.pattern)
	fmt.Println(m["name"])
}

func newTestRouter() *router {
	r := newRouter()
	r.addRouter(GET, "/", nil)
	r.addRouter(GET, "/hello/:name", nil)
	r.addRouter(GET, "/hello/b/c", nil)
	r.addRouter(GET, "/hi/:name", nil)
	r.addRouter(GET, "/assert/*filepath", nil)
	return r
}
