package store

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBConfig contains the database configuration
type DBConfig struct {
	Host string `mapstructure:"host" default:"localhost"`
	Port string `mapstructure:"port" default:"27017"`
	Name string `mapstructure:"name" default:"entropy"`
}

// New returns the database instance
func New(config *DBConfig) (*mongo.Database, error) {
	uri := fmt.Sprintf(
		"mongodb://%s:%s/%s",
		config.Host,
		config.Port,
		config.Name,
	)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	pingCtx, pingCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer pingCancel()

	err = client.Ping(pingCtx, nil)

	return client.Database(config.Name), err
}
