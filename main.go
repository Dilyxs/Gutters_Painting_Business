package main

import (
	"github.com/Dilyxs/gutters_business/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

/*
	func SetupRoues(app *fiber.App) {
		app.Get("/Login", handlers.LoginHandler())
	}
*/

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type, Accept",
	}))

	app.Get("/", utils.DashBoard) // see all clients veried and works

	app.Post("/ClientAddition/", utils.IntegrateClient) //create brand new clients works

	app.Put("/UpdateClient/", utils.UpdateClient) //Update a client works

	app.Post("/LoginVerification/", utils.VericationLogin) //used for verification

	app.Post("/DeleteClient/", utils.DeleteSingleClient) //delete a single client

	app.Listen("0.0.0.0:4321")
}
