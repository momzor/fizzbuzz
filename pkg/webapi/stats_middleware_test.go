package webapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	db "github.com/momzor/fizzbuzz/pkg/db"
	"github.com/stretchr/testify/assert"
)

//
func Test_StatsMiddleware(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDb := db.NewMockClient(mockCtrl)
	mockAm := NewMockAPIMiddleware(mockCtrl)
	mockH := NewMockAPIHandler(mockCtrl)
	mServConf := Config{
		Env:     gin.TestMode,
		BaseUrl: "localhost",
		Port:    "80",
	}
	mServer := Server{
		Conf:       mServConf,
		DBClient:   mockDb,
		Middleware: mockAm,
		Handler:    mockH,
	}
	t.Run("It should call stat middleware when starting the server", func(t *testing.T) {
		mockAm.EXPECT().StatsMiddleware().Times(1)
		mockH.EXPECT().StatsHandler().Times(1)

		w := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(w)
		r, err := http.NewRequest(http.MethodGet, FIZZBUZZ_RESOURCE_ROUTE, nil)
		fmt.Println("DEBUG12: ", err)
		mServer.InitServer()
		fmt.Println(mServer.Router)
		engine.ServeHTTP(w, r)

	})
}

// Table tests for string mapper
func Test_MapResourceFromPath(t *testing.T) {
	t.Run("It should return a well mapped resource name from path ", func(t *testing.T) {
		tests := map[string]struct {
			path   string
			expRes string
		}{
			FIZZBUZZ_RESOURCE_ROUTE: {path: FIZZBUZZ_RESOURCE_ROUTE, expRes: FIZZBUZZ_RESOURCE_NAME},
			STATS_RESOURCE_ROUTE:    {path: STATS_RESOURCE_ROUTE, expRes: STATS_RESOURCE_NAME},
			"Invalid":               {path: "MockMockMockingOnHeavensDoor", expRes: "unknown"},
		}

		for tName, tCase := range tests {

			t.Run(fmt.Sprintf("Case  : %s", tName), func(t *testing.T) {
				res := mapResourceFromPath(tCase.path)

				assert.Equal(t, res, tCase.expRes)
			})
		}
	})
}
