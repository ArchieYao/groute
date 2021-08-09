package example

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/archieyao/groute"
)

// author: ArchieYao
// date: 2021/7/15 4:33 下午
// description:

func TestRun(t *testing.T) {
	engine := groute.New()
	engine.GET("/", func(c *groute.Context) {
		c.HTML(http.StatusOK, "<h1> Hello World </h1>")
	})
	engine.GET("/hello", func(c *groute.Context) {
		c.String(http.StatusOK, "hello %s , you're at %s ", c.Query("name"), c.Path)
	})
	engine.GET("/hello/:name", func(c *groute.Context) {
		c.String(http.StatusOK, "hello %s ,you're at %s", c.Param("name"), c.Path)
	})
	engine.GET("/assets/*filepath", func(c *groute.Context) {
		c.JSON(http.StatusOK, groute.H{"filepath": c.Param("filepath")})
	})
	engine.Run(":8080")
}

func TestGroupRun(t *testing.T) {
	engine := groute.New()
	engine.GET("/", func(c *groute.Context) {
		c.HTML(http.StatusOK, "<h1> Hello World </h1>")
	})

	v1 := engine.Group("/v1")
	{
		v1.GET("/hello", func(c *groute.Context) {
			c.String(http.StatusOK, "hello %s , you're at %s ", c.Query("name"), c.Path)
		})
		v1.GET("/hello/:name", func(c *groute.Context) {
			c.String(http.StatusOK, "hello %s ,you're at %s", c.Param("name"), c.Path)
		})
	}

	v2 := engine.Group("/v2")
	{
		v2.GET("/hello", func(c *groute.Context) {
			c.String(http.StatusOK, "hello %s , you're at %s ", c.Query("name"), c.Path)
		})
		v2.GET("/hello/:name", func(c *groute.Context) {
			c.String(http.StatusOK, "hello %s ,you're at %s", c.Param("name"), c.Path)
		})
	}
	engine.Run(":9090")
}

func TestPlugins(t *testing.T) {

	//toString := base64.StdEncoding.EncodeToString([]byte("我是中国人"))
	//fmt.Println(toString)
	//decodeString, err := base64.StdEncoding.DecodeString(toString)
	//if err == nil {
	//	fmt.Println(string(decodeString))
	//}

	engine := groute.New()
	//engine.Use(plugin1())
	engine.GET("/", func(c *groute.Context) {
		c.HTML(http.StatusOK, "<h1>hello world</h1>")
	})

	groupV2 := engine.Group("/v2")
	groupV2.Use(plugin1())
	groupV2.Use(plugin2())
	{
		groupV2.GET("/hello/:name", func(c *groute.Context) {
			c.String(http.StatusOK, "hello %s ,you're at %s", c.Param("name"), c.Path)
		})
	}

	groupV3 := engine.Group("/v3")
	groupV3.Use(plugin1())
	{
		groupV3.GET("/hello/:name", func(c *groute.Context) {
			c.String(http.StatusOK, "hello %s ,you're at %s", c.Param("name"), c.Path)
		})
	}

	err := engine.Run(":9090")
	if err != nil {
		fmt.Println(fmt.Sprintf("err %v", err))
	}
}

func plugin1() groute.HandleFunc {
	return func(c *groute.Context) {
		t := time.Now()
		fmt.Println(fmt.Sprintf("[%d] %s in %v for plugin1", c.StatusCode, c.Req.RequestURI, time.Since(t)))
		c.Next()
	}
}

func plugin2() groute.HandleFunc {
	return func(c *groute.Context) {
		fmt.Println("------------------")
		name := c.Param("name")
		if name == "aaa" {
			log.Println(fmt.Sprintf("name %s cannot call this api", name))
		}
		c.Next()
	}
}

func TestFileList(t *testing.T) {
	// 读取当前目录中的所有文件和子目录
	files, err := ioutil.ReadDir(`/Users/archieyao/GoProjects/web-router`)
	if err != nil {
		panic(err)
	}
	// 获取文件，并输出它们的名字
	for _, file := range files {
		println(file.Name())
	}
}

func TestWalkfile(t *testing.T) {
	var files []string
	// 定期扫描路径，与logFileSet对比，找到不在logFileSet中的文件，并上报
	err := filepath.Walk("/Users/archieyao/GoProjects/web-router", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {

	}
	for _, v := range files {
		fmt.Println(v)
	}
}

func TestGo(t *testing.T) {
	//time.Sleep(time.Minute * 1)

	fmt.Println("xxxxxx")
}

func TestRecovery(t *testing.T) {
	engine := groute.New()
	engine.GET("/panic", func(c *groute.Context) {
		array := []int{1, 1, 2, 3}
		fmt.Println(array[10])
		c.String(http.StatusOK, "************")
	})
	err := engine.Run(":9090")
	if err != nil {
		log.Println(fmt.Sprintf("%s", err))
	}
}
