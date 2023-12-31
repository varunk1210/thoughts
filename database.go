package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func connection() (*mongo.Client, error) {
	// Database connection code goes here
	clientOptions := options.Client().ApplyURI("mongodb+srv://private1:Currency9080@currency.bqm5kzm.mongodb.net/")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	Collection = client.Database("thoughts").Collection("thoughts")

	return client, nil
}

func disconnect(client *mongo.Client) {
	// Database disconnect code goes here
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Println("Error disconnecting from MongoDB: ", err)
	}
	log.Println("Connection to MongoDB closed.")
}

func InsertDocument(data bson.M) error {
	// Insert document code goes here
	_, err := Collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

func GetAllDocuments() []*ThoughtsData {
	// Get all documents code goes here
	filter := bson.M{} // This filter will match all documents in the collection

	var results []*ThoughtsData

	cur, err := Collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem ThoughtsData
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor
	cur.Close(context.TODO())

	for _, result := range results {
		fmt.Printf("%+v\n", *result)
	}
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results
}
