package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// app := fiber.New()

	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Index Page")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("about")
	})

	app.Get("/random.txt", func(c *fiber.Ctx) error {
		return c.SendString("random.txt")
	})

	// Parameters
	app.Get("/user/:name/books/:title", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s\n", c.Params("name"))
		fmt.Fprintf(c, "%s\n", c.Params("title"))
		return nil
	})
	// Plus - greedy - not optional
	app.Get("/user/+", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("+"))
	})

	// Optional parameter
	app.Get("/user/:name?", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("name"))
	})

	// Wildcard - greedy - optional
	app.Get("/user/*", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("*"))
	})

	// This route path will match requests to "/v1/some/resource/name:customVerb", since the parameter character is escaped
	app.Get("/v1/some/resource/name\\:customVerb", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Community")
	})

	// app.Use(func(c *fiber.Ctx) error {
	// 	return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	// })

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	// Pass error to Fiber
	// 	return c.Status(fiber.StatusOK).Render("err", fiber.Map{
	// 		"Msg": "Home Page",
	// 	})
	// })

	app.Get("*", func(c *fiber.Ctx) error {
		// Pass error to Fiber
		return c.Status(fiber.StatusNotFound).Render("err", fiber.Map{
			"Msg": "404 Not Found",
		})
	})

	app.Listen(":3000")
}
