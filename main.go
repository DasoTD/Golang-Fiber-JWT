package main

import (
	"fmt"

	"github.com/DasoTD/fiber-jwt/data"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	_, err := data.CreateDBEngine()
	if err != nil {
		panic(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	//app.Post("/signup")

	//app.Post("/login")

	app.Get("/private", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": true, "path": "private"})
	})

	app.Get("/public", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": true, "path": "public"})
	})

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
	fmt.Println("hello ghdhgdhjf")
}