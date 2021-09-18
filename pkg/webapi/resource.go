package webapi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getFizzBuzz(c *gin.Context) {
	fmt.Println("HERE we are")
	r := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,..."
	c.IndentedJSON(http.StatusOK, r)

}
