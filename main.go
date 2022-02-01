package main

import (
	"github.com/DasoTD/fiber-jwt/controller"
	"github.com/DasoTD/fiber-jwt/data"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	app := fiber.New()

	_, err := data.CreateDBEngine()
	if err != nil {
		//panic(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/signup", controller.SignUp)

	app.Post("/login", controller.Login)

	private := app.Group("/private")
	private.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secretss"),
	}))

	private.Get("/private", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": true, "path": "private"})
	})

	app.Get("/public", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": true, "path": "public"})
	})

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
