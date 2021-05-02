package controllers

import (
	"log"

	"bitbucket.com/Divyanshu_B/go_fiber_rest/models"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
)

func AddProduct(c *fiber.Ctx) error {

	var todo *models.Todo //referece schema
	err := c.BodyParser(&todo)
	if err != nil {
		log.Fatal(err)
	}
	newTodo := &models.Todo{
		Title:     todo.Title,
		Completed: todo.Completed,

		Description: todo.Description,
	}
	err = mgm.Coll(todo).Create(newTodo)

	if err != nil {
		log.Fatal("Creation unsuccessful")
	}

	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Created Successfully",
	})
}
