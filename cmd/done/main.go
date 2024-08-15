package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
)

type Task struct {
	Desc string
	Date time.Time
}

func main() {
	// Initialize a new Fiber app
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/marco", func(c fiber.Ctx) error {
		return c.SendString("polo!")
	})

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.Render("index", fiber.Map{
			"Tasks": []Task{
				{Desc: "item 1"},
				{Desc: "item 2"},
				{Desc: "item 3"},
				{Desc: "lkolz 4"},
			},
		})
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
