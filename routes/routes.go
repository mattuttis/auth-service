package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mattuttis/auth-service/controllers"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
