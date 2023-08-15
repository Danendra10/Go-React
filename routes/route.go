package routes

import (
	"github.com/danendra10/gowlang-first/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// app.Use(middlewares.IsAuthenticated)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
}
