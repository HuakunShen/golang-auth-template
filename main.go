package main

import (
	"github.com/HuakunShen/golang-auth/database"
	"github.com/HuakunShen/golang-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
