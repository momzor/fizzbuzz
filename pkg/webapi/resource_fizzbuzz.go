package webapi

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
)

const (
	FIZZBUZZ_RESOURCE_ROUTE = "/fizzbuzz"
	FIZZBUZZ_RESOURCE_NAME  = "fizzbuzz"
)

// Fizzbuzz represents the resource payload
type Fizzbuzz struct {
	Int1  int    `schema:"int1,required"`
	Int2  int    `schema:"int2,required"`
	Limit int    `schema:"limit,required"`
	Str1  string `schema:"str1,required"`
	Str2  string `schema:"str2,required"`
}

// FizzBuzzHandler godoc
// @Summary Return Fizzbuzz resource
// @Description get custom Fizzbuzz
// @Produce json
// @Param int1 query int true "int1 query parameter"
// @Param int2 query int true "int2 query parameter"
// @Param limit query int true "limit query parameter"
// @Param str1 query string true "str1 query parameter"
// @Param str2 query string true "str2 query parameter"
// @Router /fizzbuzz [get]
func FizzBuzzHandler(c *gin.Context) {
	var f Fizzbuzz

	if err := schema.NewDecoder().Decode(&f, c.Request.URL.Query()); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{
			Status:  http.StatusBadRequest,
			Errors:  []string{"invalid parameters, please refer to the api documentation", err.Error()},
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, buildFizzBuzz(f))
}

/* BuildFizzBuzz builds a fizz buzz string rom a Fizzbuzz payload
S
*/
func buildFizzBuzz(f Fizzbuzz) (fb string) {
	var r []string
	bothMultiplier := f.Int1 * f.Int2

	for i := 1; i <= f.Limit; i++ {
		// no Zero division
		if f.Int1 <= 0 || f.Int2 <= 0 {
			r = append(r, strconv.Itoa(i))
			continue
		}

		if i%bothMultiplier == 0 {
			r = append(r, f.Str1+f.Str2)
			continue
		}
		if i%f.Int1 == 0 {
			r = append(r, f.Str1)
			continue
		}
		if i%f.Int2 == 0 {
			r = append(r, f.Str2)
			continue
		}

		r = append(r, strconv.Itoa(i))
	}

	return strings.Join(r, ",")

}
