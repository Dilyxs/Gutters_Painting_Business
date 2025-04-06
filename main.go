package main

import (
	"github.com/Dilyxs/gutters_business/utils"
	"github.com/gofiber/fiber/v2"
)

/*
	func SetupRoues(app *fiber.App) {
		app.Get("/Login", handlers.LoginHandler())
	}
*/

func main() {
	app := fiber.New()

	app.Get("/", utils.DashBoard) // see all clients veried and works

	app.Post("/ClientAddition/", utils.IntegrateClient) //create brand new clients works

	app.Put("/UpdateClient/", utils.UpdateClient) //Update a client works

	app.Listen("0.0.0.0:4321")

}
