package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/majjikishore007/todo-fiber.git/controllers"
)


func TodoRoutes (router fiber.Router) {
	router.Post("/", controllers.CreateTodo)
	router.Get("/", controllers.GetTodos)
	router.Get("/:id", controllers.GetTodo)
	router.Put("/:id", controllers.UpdateTodo)
	router.Delete("/:id", controllers.DeleteTodo)
}