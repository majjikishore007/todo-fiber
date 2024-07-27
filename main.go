package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/majjikishore007/todo-fiber.git/config"
	"github.com/majjikishore007/todo-fiber.git/routes"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Database connection
	config.ConnectToMongo()


	// Middleware
	app.Use(cors.New())



	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	
	api := app.Group("/api")
	routes.TodoRoutes(api.Group("/todos"))
	
	
	app.Listen(":3000")
}