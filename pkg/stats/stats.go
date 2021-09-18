package stats

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type AccessEvent struct {
	resource string
	httpverb string
	date     time.Time
	userinfo string
}

func accessiLstner() {

}

//AccessiLstner middleware
func AccessiLstner() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Caught request")
		//todo here handle request for Stats
		c.Next()
		time.Sleep(time.Nanosecond * 3000)
		fmt.Print("post next")
	}
}
