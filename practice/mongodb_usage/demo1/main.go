package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//import "go.mongodb.org/mongo-driver/mongo"
//import "github.com/mongodb/mongo-go-driver/mongo"

func main() {
	var (
		appName       string
		clientOptions *options.ClientOptions
		client        *mongo.Client
		err           error
		database      *mongo.Database
		collection    *mongo.Collection
	)

	// 1, 建立连接
	clientOptions = &options.ClientOptions{
		AppName: &appName,
		//ConnectTimeout: time.Duration(5 * time.Second),
		Hosts: []string{"mongodb://36.111.184.221:27017"},
	}

	if client, err = mongo.Connect(context.TODO(), clientOptions); err != nil {
		fmt.Println(err)
		return
	}

	// 2, 选择数据库my_db
	database = client.Database("my_db")

	// 3, 选择表my_collection
	collection = database.Collection("my_collection")

	collection = collection
}
