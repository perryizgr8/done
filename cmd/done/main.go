package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	app.Get("/marco", func(c fiber.Ctx) error {
		return c.SendString("polo!")
	})

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("done")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
