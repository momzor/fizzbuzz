package webapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// Test_FizzBuzzHandler Run test units for webapi FizzBuzzhandler function
func Test_FizzBuzzHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	tests := map[string]struct {
		urlParams    string
		expectedCode int
	}{
		"valid":                   {urlParams: "?int1=3&int2=3&limit=10&str1=fizz&str2=buzzzz", expectedCode: http.StatusOK},
		"missing required str2":   {urlParams: "?int1=3&int2=3&limit=10&str1=fizz&str2=", expectedCode: http.StatusBadRequest},
		"missing required str1":   {urlParams: "?int1=3&int2=3&limit=10&&str2=buzzzz", expectedCode: http.StatusBadRequest},
		"missing required limit":  {urlParams: "?int1=3&int2=3&str1=fizz&str2=buzzzz", expectedCode: http.StatusBadRequest},
		"missing required int1":   {urlParams: "?&int2=3&limit=10&str1=fizz&str2=buzzzz", expectedCode: http.StatusBadRequest},
		"missing required int2":   {urlParams: "?int1=3&limit=10&str1=fizz&str2=buzzzz", expectedCode: http.StatusBadRequest},
		"invalid value for str1":  {urlParams: "?int1=3&int2=3&limit=10&str1=&str2=buzzzz", expectedCode: http.StatusBadRequest},
		"invalid value for str2":  {urlParams: "?int1=3&int2=3&limit=10&str1=fizz&str2=", expectedCode: http.StatusBadRequest},
		"invalid value for int1":  {urlParams: "?int1=INVALIDMOCK&int2=3&limit=10&str1=fizz&str2=buzzzz", expectedCode: http.StatusBadRequest},
		"invalid value for int2":  {urlParams: "?int1=3&int2=INVALIDMOCK&limit=10&str1=fizz&str2=buzzzz", expectedCode: http.StatusBadRequest},
		"invalid value for limit": {urlParams: "?int1=3&int2=3&limit=INVALIDMOCK&str1=fizz&str2=buzzzz", expectedCode: http.StatusBadRequest},
	}

	for tName, tCase := range tests {

		t.Run(fmt.Sprintf("It should put error in response if params are invalids : %s", tName), func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest(http.MethodGet, tCase.urlParams, nil)

			FizzBuzzHandler(c)

			assert.Equal(t, w.Code, tCase.expectedCode)
		})

	}

}

// Test_buildFizzBuzz Run test the builder which build the fizzbuzz string from parameters
func Test_buildFizzBuzz(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	//TODO : REPLACE this by fixture files for more tests cases
	tests := map[string]struct {
		f   Fizzbuzz
		res string
	}{
		"original- short": {f: Fizzbuzz{
			Int1:  3,
			Int2:  5,
			Limit: 30,
			Str1:  "fizz",
			Str2:  "buzz",
		}, res: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz"},
		"1&1- short": {f: Fizzbuzz{
			Int1:  1,
			Int2:  1,
			Limit: 20,
			Str1:  "fizz",
			Str2:  "buzz",
		}, res: "fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz"},
		"0&0": {f: Fizzbuzz{
			Int1:  0,
			Int2:  0,
			Limit: 10,
			Str1:  "fizz",
			Str2:  "buzz",
		}, res: "1,2,3,4,5,6,7,8,9,10"},
	}

	for tName, tCase := range tests {

		t.Run(fmt.Sprintf("It should put error in response if params are invalids : %s", tName), func(t *testing.T) {
			res := buildFizzBuzz(tCase.f)

			assert.Equal(t, res, tCase.res)
		})

	}

}
