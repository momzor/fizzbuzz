package webapi

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// AccessEvent represents a db document of the
// collection "access_event" holding all api access events
type AccessEvent struct {
	Resource   string              `bson:"resource"`
	Method     string              `bson:"method"`
	Date       time.Time           `bson:"date"`
	Uri        string              `bson:"uri"`
	Parameters map[string][]string `bson:"parameters"`
}

//AccessiLstner middleware
func (s *Server) statsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// call c.Next() in order to prevent blocking user if db insert fails
		c.Next()

		var aE = AccessEvent{
			Parameters: c.Request.URL.Query(),
			Resource:   mapResourceFromPath(c.Request.URL.Path),
			Method:     c.Request.Method,
			Date:       time.Now(),
		}

		_, err := s.DBClient.Collection(STATS_DB_COLLECTION_NAME).InsertOne(context.Background(), &aE)
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to save BSON to Mongo DB: %s", err))
		}
	}
}

// resource mapper
func mapResourceFromPath(p string) string {
	switch p {
	case FIZZBUZZ_RESOURCE_ROUTE:
		return FIZZBUZZ_RESOURCE_NAME
	case STATS_RESOURCE_ROUTE:
		return STATS_RESOURCE_NAME
	default:
		return "unknown"
	}
}