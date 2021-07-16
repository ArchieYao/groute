package main

import (
	"archieyao.github.com/web-router/src/base"
	"log"
	"net/http"
)

// author: ArchieYao
// date: 2021/7/8 8:17 下午
// description:

func main() {
	log.Println("start web-router.")

	handleEngine := base.New()

	handleEngine.GET("/", func(c *base.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	handleEngine.GET("/hello", func(c *base.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s .", c.Query("name"), c.Path)
	})

	handleEngine.POST("/login", func(c *base.Context) {
		c.JSON(http.StatusOK, base.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	handleEngine.Run(":8080")
}
