package database

import (
	"sync"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB_JOB *mongo.Database
var DB *mongo.Database
var CACHE *redis.Client

// PubSub type definition
type (
	PubSub struct {
		Conn *redis.PubSub
		Quit chan struct{}
	}

	Map struct {
		sync.Mutex
		M map[string]*PubSub
	}
)

var Subscriptions = Map{
	M: make(map[string]*PubSub),
}

// Connect creates a connection to database
func Connect() {
	ConnectRabbitMQ()

}
