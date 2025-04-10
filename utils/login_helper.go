package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/Dilyxs/gutters_business/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Client struct {
	//ID             primitive.ObjectID `json:"id" bson:"_id"`
	Name           string `json:"name" bson:"Name"`
	Phone          int64  `json:"phone" bson:"Phone"`
	Address        string `json:"address" bson:"Address"`
	Message        string `json:"message" bson:"Message"`
	Estimation     bool   `json:"estimation" bson:"Estimation"`
	BookingTime    string `json:"booking_time" bson:"BookingTime"`
	SignedCustomer bool   `json:"signed_customer" bson:"SignedCustomer"`
	WorkDone       bool   `json:"work_done" bson:"WorkDone"`
	WorkTime       string `json:"work_time" bson:"WorkTime"`
}

func CheckLoginValid(username, password string) bool {
	collection := database.Connect("Login")
	filter := bson.M{"username": username}
	ctx := context.TODO()

	var user User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil || user.Password != password {
		return false
	}
	return true
}

func VericationLogin(c *fiber.Ctx) error {

	var credentials User
	err := c.BodyParser(&credentials)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}
	Correct_or_nah := CheckLoginValid(credentials.Username, credentials.Password)

	if Correct_or_nah {
		return c.Status(200).JSON(fiber.Map{"success": true})
	} else {
		return c.Status(400).JSON(fiber.Map{"success": false})

	}

}

func DashBoard(c *fiber.Ctx) error {
	ctx := context.TODO()
	collection := database.Connect("Form")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not fetch data"})
	}
	defer cursor.Close(ctx)

	var clients []Client
	for cursor.Next(ctx) {
		var client Client
		if err := cursor.Decode(&client); err != nil {
			fmt.Printf("err is %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to decode client"})
		}
		clients = append(clients, client)
	}

	if err := cursor.Err(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cursor error"})
	}

	return c.Status(200).JSON(clients)
}

func UpdateClient(c *fiber.Ctx) error {
	var body Client

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input format"})
	}

	if body.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "name is required"})
	}
	filter := bson.M{"Phone": body.Phone}
	log.Println("Update filter:", filter)

	update := bson.M{
		"$set": bson.M{
			"Name":           body.Name,
			"Phone":          body.Phone,
			"Address":        body.Address,
			"Message":        body.Message,
			"Estimation":     body.Estimation,
			"BookingTime":    body.BookingTime,
			"SignedCustomer": body.SignedCustomer,
			"WorkDone":       body.WorkDone,
			"WorkTime":       body.WorkTime,
		},
	}

	ctx := context.TODO()
	collection := database.Connect("Form")
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updated Client
	err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updated)
	if err != nil {
		log.Println("Update error:", err)
		if err.Error() == "mongo: no documents in result" {
			return c.Status(404).JSON(fiber.Map{"error": "client not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "could not update client"})
	}

	return c.Status(200).JSON(updated)
}

func IntegrateClient(c *fiber.Ctx) error {
	var body Client
	collection := database.Connect("Form")

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}

	_, err := collection.InsertOne(context.TODO(), body)
	if err != nil {
		fmt.Printf("here is error, %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "could not save data"})
	}

	return c.Status(201).JSON(body)
}

func DeleteSingleClient(c *fiber.Ctx) error {

	var payload Client
	collection := database.Connect("Form")

	// Parse the body
	if err := c.BodyParser(&payload); err != nil {
		log.Printf("Error parsing delete payload: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Attempt deletion
	res, err := collection.DeleteOne(context.TODO(), bson.M{"Phone": payload.Phone})
	if err != nil {
		log.Printf("Error deleting client: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete client from database",
		})
	}

	if res.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No client found with that number",
		})
	}

	log.Printf("Client with ID %v successfully deleted", payload.Phone)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Client successfully deleted",
	})
}
