package main

import (
	"Go_Exercise/mongo-go-driver/db"
	dbservices "Go_Exercise/mongo-go-driver/dbServices"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// main function
func main() {
	session, err := db.GetConnection()
	if err != nil {
		fmt.Errorf("Error while getting database connection ", err.Error())
		return
	}

	// db and collection references
	cd := session.Database("testDB").Collection("newColl")

	// insert single document
	_, err = cd.InsertOne(context.Background(), bson.M{
		"rollNo": 1,
		"name":   "Mayur Wadekar",
	})

	// error handling
	if err != nil {
		fmt.Errorf("Error while inserting document in connection ", err.Error())
		return
	}

	// creating indexes
	err = dbservices.EnsureIndex(cd, []string{"rollNo"})
	if err != nil {
		fmt.Errorf("Error while creating index on connection ", err.Error())
		return
	}

	// wrap up
	fmt.Println("Everythin works fine... :)")
}
