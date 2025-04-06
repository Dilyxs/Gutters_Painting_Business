package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Dilyxs/gutters_business/database"
	"go.mongodb.org/mongo-driver/bson"
)

// Old format
type OldClient struct {
	Name      string `bson:"name"`
	Phone     int64  `bson:"phone"`
	Address   string `bson:"address"`
	Message   string `bson:"message"`
	Timestamp string `bson:"timestamp"`
}

// New format
type Client struct {
	Name           string `bson:"Name"`
	Phone          int64  `bson:"Phone"`
	Address        string `bson:"Address"`
	Message        string `bson:"Message"`
	Estimation     bool   `bson:"Estimation"`
	BookingTime    string `bson:"BookingTime"`
	SignedCustomer bool   `bson:"SignedCustomer"`
	WorkDone       bool   `bson:"WorkDone"`
	WorkTime       string `bson:"WorkTime"`
}

func mainAlt() {

	ctx := context.TODO()

	collection := database.Connect("Form")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			log.Println("decode error:", err)
			continue
		}

		// Check if already migrated (has "Name" field)
		if _, alreadyMigrated := doc["Name"]; alreadyMigrated {
			continue
		}

		// Now decode to OldClient
		var old OldClient
		bsonBytes, _ := bson.Marshal(doc)
		bson.Unmarshal(bsonBytes, &old)

		// Convert phone
		var phoneInt int64
		fmt.Sscanf(fmt.Sprintf("%d", old.Phone), "%v", &phoneInt)

		newClient := Client{
			Name:           old.Name,
			Phone:          phoneInt,
			Address:        old.Address,
			Message:        old.Message,
			BookingTime:    old.Timestamp,
			Estimation:     false,
			SignedCustomer: false,
			WorkDone:       false,
			WorkTime:       "",
		}

		// Use _id to update the doc in-place
		filter := bson.M{"_id": doc["_id"]}

		_, err := collection.ReplaceOne(ctx, filter, newClient)
		if err != nil {
			log.Println("replace error:", err)
		}
	}

	fmt.Println("Migration done. Only updated old-format clients.")
}
