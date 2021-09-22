package webapi

import (
	"github.com/gin-gonic/gin"
	"github.com/momzor/fizzbuzz/pkg/db"
)

//go:generate mockgen -source=api_handler.go -destination=./api_handler_mock.go -package=webapi

// Handler interface for Api server
type APIHandler interface {
	StatsHandler() gin.HandlerFunc
}

type apiHandler struct {
	db db.Client
}

// NewAPIHandler create an api handler with injected db client
func NewAPIHandler(db db.Client) APIHandler {
	return &apiHandler{
		db: db,
	}
}
