package db_test

import (
	"context"
	"testing"
	"time"

	"github.com/Smart-Pot/pkg/db"
	"github.com/stretchr/testify/assert"
)


const (
	_connectionURI = "mongodb://localhost:27017"
	_collection = "todos"
	_dbName = "test"

)

func TestConnect(t *testing.T) {
	t.Run("Connect",func(t *testing.T){
		assert.Equal(t,db.IsConnected(),false)

		//err := db.Connect("mongodb://wrongURI:27017",_collection,_dbName)
		//assert.NotNil(t,err)

		err := db.Connect(_connectionURI,_collection,_dbName)
		assert.Nil(t,err)
		
		assert.Equal(t,db.IsConnected(),true)
	})

	t.Run("Collection",func(t *testing.T){
		c := db.Collection()
		assert.NotNil(t,c)

		type todo struct {
			Task string
		}

		rec := todo{
			Task: "insert a record",
		}
		
		ctx,cancel := context.WithTimeout(context.Background(),time.Second * 5)
		defer cancel()

		_, err := c.InsertOne(ctx,rec)
		assert.Nil(t,err)

		err = c.Drop(ctx)
		assert.Nil(t,err)
		
	})
}
