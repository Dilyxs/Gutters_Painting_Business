package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbUsername     string
	dbPassword     string
	dbName         string
	dbCollection_1 string
	dbCollection_2 string

	client       *mongo.Client
	collection_1 *mongo.Collection
	collection_2 *mongo.Collection
	db           *mongo.Database
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUsername = os.Getenv("DBUSERNAME")
	dbPassword = os.Getenv("DBPASSWORD")
	dbName = os.Getenv("DBNAME")
	dbCollection_1 = os.Getenv("dbCollectionLOGIN")
	dbCollection_2 = os.Getenv("dbCollectionDATABASE")
}

func Connect(collectionName string) *mongo.Collection {
	if client == nil {
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(fmt.Sprintf(
			"mongodb+srv://%s:%s@clusterfirst.1zrfovm.mongodb.net/?retryWrites=true&w=majority&appName=ClusterFirst",
			dbUsername, dbPassword)).SetServerAPIOptions(serverAPI)

		var err error
		client, err = mongo.Connect(context.TODO(), opts)
		if err != nil {
			log.Fatal("MongoDB connection error:", err)
		}

		db = client.Database(dbName)
	}

	switch collectionName {
	case dbCollection_1:
		if collection_1 == nil {
			collection_1 = db.Collection(collectionName)
			fmt.Println("Connected to MongoDB, collection")
		}
		return collection_1
	case dbCollection_2:
		if collection_2 == nil {
			collection_2 = db.Collection(collectionName)
			fmt.Println("Connected to MongoDB, collection")
		}
		return collection_2
	default:
		log.Printf("Unknown collection name: %s", collectionName)
		return nil
	}
}
