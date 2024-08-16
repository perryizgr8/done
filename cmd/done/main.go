package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"github.com/perryizgr8/done/pkg/common"
	"github.com/perryizgr8/done/pkg/dbase"
)

func main() {
	// Initialize database
	dbase.Init()

	// Initialize a new Fiber app
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/marco", func(c fiber.Ctx) error {
		return c.SendString("polo!")
	})

	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.Render("index", fiber.Map{
			"Tasks": dbase.List(),
		})
	})

	app.Post("/add", func(c fiber.Ctx) error {
		// Add a new task
		desc := c.FormValue("newtask")
		log.Println("desc:", desc)
		err := dbase.Add(common.Task{Desc: desc, Done: time.Now()})
		if err != nil {
			return err
		}
		return c.Redirect().To("/")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
