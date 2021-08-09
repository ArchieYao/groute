package groute

import (
	"fmt"
	"log"
	"time"
)

// ArchieYao created at 2022/5/26 17:55

func ReqTimeCostLog() HandleFunc {
	return func(c *Context) {
		now := time.Now()
		c.Next()
		log.Println(fmt.Sprintf("[%d] [%s] execute in %v", c.StatusCode, c.Req.RequestURI, time.Since(now)))
	}
}
