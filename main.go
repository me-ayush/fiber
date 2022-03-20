package main

import "github.com/gofiber/fiber/v2"

func indexRoute(res *fiber.Ctx) error {
	return res.SendString("Hello World tghis is ayush")
}

func main() {
	app := fiber.New()
	app.Get("/", indexRoute)

	app.Listen(":3000")
}
