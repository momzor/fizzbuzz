package webapi

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
)

type InputFilter struct {
	Int1  int    `schema:"int1,required,numeric"`
	Int2  int    `schema:"int2,required,numeric"`
	Limit int    `schema:"limit,required,number"`
	Str1  string `schema:"str1,required"`
	Str2  string `schema:"str2,required"`
}

type ApiResponse struct {
	Status  int
	Errors  []string `json:"errors"`
	Message string   `json:"message"`
}

func getFizzBuzz(c *gin.Context) {
	//Query Validation
	var f InputFilter

	if err := schema.NewDecoder().Decode(&f, c.Request.URL.Query()); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ApiResponse{
			Status:  http.StatusBadRequest,
			Errors:  []string{err.Error()},
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, buildFizzBuzz(f))
}

func buildFizzBuzz(f InputFilter) (fb string) {
	var r []string
	bothMultiplier := f.Int1 * f.Int2

	for i := 1; i < f.Limit; i++ {
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
