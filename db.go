package main

import (
	"context" //for timeout and cancel the blocking call of database
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func getMongoDbConn() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://ayush:123@fiber-go.e7fh5.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}

func getMongoDbCollection(dbname string, collectionName string) (*mongo.Collection, error) {
	client, err := getMongoDbConn()
	if err != nil {
		log.Fatal(err)
	}
	collections := client.Database(dbname).Collection(collectionName)
	return collections, err
}
