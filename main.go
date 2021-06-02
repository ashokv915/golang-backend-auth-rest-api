package main

import (
	"github.com/gofiber/fiber"
	"go-auth/react-golang/database"
	"go-auth/react-golang/routes"
	"github.com/gofiber/fiber/middleware/cors"
)

func main() {
	database.Connect()
    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

    app.Listen(":8000")
}
