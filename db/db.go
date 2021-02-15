// Package db implements an abstraction layer for MongoDB connection.
package db

import (
	"context"
	"time"

	"github.com/Smart-Pot/pkg"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ConnectionTimeout time.Duration = 10 * time.Second
)

var (
	_collection *mongo.Collection
	_instance *mongo.Database
	_connected bool
)

// Connect connects to a mongo db database
func Connect(uri,name,collection string) error {
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(),ConnectionTimeout)	
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return err
	}
	_instance = client.Database(name)
	_collection = _instance.Collection(collection)
	_connected = true
	return nil
}



// Collection returns default collection,
// if db is not connected panics
// if collection is not defined returns nil
func Collection() *mongo.Collection {
	if !_connected {
		panic("db is not connected")
	}
	return _collection
}


// IsConnected returns database connection is established or not
func IsConnected() bool {
	return _connected
}


// PkgConfig returns information from pkg config
func PkgConfig(collection string) (string,string,string) {
	return pkg.Config.Database.Addr,pkg.Config.Database.DBName, collection
}