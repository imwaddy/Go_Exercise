package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// single instance variables
var instance *mongo.Client
var mutex sync.Mutex
var once sync.Once

var (
	username  = ""
	password  = ""
	dbname    = ""
	poolLimit = 100
)

// GetConnection - return mongodb connection
func GetConnection() (*mongo.Client, error) {
	once.Do(func() {
		defer mutex.Unlock()
		mutex.Lock()
		clientOption := options.Client()
		clientOption.SetHosts([]string{"localhost:27017"}).
			SetConnectTimeout(time.Second * 3).   // sets connection timeout
			SetMaxPoolSize(uint64(poolLimit)).    // sets connectionpool limit
			SetReadPreference(readpref.Primary()) // sets read from primary database

		// if username not blank then only perform authentication
		if username != "" {
			cred := options.Credential{}
			cred.Username = username
			cred.Password = password
			cred.AuthSource = dbname
			clientOption.SetAuth(cred)
		}
		client, err := mongo.NewClient(clientOption)
		if err != nil {
			fmt.Errorf(err.Error())
			return
		}
		err = client.Connect(context.Background())
		if err != nil {
			fmt.Errorf(err.Error())
			return
		}
		err = client.Ping(context.Background(), readpref.Primary())
		if err != nil {
			fmt.Errorf(err.Error())
			return
		}
		instance = client
	})
	return instance, nil
}
