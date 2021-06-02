package routes

import (
	"github.com/gofiber/fiber"
	"go-auth/react-golang/controllers"
)
func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("api/login", controllers.Login)
	app.Get("api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)
}
