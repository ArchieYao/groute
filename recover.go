package groute

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// ArchieYao created at 2022/5/26 17:35

func Recovery() HandleFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("%s\n\n", trace(fmt.Sprintf("%s", err)))
				c.FailWithCustomCode(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}

func trace(msg string) string {
	var uintptrs [32]uintptr
	n := runtime.Callers(2, uintptrs[:])

	var strBuilder strings.Builder
	strBuilder.WriteString(msg + "\n Traceback:")
	for _, pc := range uintptrs[:n] {
		funcForPC := runtime.FuncForPC(pc)
		file, line := funcForPC.FileLine(pc)
		strBuilder.WriteString(fmt.Sprintf("\n\t %s:%d", file, line))
	}
	return strBuilder.String()
}
