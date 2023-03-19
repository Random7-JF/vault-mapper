package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// Create a new engine
	htmlEngine := html.New("./views", ".html")
	htmlEngine.Debug(true)
	htmlEngine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: htmlEngine,
	})
	routes(app)

	app.Listen(":3000")
}

func routes(app *fiber.App) {

	app.Get("/", indexHandler)
	app.Post("/", postIndexHandler)
}

func indexHandler(c *fiber.Ctx) error {
	// Render index
	return c.Render("pages/index", fiber.Map{
		"Title": "Hello, World!",
	})
}

func postIndexHandler(c *fiber.Ctx) error {
	test := c.FormValue("test")
	fmt.Println(test)
	return c.SendString(test)
}
