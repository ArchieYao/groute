package main

import (
	"archieyao.github.com/web-router/src/base"
	"net/http"
	"testing"
)

// author: ArchieYao
// date: 2021/7/15 4:33 下午
// description:

func TestRun(t *testing.T) {
	engine := base.New()
	engine.GET("/", func(c *base.Context) {
		c.HTML(http.StatusOK, "<h1> Hello World </h1>")
	})
	engine.GET("/hello", func(c *base.Context) {
		c.String(http.StatusOK, "hello %s , you're at %s ", c.Query("name"), c.Path)
	})
	engine.GET("/hello/:name", func(c *base.Context) {
		c.String(http.StatusOK, "hello %s ,you're at %s", c.Param("name"), c.Path)
	})
	engine.GET("/assets/*filepath", func(c *base.Context) {
		c.JSON(http.StatusOK, base.H{"filepath": c.Param("filepath")})
	})
	engine.Run(":8080")
}
