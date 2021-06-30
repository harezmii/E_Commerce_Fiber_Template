package main

import (
	"apirestecommerce/config"
	"apirestecommerce/database"
	"apirestecommerce/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db := database.Connection()

	route.CategoryRoutes(app,db)

	listenError := app.Listen(config.GetEnv("SERVER_ADDRESS")+":"+config.GetEnv("SERVER_PORT"))
	if listenError != nil {
		println("Server Connection Error")
	}
}
