package controllers

import "github.com/gofiber/fiber/v2"

// Hello returns a simple Hello message
func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!!")
}
