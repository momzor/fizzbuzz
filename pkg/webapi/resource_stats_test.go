package webapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	"github.com/momzor/fizzbuzz/pkg/db"
)

//Test_StatsHandler handle
func Test_StatsHandler(t *testing.T) {
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
	t.Run("It should call Stats handler if route match  ", func(t *testing.T) {
		mockAm.EXPECT().StatsMiddleware().Times(1)
		mockH.EXPECT().StatsHandler().Times(1)
		w := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(w)
		r, _ := http.NewRequest(http.MethodGet, STATS_RESOURCE_ROUTE, nil)
		mServer.InitServer()
		fmt.Println(mServer.Router)
		engine.ServeHTTP(w, r)

	})

}
