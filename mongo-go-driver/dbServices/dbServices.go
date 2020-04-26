package dbservices

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// EnsureIndex will create index on collection provided
func EnsureIndex(cd *mongo.Collection, indexQuery []string) error {

	// options for index
	opts := options.CreateIndexes().SetMaxTime(5 * time.Second)

	// index model
	index := []mongo.IndexModel{}

	// creating multiple index query
	for _, val := range indexQuery {
		temp := mongo.IndexModel{}
		temp.Keys = bsonx.Doc{{Key: val, Value: bsonx.Int32(1)}}
		index = append(index, temp)
	}

	// executng index query
	_, err := cd.Indexes().CreateMany(context.Background(), index, opts)
	if err != nil {
		fmt.Errorf("Error while executing index Query", err.Error())
		return err
	}

	// if executed successfully then return nil
	return nil
}
