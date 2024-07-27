package controllers

import "github.com/gofiber/fiber/v2"

func CreateTodo(c *fiber.Ctx) error {
	return c.SendString("New todo")
}

func GetTodos(c *fiber.Ctx) error {
	print("All todos")
	return c.SendString("All todos")
}

func GetTodo(c *fiber.Ctx) error {
	return c.SendString("Single todo")
}

func UpdateTodo(c *fiber.Ctx) error {
	return c.SendString("Update todo")
}


func DeleteTodo(c *fiber.Ctx) error {
	return c.SendString("Delete todo")
}

