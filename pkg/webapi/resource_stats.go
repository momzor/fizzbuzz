package webapi

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	STATS_RESOURCE_ROUTE     = "/stats"
	STATS_RESOURCE_NAME      = "stats"
	STATS_DB_COLLECTION_NAME = "access_events"
)

type Stats struct {
	Hits       int                 `bson:"count"`
	Parameters map[string][]string `bson:"parameters"`
	Resource   string              `bson:"resource"`
}

// DB Agregate for retreiving most used params for fizzbuzz resource
func (s *Server) StatsHandler(c *gin.Context) {
	pip := []bson.M{
		{"$match": bson.M{"method": http.MethodGet, "resource": FIZZBUZZ_RESOURCE_NAME}},
		{"$addFields": bson.M{"sortedParameters": bson.M{
			"int1":  "$parameters.int1",
			"int2":  "$parameters.int2",
			"str1":  "$parameters.str1",
			"str2":  "$parameters.str2",
			"limit": "$parameters.limit",
		}}},
		{"$group": bson.M{
			"_id": bson.M{
				"grpParameters": "$sortedParameters",
			},
			"parameters": bson.M{"$first": "$sortedParameters"},
			"resource":   bson.M{"$first": "$resource"},
			"count":      bson.M{"$sum": 1}}},
		{"$sort": bson.M{"count": -1}},
		{"$limit": 1}}

	csr, err := s.DBClient.Collection(STATS_DB_COLLECTION_NAME).Aggregate(context.Background(), pip)
	if err != nil {
		log.Fatal(err)
	}

	var res []Stats
	if err = csr.All(c, &res); err != nil {
		panic(err)
	}
	if len(res) > 0 {
		c.IndentedJSON(http.StatusOK, res[0])
	} else {
		c.IndentedJSON(http.StatusOK, Stats{})
	}
}
