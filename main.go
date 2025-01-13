package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mattuttis/auth-service/database"
	"github.com/mattuttis/auth-service/routes"
)

func main() {
	_, err := database.ConnectDB()
	if err != nil {
		panic("could not connect to db")
	}

	fmt.Println("Connected to db")

	app := fiber.New()

	//app.Use(cors.New(cors.Config{
	//	AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
	//	AllowHeaders:     "Content-Type,Authorization,Accept,Origin,Access-Control-Request-Method,Access-Control-Request-Headers,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Access-Control-Allow-Methods,Access-Control-Expose-Headers,Access-Control-Max-Age,Access-Control-Allow-Credentials",
	//	AllowCredentials: true,
	//}))

	routes.SetUpRoutes(app)

	err = app.Listen(":8000")
	if err != nil {
		panic("could not start server")
	}
}
