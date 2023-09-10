package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectionDB()(*mongo.Client,error){
	ctx := context.TODO()

	// Opens database
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	err = mongoclient.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	database := mongoclient.Database("mongo_demo")
	fmt.Println("database", database)

	dbList, err := mongoclient.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to list database names: %v", err)
	}
	dbExists := false
	for _, dbName := range dbList {
		if dbName == "mongo_demo" {
			dbExists = true
			break
		}
	}

	if !dbExists {
		// Create the database and its collections
		err = database.CreateCollection(ctx, "users")
		if err != nil {
			return nil, fmt.Errorf("failed to create users collection: %v", err)
		}

		log.Println("Created database:", "mongo_demo")
	} else {
		log.Println("Connected to database:", "mongo_demo")
	}

	return mongoclient, nil
}