package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//go:generate mockgen -source=db.go -destination=./test/db_mock.go -package=db

var (
	// ErrNoDocuments record not found error
	ErrNoDocuments error = mongo.ErrNoDocuments
	// ErrAbsolutePath error raised when a migrationDir is absolute
	ErrAbsolutePath = errors.New("path should not be absolute")
	// DefaultTimeout is the default timeout for any interaction with the Mongo DB
	DefaultTimeout = 5 * time.Second
)

// Config holds database configuration
// URI examples:
//      "mongodb+srv://user:password@host:port/database?retryWrites=true&w=majority"
//      "mongodb://user:password@host:port"
type Config struct {
	URI string
	DB  string

	Timeout time.Duration
}

// Client represents the client to interact with a Mongo DB
type Client interface {
	Health(context.Context) error
	Disconnect(context.Context) error
	Collection(name string, opts ...*options.CollectionOptions) Collection
}

// Collection is the interface to interact with a mongo collection
type Collection interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) Result
}

// Result is the interface to interact with Query results
type Result interface {
	Decode(v interface{}) error
}

type mongoClient struct {
	client  *mongo.Client
	db      *mongo.Database
	timeout time.Duration
}

type mongoCollection struct {
	c *mongo.Collection
}

// NewClient creates a new database connection
func NewClient(ctx context.Context, conf Config) (Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.URI))
	if err != nil {
		return nil, fmt.Errorf("Failed to connect: %s", err)
	}
	timeout := conf.Timeout
	if timeout.Nanoseconds() == 0 {
		timeout = DefaultTimeout
	}
	mongoClient := &mongoClient{
		client:  client,
		db:      client.Database(conf.DB),
		timeout: timeout,
	}

	return mongoClient, nil
}

func (m mongoClient) Health(ctx context.Context) error {
	subCtx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()
	return m.client.Ping(subCtx, readpref.Primary())
}

func (m mongoClient) Disconnect(ctx context.Context) error {
	subCtx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()
	return m.client.Disconnect(subCtx)
}

func (m mongoClient) Collection(name string, opts ...*options.CollectionOptions) Collection {
	return mongoCollection{c: m.db.Collection(name, opts...)}
}

func (m mongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return m.c.InsertOne(ctx, document, opts...)
}

func (m mongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) Result {
	return m.c.FindOne(ctx, filter, opts...)
}
